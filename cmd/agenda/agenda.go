package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vilasbtw/google-calendar-cli/internal/calendar"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "Check agenda",
	Long:  "Find all your agenda events",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.InsertAgenda(args[0])

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("success!")
	},
}
