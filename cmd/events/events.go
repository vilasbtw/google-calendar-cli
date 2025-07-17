package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Check events",
	Long:  "Find all your calendar events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose the command: today or week")
	},
}
