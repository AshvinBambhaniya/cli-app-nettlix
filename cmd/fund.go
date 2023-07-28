/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	// "log"

	// "strconv"
	// "strings"

	// "github.com/PlatformEngineers/cas-go-cli/models"

	"github.com/spf13/cobra"
)

const pathforCSV = `csv/navAMFI.csv`

// fundCmd represents the title command
var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Gives net asset values",
	Long:  `it give the unique scheme code of user and their respective mutual fund plans. It gives net asset value growth across all report.`,
}

func init() {
	rootCmd.AddCommand(fundCmd)
}
