package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vilasbtw/google-calendar-cli/internal/calendar"
)

func init() {
	EventsCmd.AddCommand(EventsWeekCmd)
	EventsCmd.AddCommand(EventsTodayCmd)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Check events",
	Long:  "Find all your calendar events",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
