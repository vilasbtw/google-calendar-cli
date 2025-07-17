// Root Cobra command

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	agenda "github.com/vilasbtw/google-calendar-cli/cmd/agenda"
	events "github.com/vilasbtw/google-calendar-cli/cmd/events"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calendar",
		Short: "Your calendar CLI",
	}
	rootCmd.AddCommand(events.EventsCmd)
	rootCmd.AddCommand(agenda.AgendaCmd)

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
