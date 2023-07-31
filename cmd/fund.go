/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "log"

	"strconv"
	"strings"

	"fmt"
	"log"

	"github.com/PlatformEngineers/cas-go-cli/models"

	"github.com/spf13/cobra"
)

const pathforCSV = `csv/NavAMFI.csv`

// fundCmd represents the title command
var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Gives net asset values",
	Long:  `it give the unique scheme code of user and their respective mutual fund plans. It gives net asset value growth across all report.`,
	Run: func(cmd *cobra.Command, args []string) {

		records, err := models.ReadFile(pathforCSV)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(records)
		count, err := cmd.Flags().GetBool("count")

		if err != nil {
			log.Fatal(err)
		}

		if count {
			fmt.Println("\nTotal Funds Count :", models.GetFundCount(records))
		}

		// /show runtime with max and min runtime records
		nav, err := cmd.Flags().GetBool("nav")
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
		if nav {
			minRow, MaxRow := models.GetMaxMinNetAssetValue(records)
			if max {

				fmt.Println("\nmax Net_Asset_Value is: ", MaxRow.Net_Asset_Value)
				fmt.Println("Scheme Code: ", MaxRow.Scheme_Code)
				fmt.Println("ISIN Growth: ", MaxRow.ISIN_Growth)
				fmt.Println("Scheme Name: ", MaxRow.Scheme_Name)
				fmt.Println("Net Asset Value: ", MaxRow.Net_Asset_Value)
				fmt.Println("Date: ", MaxRow.Date)
			}
			if min {
				fmt.Println("\nmin Net_Asset_Value is: ", minRow.Net_Asset_Value)
				fmt.Println("Scheme Code: ", minRow.Scheme_Code)
				fmt.Println("ISIN Growth: ", minRow.ISIN_Growth)
				fmt.Println("Scheme Name: ", minRow.Scheme_Name)
				fmt.Println("Net Asset Value: ", minRow.Net_Asset_Value)
				fmt.Println("Date: ", minRow.Date)
			}
		}

		//show type flag command
		scheme_name, err := cmd.Flags().GetString("scheme_name")
		if err != nil {
			log.Fatal(err)
		}

	
		if scheme_name != "" {
			searchScheme := models.ListByFundSearch(records, strings.ToLower(scheme_name))
				for i := 0; i < len(searchScheme); i++ {
					fmt.Println("\n", searchScheme[i])
				}
		}

		isin, err := cmd.Flags().GetString("isin")
		if err != nil {
			log.Fatal(err)
		}

	
		if isin != "" {
			searchisin := models.GetIsinScheme(records, strings.ToUpper(isin))
			// fmt.Println(searchisin)
			if searchisin.Net_Asset_Value != 0{
				fmt.Println("Scheme Code: ", searchisin.Scheme_Code)
				fmt.Println("ISIN Growth: ", searchisin.ISIN_Growth)
				fmt.Println("ISIN Div Reinvestment: ", searchisin.ISIN_Div_Reinvestment)
				fmt.Println("Scheme Name: ", searchisin.Scheme_Name)
				fmt.Println("Net Asset Value: ", searchisin.Net_Asset_Value)
				fmt.Println("Date: ", searchisin.Date)
			}else{
				fmt.Println("ISIN code is not available")
			}
		}

		date, err := cmd.Flags().GetInt("date")
		if err != nil {
			log.Fatal(err)
		}
		

		month, err := cmd.Flags().GetString("month")
		if err != nil {
			log.Fatal(err)
		}

		if (date != 0){
			if(month == ""){
				dateStr := strconv.Itoa(date)
				searchDate := models.GetSchemeByDate(records,dateStr)
				for i := 0; i < len(searchDate); i++ {
					fmt.Println("\n", searchDate[i])
				}
			}
			
		}

		if(month != ""){
			if(date == 0){
				searchDateMonth := models.GetSchemeByMonth(records,month)
				for i := 0; i < len(searchDateMonth); i++ {
					fmt.Println("\n", searchDateMonth[i])
				}
			}
			
		}


		if (date != 0){
			if(month != ""){
				dateStr := strconv.Itoa(date)
				searchDateMonth := models.GetSchemeByDateMonth(records,dateStr,month)
				for i := 0; i < len(searchDateMonth); i++ {
					fmt.Println("\n", searchDateMonth[i])
				}
			}
			
		}

	},
}


func init() {
	rootCmd.AddCommand(fundCmd)
	fundCmd.Flags().String("scheme_name", "", "get the list of the schema based on the name")
	fundCmd.Flags().String("isin", "", "get the schema based on the ISIN Div Payout/ ISIN Growth")
	fundCmd.Flags().BoolP("count", "c", false, "Count of the Funds")
	fundCmd.Flags().BoolP("nav", "n", false, "get the Scheme by the --min NAV and --max NAV")
	fundCmd.Flags().BoolP("min", "", false, "get the minimum NAV Scheme")
	fundCmd.Flags().BoolP("max", "", false, "get the maximum NAV Scheme")
	fundCmd.Flags().String("month", "", "get the scheme according to month")
	fundCmd.Flags().Int("date", 0, "get the scheme according to date")
}