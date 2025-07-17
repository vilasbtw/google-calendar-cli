package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vilasbtw/google-calendar-cli/internal/calendar"
)

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Check events",
	Long:  "Find all your calendar events",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		id, err := c.GetAgendaID()

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("id:", id)
	},
}
