package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vilasbtw/google-calendar-cli/internal/calendar"
)

var EventsTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "List all events of this today",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListTodayEvents()
	},
}
