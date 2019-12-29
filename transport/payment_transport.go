package transport

import (
	"context"

	pb "achuala.in/payhub/transport/proto"
)

// PaymentTransport -
type PaymentTransport struct {
}

// InitiatePayment -
func (pt *PaymentTransport) InitiatePayment(ctx context.Context, req *pb.InitiatePaymentRq) (*pb.InitiatePaymentRs, error) {
	return &pb.InitiatePaymentRs{TransactionRefId: "1234", Status: "ACCEPTED"}, nil
}

// GetPayment -
func (pt *PaymentTransport) GetPayment(ctx context.Context, req *pb.GetPaymentRq) (*pb.GetPaymentRs, error) {
	return &pb.GetPaymentRs{TransactionRefId: req.TransactionRefId, Status: "PROCESSING"}, nil
}
