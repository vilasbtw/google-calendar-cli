package calendar

// Connecting our client to GCP

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "Roles"

var (
	ErrAgendaNotFound          = errors.New("unable to find the agenda")
	ErrUnableToInsertAgenda    = errors.New("unable to insert the agenda")
	ErrUnableToFindWeekEvents  = errors.New("unable to find week events")
	ErrUnableToFindTodayEvents = errors.New("unable to find today events")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarId string
}

func NewClient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("./credentials.json")

	if err != nil {
		log.Fatal("unable to read credentials.json")
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("unable to create Google Calendar Service: %s", err.Error())
	}

	return &Calendar{
		Service: service,
	}
}

func (c *Calendar) GetAgendaID() error {
	// GET request for agenda "Roles"
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		return ErrAgendaNotFound
	}

	for _, v := range list.Items {
		if v.Summary == AGENDA {
			c.CalendarId = v.Id
		}
	}
	return nil
}

func (c *Calendar) InsertAgenda(id string) error {
	entry := &gCalendar.CalendarListEntry{
		Id: c.CalendarId,
	}

	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrUnableToInsertAgenda
	}

	return nil
}

func (c *Calendar) ListWeekEvents() error {
	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)
	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrUnableToFindWeekEvents
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s | at %s\n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}

func (c *Calendar) ListTodayEvents() error {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrUnableToFindTodayEvents
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s | at %s\n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}
