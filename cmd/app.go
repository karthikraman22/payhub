package main

import (
	"context"
	"net"
	"net/http"
	"os"

	pb "achuala.in/payhub/transport/proto"

	"achuala.in/payhub/config"
	promZap "achuala.in/payhub/core/zap"
	grpSvc "achuala.in/payhub/transport"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func main() {
	// Configuration
	conf, err := config.GetConfig(os.Getenv("BUILD_ENV"), nil)
	if err != nil {
		panic(err) // othing we can do
	}
	if conf.GetString("environment") != "production" {
		conf.Debug()
	}

	/*
	 * PreRequisite: Health Check + Expose status Prometheus metrics gauge
	 * **************************** */
	healthChecker := healthcheck.NewMetricsHandler(prometheus.DefaultRegisterer, "health_check")
	healthChecker.AddLivenessCheck("Goroutine Threshold", healthcheck.GoroutineCountCheck(conf.GetInt("health_check.goroutine_threshold")))

	// Expose to HTTP
	http.HandleFunc(conf.GetString("health_check.route.group")+conf.GetString("health_check.route.live"), healthChecker.LiveEndpoint)
	http.HandleFunc(conf.GetString("health_check.route.group")+conf.GetString("health_check.route.ready"), healthChecker.ReadyEndpoint)

	// Logger
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.UnmarshalText([]byte(conf.GetString("log.level")))
	zapConfig.Development = conf.GetString("environment") != "production"
	// Expose log level Prometheus metrics
	hook := promZap.MustNewPrometheusHook([]zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel,
		zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel, zapcore.DebugLevel})
	logger, _ := zapConfig.Build(zap.Hooks(hook))

	// Sentry
	if conf.GetString("sentry.dsn") != "" {
		atom := zap.NewAtomicLevel()
		err := atom.UnmarshalText([]byte(conf.GetString("sentry.log_level")))
		if err != nil {
			logger.Fatal("Failed to parse Zap-Sentry log level", zap.String("sentry.log_level", conf.GetString("sentry.log_level")))
		}

		cfg := zapsentry.Configuration{
			Level: atom.Level(), //when to send message to sentry
			Tags: map[string]string{
				"component": conf.GetString("app.name"),
			},
		}
		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(conf.GetString("sentry.dsn")))
		//in case of err it will return noop core. so we can safely attach it
		if err != nil {
			logger.Fatal("failed to init sentry / zap")
		}
		logger = zapsentry.AttachCoreToLogger(core, logger)
	}
	defer logger.Sync()

	/*
	 * PreRequisite: gRPC
	 * **************************** */
	lis, err := net.Listen("tcp", ":"+conf.GetString("app.grpc.port"))
	if err != nil || lis == nil {
		logger.Sugar().Errorf("gRPC failed to listen: %v", err)
	}
	logger.Info("gRPC is about to start listening for incoming requests", zap.String("port", conf.GetString("app.grpc.port")))

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	defer grpcServer.GracefulStop()

	paymentGrpcService := grpSvc.PaymentTransport{}
	pb.RegisterPaymentServiceServer(grpcServer, &paymentGrpcService)

	// Start listening to gRPC requests
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("gRPC failed listening for incoming requests",
				zap.String("port", conf.GetString("app.grpc.port")),
				zap.String("error", err.Error()),
			)

			healthChecker.AddReadinessCheck("gRPC", func() error { return err }) // Permanent, take us down.
		} else {
			logger.Info("gRPC is listening for incoming requests", zap.String("port", conf.GetString("app.grpc.port")))
		}

	}()

	// Register gRPC as HTTP Gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, ":"+conf.GetString("app.grpc.port"), opts)
	if err != nil {
		logger.Sugar().Errorf("Failed to register Payment Service %+v", err)
	}

	/*
	 * Start listening for incoming HTTP requests
	 * **************************** */
	logger.Info("Starting..")
	http.ListenAndServe(":"+conf.GetString("app.port"), mux)
}
