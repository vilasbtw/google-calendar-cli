// Root Cobra command

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	events "github.com/vilasbtw/google-calendar-cli/cmd/events"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calendar",
		Short: "Your calendar CLI",
	}
	rootCmd.AddCommand(events.EventsCmd)
	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
