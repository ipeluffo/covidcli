package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ipeluffo/covidcli/models"

	"github.com/urfave/cli/v2"
)

const (
	flagFrom    = "from"
	flagTo      = "to"
	flagCountry = "country"
	dateFormat  = "2006-01-02"
)

func commandAction(ctx *cli.Context) error {
	var from *time.Time = ctx.Timestamp(flagFrom)
	var to *time.Time = ctx.Timestamp(flagTo)
	var country string = ctx.String(flagCountry)

	// I'd like to use https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#9739c95f-ef1d-489b-97a9-0a6dfe2f74d8
	// but dates filters don't work
	apiURL := "https://api.covid19api.com/country/%s?from=%s&to=%s"
	url := fmt.Sprintf(apiURL, country, from.Format(dateFormat), to.Format(dateFormat))

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var records []models.Stats
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &records)
	// Alternative
	// json.NewDecoder(resp.Body).Decode(&records)

	var lastDeaths int = 0
	var lastDeathsDiff int = 0

	for _, stats := range records {
		// Skip cities stats.
		// This is a workaround since we cannot use totals per country endpoint.
		if stats.Province != "" {
			continue
		}

		diff := stats.Deaths - lastDeaths
		var diffEmoji string
		var deathsEmoji string

		if diff == 0 {
			deathsEmoji = "🎉"
		} else {
			deathsEmoji = "💀"
		}

		if diff > lastDeathsDiff {
			diffEmoji = "📈"
		} else if diff < lastDeathsDiff {
			diffEmoji = "📉"
		}

		// This stat usually comes as zero, if this is case then calculate it
		var active int = stats.Active
		if active == 0 {
			active = stats.Confirmed - stats.Deaths - stats.Recovered
		}

		fmt.Println(stats.Date)
		fmt.Println("Confirmed:", stats.Confirmed)
		fmt.Println("Deaths:", stats.Deaths, deathsEmoji, "Diff:", diff, diffEmoji)
		fmt.Println("Recovered:", stats.Recovered)
		fmt.Println("Active:", active)
		fmt.Println()

		lastDeaths = stats.Deaths
		lastDeathsDiff = diff
	}

	return nil
}

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Ignacio Peluffo",
				Email: "ipeluffo@gmail.com",
			},
		},
		Version: "2020.04.19",
		Name:    "covidcli",
		Usage:   "Get COVID-19 stats on your terminal",
		Action:  commandAction,
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:     flagFrom,
				Required: true,
				Layout:   "2006-01-02",
			},
			&cli.TimestampFlag{
				Name:     flagTo,
				Required: true,
				Layout:   "2006-01-02",
			},
			&cli.StringFlag{
				Name:     flagCountry,
				Required: true,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
