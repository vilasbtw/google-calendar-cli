package calendar

// Connecting our client to GCP

import (
	"context"
	"errors"
	"log"
	"os"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "Roles"

var (
	ErrAgendaNotFound       = errors.New("unable to find the agenda")
	ErrUnableToInsertAgenda = errors.New("unable to insert the agenda")
)

type Calendar struct {
	Service *gCalendar.Service
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

func (c *Calendar) GetAgendaID() (string, error) {
	// GET request for agenda "Roles"
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatal("unable to list agendas.")
	}

	for _, v := range list.Items {
		if v.Summary == AGENDA {
			return v.Id, nil
		}
	}
	return "", ErrAgendaNotFound
}

func (c *Calendar) InsertAgenda(id string) error {

	entry := &gCalendar.CalendarListEntry{
		Id: id,
	}

	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrUnableToInsertAgenda
	}

	return nil
}
