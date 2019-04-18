package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// API demos.
	cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "Serve our web application",
		Run: func(cmd *cobra.Command, args []string) {
			Serve(port)
		},
	}
)

// Flag vars.
var (
	port int
)

func init() {
	cmdServe.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port to listen on")

	rootCmd.AddCommand(cmdServe)
}

func Serve(port int) {
	fmt.Printf("Listening on port: %d\n", port)
}
