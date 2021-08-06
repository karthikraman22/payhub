FROM golang:latest as builder

# Arguments to Env variables
ARG build_env
ARG app_port
ARG app_grpc_port
ARG protobuf_release_tag

ENV BUILD_ENV $build_env
ENV APP_PORT $app_port
ENV APP_GRPC_PORT $app_grpc_port
ENV PROTOBUF_RELEASE_TAG $protobuf_release_tag

# Path
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
WORKDIR /app

# Copy our code
ADD . .

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod vendor 

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app1 ./cmd/app.go


FROM scratch
#FROM alpine:latest

COPY --from=builder /app1 /app
#RUN chmod +x /app
EXPOSE ${APP_GRPC_PORT} ${APP_PORT}
CMD ["/app"]