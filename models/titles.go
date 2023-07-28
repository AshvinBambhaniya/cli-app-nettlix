package models

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Title struct {
	Id                   string
	Title                string
	Show_type            string
	Description          string
	Release_year         int
	Age_certification    int
	Runtime              int
	Genres               string
	Production_countries string
	Seasons              float64
	Imdb_id              string
	Imdb_score           float64
	Imdb_votes           float64
	Tmdb_popularity      float64
	Tmdb_score           float64
}

type Titles []Title

func ReadCsvFile(filePath string) (Titles, error) {
	var titles Titles
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
	for _, rowTitle := range records {
		Release_year, err := strconv.Atoi(rowTitle[4])
		if err != nil {
			Release_year = 0
		}
		Age_certification, err := strconv.Atoi(rowTitle[5])
		if err != nil {
			Age_certification = 0
		}
		runtime, err := strconv.Atoi(rowTitle[6])
		if err != nil {
			runtime = 0
		}
		Seasons, err := strconv.ParseFloat(rowTitle[9], 64)
		if err != nil {
			Seasons = 0
		}
		Imdb_score, err := strconv.ParseFloat(rowTitle[11], 64)
		if err != nil {
			Imdb_score = 0
		}
		Imdb_votes, err := strconv.ParseFloat(rowTitle[12], 64)
		if err != nil {
			Imdb_votes = 0
		}
		Tmdb_popularity, err := strconv.ParseFloat(rowTitle[13], 64)
		if err != nil {
			Tmdb_popularity = 0
		}
		Tmdb_score, err := strconv.ParseFloat(rowTitle[14], 64)
		if err != nil {
			Tmdb_popularity = 0
		}

		list := Title{
			Id:                   rowTitle[0],
			Title:                rowTitle[1],
			Show_type:            rowTitle[2],
			Description:          rowTitle[3],
			Release_year:         Release_year,
			Age_certification:    Age_certification,
			Runtime:              runtime,
			Genres:               rowTitle[7],
			Production_countries: rowTitle[8],
			Seasons:              Seasons,
			Imdb_id:              rowTitle[10],
			Imdb_score:           Imdb_score,
			Imdb_votes:           Imdb_votes,
			Tmdb_popularity:      Tmdb_popularity,
			Tmdb_score:           Tmdb_score,
		}
		if list.Id != "" {
			titles = append(titles, list)
		}

	}
	return titles, nil
}

func GetCount(records Titles, showType string) int {
	var totalCount int
	if showType != "" {
		shows := ListByType(records, showType)
		totalCount = len(shows)
	} else {
		totalCount = len(records)
	}
	return totalCount

}

func ListByType(records Titles, showType string) []string {
	var showTitle []string
	for _, v := range records {
		if v.Show_type == showType {
			showTitle = append(showTitle, v.Title)
		}
	}
	return showTitle
}

func ListByYear(records Titles, year int) []string {
	var showTitle []string
	for _, v := range records {
		if v.Release_year == year {
			showTitle = append(showTitle, v.Title)
		}
	}
	return showTitle
}

func ListUniqueGeneres(records Titles) []string {
	var uniqueGeneres []string
	for _, v := range records {
		replaceGeneres := strings.Replace(v.Genres, "[", "", 1)
		finalGeneres := strings.Replace(replaceGeneres, "]", "", 1)
		splitGeneres := strings.Split(finalGeneres, ",")
		for _, list := range splitGeneres {
			list = strings.TrimSpace(list)
			if !lo.Contains(uniqueGeneres, list) {
				uniqueGeneres = append(uniqueGeneres, list)
			}
		}
	}
	return uniqueGeneres
}

func GetMaxMinRuntimeShow(records Titles) (Title, Title) {
	var min Title
	var max Title
	for _, row := range records {
		if row.Runtime == 0 {
			continue
		}
		if (min == Title{} && max == Title{}) {
			min = row
			max = row
		}
		if row.Runtime < min.Runtime {
			min = row
		}
		if row.Runtime > max.Runtime {
			max = row
		}
	}
	return min, max

}

func ListUniqueCountries(records Titles) []string {
	var uniqueCountries []string
	for _, v := range records {
		replaceCountries := strings.Replace(v.Production_countries, "[", "", 1)
		finalCountries := strings.Replace(replaceCountries, "]", "", 1)
		splitCountries := strings.Split(finalCountries, ",")
		for _, list := range splitCountries {
			list = strings.TrimSpace(list)
			if !lo.Contains(uniqueCountries, list) {
				uniqueCountries = append(uniqueCountries, list)
			}
		}
	}
	return uniqueCountries
}

func ListBySearch(records Titles, movie string) Titles {

	var searchedMovie Titles
	for _, row := range records {
		if strings.HasPrefix(strings.ToLower(row.Title), strings.ToLower(movie)) {
			searchedMovie = append(searchedMovie, row)
		}
	}
	return searchedMovie
}
