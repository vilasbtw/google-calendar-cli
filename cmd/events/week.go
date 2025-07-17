package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vilasbtw/google-calendar-cli/internal/calendar"
)

var EventsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "List all events of this week",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListWeekEvents()
	},
}
