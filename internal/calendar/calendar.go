package calendar

// Connecting our client to GCP

import (
	"context"
	"log"
	"os"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Calendar struct {
	Service *gCalendar.Service
}

func NewClient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("./credentials.json")

	if err != nil {
		log.Fatal("Unable to read credentials.json")
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatal("Unable to create Google Calendar Service: %s", err.Error())
	}

	return &Calendar{
		Service: service,
	}
}
