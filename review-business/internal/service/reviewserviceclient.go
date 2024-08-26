package service

import (
	"context"
	pb "review-business/api/reply/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewReviewServiceClient() pb.ReplyClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return pb.NewReplyClient(conn)
}
