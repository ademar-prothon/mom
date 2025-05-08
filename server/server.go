package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ademar-prothon/mom/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type motoTrackServer struct {
	pb.UnimplementedMotoTrackServer

	sessions []*pb.Session
}

func (m *motoTrackServer) StartSession(ctx context.Context, req *pb.StartSessionRequest) (*pb.StartSessionResponse, error) {

	log.Printf("StartSession called for bike: %s", req.BikeId)
	sessionID := fmt.Sprintf("session-%s", req.BikeId)
	sid := &pb.SessionId{SessionId: sessionID}
	m.sessions = append(m.sessions, &pb.Session{Id: sid})
	return &pb.StartSessionResponse{
		SessionId: sessionID,
	}, nil
}

func (m *motoTrackServer) EndSession(ctx context.Context, id *pb.SessionId) (*pb.Ack, error) {
	log.Printf("EndSession called for bike: %s", id.SessionId)
	// here remove session from array
	return nil, status.Errorf(codes.Unimplemented, "method EndSession not implemented")
}

func (m *motoTrackServer) StreamTelemetry(stream pb.MotoTrack_StreamTelemetryServer) error {
	// here add TelemetryData in a session
	return status.Errorf(codes.Unimplemented, "method StreamTelemetry not implemented")
}

func (m *motoTrackServer) GetSessionData(ctx context.Context, id *pb.SessionId) (*pb.Session, error) {
	// here return Session from the server
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionData not implemented")
}
