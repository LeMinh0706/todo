package streaming

import (
	"time"

	"github.com/LeMinh0706/todo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

// StreamServerTime implements proto.StreamingServiceServer.
func (s *Service) StreamServerTime(req *proto.StreamServerRequest, stream grpc.ServerStreamingServer[proto.StreamServerResponse]) error {
	if req.GetIntervalSeconds() == 0 {
		return status.Error(codes.InvalidArgument, "interval must be set")
	}

	interval := time.Duration(req.GetIntervalSeconds()) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			currentTime := time.Now()
			resp := &proto.StreamServerResponse{
				CurrentTime: timestamppb.New(currentTime),
			}

			if err := stream.Send(resp); err != nil {
				return err
			}

		}
	}

}

func NewService() proto.StreamingServiceServer {
	return &Service{}
}
