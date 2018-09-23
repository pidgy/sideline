package sideline

// staff.go acts as a common library for sideline methods

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

var apiKey = ""

// format
const (
	formatXml = "xml"
	formatJson = "json"
)

// service-name
const (
	serviceNflTeams = "nfl-teams"
	servicePlayers = "players"
	serviceWeeklyProjections = "weekly-projections"
	serviceWeeklyRankings = "weekly-rankings"
	serviceSchedule = "schedule"
)

func APIKey(key string) {
	apiKey = key
}

func ffnGetRequest(args ...string) []byte {
	request := "https://www.fantasyfootballnerd.com/service"

	for i, str := range args {
		if i == 1 {
			request = makeString(request, "/", str, "/", apiKey)
		} else {
			request = makeString(request, "/", str)
		}
	}

	response, err := http.Get(request)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("The HTTP request failed")
		return nil
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to read HTTP response")
	}

	return contents
}

func makeString(s ...string) string {
	var buffer bytes.Buffer
	for _, str := range s {
		buffer.WriteString(str)
	}
	return buffer.String()
}
