package transport

import (
	"context"

	pb "achuala.in/payhub/transport/proto"
)

// PaymentTransport -
type PaymentTransport struct {
}

// InitiatePayment -
func (pt *PaymentTransport) InitiatePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{TransactionRefId: "1234", Status: "ACCEPTED"}, nil
}
