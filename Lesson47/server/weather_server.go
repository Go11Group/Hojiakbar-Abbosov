package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/weather"
)

type server1 struct {
    pb.UnimplementedWeatherServiceServer
}

func (s *server1) GetCurrentWeather(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
    return &pb.WeatherResponse{
        Temperature: "25°C",
        Humidity:    "60%",
        WindSpeed:   "10 km/h",
    }, nil
}

func (s *server1) GetWeatherForecast(ctx context.Context, req *pb.ForecastRequest) (*pb.ForecastResponse, error) {
    forecasts := []*pb.WeatherForecast{
        {Date: "2024-06-28", Temperature: "26°C", Description: "Sunny"},
        {Date: "2024-06-29", Temperature: "24°C", Description: "Partly Cloudy"},
    }
    return &pb.ForecastResponse{Forecasts: forecasts}, nil
}

func (s *server1) ReportWeatherCondition(ctx context.Context, req *pb.WeatherConditionReport) (*pb.WeatherConditionResponse, error) {
    log.Printf("Reported condition: %s in %s", req.Condition, req.Location)
    return &pb.WeatherConditionResponse{Success: true, Message: "Condition reported successfully"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterWeatherServiceServer(s, &server1{})
    log.Println("WeatherService server is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
