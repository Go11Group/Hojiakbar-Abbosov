package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    weatherpb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/weather"
    transportpb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/transport"

)

func callWeatherService(client weatherpb.WeatherServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    weatherRes, err := client.GetCurrentWeather(ctx, &weatherpb.WeatherRequest{Location: "New York"})
    if err != nil {
        log.Fatalf("could not get weather: %v", err)
    }
    log.Printf("Current Weather: %s, %s, %s", weatherRes.Temperature, weatherRes.Humidity, weatherRes.WindSpeed)
}

func callTransportService(client transportpb.TransportServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    busRes, err := client.GetBusSchedule(ctx, &transportpb.BusScheduleRequest{BusNumber: "123"})
    if err != nil {
        log.Fatalf("could not get bus schedule: %v", err)
    }
    log.Printf("Bus Schedule: %s", busRes.Schedule)
}

func main() {
    weatherConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer weatherConn.Close()
    weatherClient := weatherpb.NewWeatherServiceClient(weatherConn)

    transportConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer transportConn.Close()
    transportClient := transportpb.NewTransportServiceClient(transportConn)

    callWeatherService(weatherClient)
    callTransportService(transportClient)
}
