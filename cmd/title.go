/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PlatformEngineers/cas-go-cli/models"

	"github.com/spf13/cobra"
)

const pathCSV = `csv/titles.csv`

// titleCmd represents the title command
var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Total count of the titles",
	Long:  `All the titles of the dataset counts are avalialbe by the count command`,
	Run: func(cmd *cobra.Command, args []string) {
		records, err := models.ReadCsvFile(pathCSV)
		if err != nil {
			log.Fatal(err)
		}
		//count flag command
		count, err := cmd.Flags().GetBool("count")

		if err != nil {
			log.Fatal(err)
		}

		//show type flag command
		showtype, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}

		if showtype != "" {
			shows := models.ListByType(records, strings.ToUpper(showtype))
			for i := 0; i < len(shows); i++ {
				fmt.Println(i, shows[i])
			}
		}

		if count {
			fmt.Println("\nTotal Titles Count :", models.GetCount(records, strings.ToUpper(showtype)))
		}

		//show Year flag command
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			log.Fatal(err)
		}
		yearStr := strconv.Itoa(year)
		if len(yearStr) == 4 {
			shows := models.ListByYear(records, year)
			fmt.Println("#\t Year \t Title")
			fmt.Println("--------------------------------")
			for i := 0; i < len(shows); i++ {
				fmt.Println(i, "\t", year, "\t", shows[i])
			}
		}

		// /show runtime with max and min runtime records
		runtime, err := cmd.Flags().GetBool("runtime")
		if err != nil {
			log.Fatal(err)
		}
		max, err := cmd.Flags().GetBool("max")
		if err != nil {
			log.Fatal(err)
		}
		min, err := cmd.Flags().GetBool("min")
		if err != nil {
			log.Fatal(err)
		}
		if runtime {
			minRow, MaxRow := models.GetMaxMinRuntimeShow(records)
			if max {

				fmt.Println("\nmax runtime is: ", MaxRow.Runtime)
				fmt.Println("Id\t \ttitle\t \t\ttype\t release_year\t age_certification\t runtime\t")
				fmt.Println(MaxRow.Id, "\t", MaxRow.Title, "\t", MaxRow.Show_type, "\t\t", MaxRow.Release_year, "\t\t", MaxRow.Age_certification, "\t\t", MaxRow.Runtime)
			}
			if min {
				fmt.Println("\nmin runtime is: ", minRow.Runtime)
				fmt.Println("Id\t \ttitle\t type\t release_year\t age_certification\t runtime\t")
				fmt.Println(minRow.Id, "\t", minRow.Title, "\t", minRow.Show_type, "\t", minRow.Release_year, "\t\t", minRow.Age_certification, "\t\t\t", minRow.Runtime)
			}
		}
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Display the searched movie result",
	Long:  `The command will display the show of movie passed as an argument`,
	Run: func(cmd *cobra.Command, args []string) {

		records, err := models.ReadCsvFile(pathCSV)
		if err != nil {
			log.Fatal(err)
		}
		if len(args) != 0 {
			movie := args[0]
			searchMovie := models.ListBySearch(records, movie)
			for i := 0; i < len(searchMovie); i++ {
				fmt.Println("\n", searchMovie[i])
			}
		}

	}}

func init() {
	rootCmd.AddCommand(titleCmd)
	titleCmd.AddCommand(searchCmd)
	titleCmd.Flags().BoolP("count", "c", false, "Count of the titles")
	titleCmd.Flags().String("type", "", "get the list of the movies based on the type")
	titleCmd.Flags().Int("year", 0, "get list of the title based on the year")
	titleCmd.Flags().BoolP("runtime", "r", false, "get the shows by the --min runtime show and --max runtime show")
	titleCmd.Flags().BoolP("min", "", false, "get the minimum runtime show")
	titleCmd.Flags().BoolP("max", "", false, "get the maximum runtime show")
}
