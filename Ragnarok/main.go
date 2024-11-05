package cmd

import (
	"fmt"
	"os"

	"github.com/FOXHOUND0x/ragnarok/pkg/monitor"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ragnarok",
	Short: "Ragnarok is a CLI tool for monitoring Docker containers",
	Long:  `Ragnarok is a command-line tool written in Go for monitoring the health of Docker containers running on your local machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Ragnarok...")

		mon, err := monitor.NewMonitor()
		if err != nil {
			fmt.Println("Error initializing monitor:", err)
			return
		}

		mon.DisplayContainerHealth()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
