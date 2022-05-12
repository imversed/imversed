package main

import (
	"github.com/spf13/cobra"
)

const (
	flagLong = "long"
)

func init() {
	infoCmd.Flags().Bool(flagLong, false, "Print full information")
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print version info",
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}
