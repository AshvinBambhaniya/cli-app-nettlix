package models

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	// "github.com/samber/lo"
)

type Fund struct {
	Scheme_Code           string
	ISIN_Growth           string
	ISIN_Div_Reinvestment string
	Scheme_Name           string
	Net_Asset_Value       float64
	Date                  string
}

type Funds []Fund

func ReadFile(filePath string) (Funds, error) {
	var funds Funds
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, rowFund := range records {
		Net_Asset_Value, err := strconv.ParseFloat(rowFund[4], 64)
		if err != nil {
			Net_Asset_Value = 0
		}

		list := Fund{
			Scheme_Code:           rowFund[0],
			ISIN_Growth:           rowFund[1],
			ISIN_Div_Reinvestment: rowFund[2],
			Scheme_Name:           rowFund[3],
			Net_Asset_Value:       Net_Asset_Value,
			Date:                  rowFund[5],
		}
		if list.Scheme_Code != "" {
			funds = append(funds, list)
		}

	}
	return funds, nil
}

func ListByFundSearch(records Funds, fundname string) Funds {

	var searchedFund Funds
	for _, row := range records {
		if strings.HasPrefix(strings.ToLower(row.Scheme_Name), strings.ToLower(fundname)) {
			searchedFund = append(searchedFund, row)
		}
	}
	return searchedFund
}

func GetFundCount(records Funds) int {
	var isinCount []string

	for _, v := range records {
		if v.Net_Asset_Value != 0 {
			isinCount = append(isinCount, v.ISIN_Growth)
		}
	}

	return len(isinCount)

}

func GetMaxMinNetAssetValue(records Funds) (Fund, Fund) {
	var min Fund
	var max Fund
	for _, row := range records {
		if row.Net_Asset_Value == 0 {
			continue
		}
		if (min == Fund{} && max == Fund{}) {
			min = row
			max = row
		}
		if row.Net_Asset_Value < min.Net_Asset_Value {
			min = row
		}
		if row.Net_Asset_Value > max.Net_Asset_Value {
			max = row
		}
	}
	return min, max

}