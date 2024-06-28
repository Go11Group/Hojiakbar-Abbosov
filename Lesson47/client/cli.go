package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	weatherpb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/weather"
	transportpb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson46/genproto/transport"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s [weather|transport] <command> [args]", os.Args[0])
	}

	service := os.Args[1]
	command := os.Args[2]

	switch service {
	case "weather":
		runWeatherCommand(command)
	case "transport":
		runTransportCommand(command)
	default:
		log.Fatalf("Unknown service: %s", service)
	}
}

func runWeatherCommand(command string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := weatherpb.NewWeatherServiceClient(conn)

	switch command {
	case "current":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.GetCurrentWeather(ctx, &weatherpb.WeatherRequest{Location: "New York"})
		if err != nil {
			log.Fatalf("Error getting weather: %v", err)
		}
		fmt.Printf("Current Weather: %s, %s, %s\n", res.Temperature, res.Humidity, res.WindSpeed)
	case "forecast":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.GetWeatherForecast(ctx, &weatherpb.ForecastRequest{Location: "New York", Days: 2})
		if err != nil {
			log.Fatalf("Error getting forecast: %v", err)
		}
		for _, forecast := range res.Forecasts {
			fmt.Printf("Date: %s, Temperature: %s, Description: %s\n", forecast.Date, forecast.Temperature, forecast.Description)
		}
	case "report":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.ReportWeatherCondition(ctx, &weatherpb.WeatherConditionReport{Location: "New York", Condition: "Sunny"})
		if err != nil {
			log.Fatalf("Error reporting condition: %v", err)
		}
		fmt.Printf("Report success: %t, Message: %s\n", res.Success, res.Message)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func runTransportCommand(command string) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := transportpb.NewTransportServiceClient(conn)

	switch command {
	case "schedule":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.GetBusSchedule(ctx, &transportpb.BusScheduleRequest{BusNumber: "123"})
		if err != nil {
			log.Fatalf("Error getting schedule: %v", err)
		}
		fmt.Printf("Bus Schedule: %s\n", res.Schedule)
	case "location":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.TrackBusLocation(ctx, &transportpb.BusLocationRequest{BusNumber: "123"})
		if err != nil {
			log.Fatalf("Error tracking location: %v", err)
		}
		fmt.Printf("Bus Location: %s\n", res.Location)
	case "report":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.ReportTrafficJam(ctx, &transportpb.TrafficJamReport{Location: "Main St", Description: "Heavy traffic"})
		if err != nil {
			log.Fatalf("Error reporting traffic jam: %v", err)
		}
		fmt.Printf("Report success: %t, Message: %s\n", res.Success, res.Message)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
