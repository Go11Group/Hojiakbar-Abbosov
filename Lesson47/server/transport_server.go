package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
	pb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/transport"


)

type server struct {
    pb.UnimplementedTransportServiceServer
}

func (s *server) GetBusSchedule(ctx context.Context, req *pb.BusScheduleRequest) (*pb.BusScheduleResponse, error) {
    return &pb.BusScheduleResponse{Schedule: "Bus 123: 10:00 AM, 12:00 PM, 2:00 PM"}, nil
}

func (s *server) TrackBusLocation(ctx context.Context, req *pb.BusLocationRequest) (*pb.BusLocationResponse, error) {
    return &pb.BusLocationResponse{Location: "Near Central Park"}, nil
}

func (s *server) ReportTrafficJam(ctx context.Context, req *pb.TrafficJamReport) (*pb.TrafficJamResponse, error) {
    log.Printf("Reported traffic jam: %s at %s", req.Description, req.Location)
    return &pb.TrafficJamResponse{Success: true, Message: "Traffic jam reported successfully"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTransportServiceServer(s, &server{})
    log.Println("TransportService server is running on port 50052...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
