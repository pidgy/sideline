package sideline

// wideout.go performs RESTful GET requests for the data and marshals
// the data according to specifications defined in scout.go

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
)

type Wideout struct {
	name string
}

// Games will return a slice of nflMatchup types for the current week in football
// as well as the string value for the current week
func (w Wideout) Games() (nflMatchups, string) {
	var matchups nflMatchups

	response := ffnGetRequest(serviceSchedule, formatJson)
	playerMap := map[string]interface{}{}

	err := json.Unmarshal(response, &playerMap)
	if err != nil {
		log.Error().Err(err).Str("Function", "Teams").Msg("Failed to marshal HTTP response")
		return nil, ""
	}

	currentWeek := playerMap["currentWeek"]

	for key, value := range playerMap {
		if key == "Schedule" {
			for _, data := range value.([]interface{}) {
				stats := data.(map[string]interface{})
				if stats["gameWeek"] == currentWeek {
					matchups = append(matchups, nflMatchup{
						gameId: stats["gameId"].(string),
						gameWeek: stats["gameWeek"].(string),
						gameDate: stats["gameDate"].(string),
						awayTeam: stats["awayTeam"].(string),
						homeTeam: stats["homeTeam"].(string),
						gameTimeET: stats["gameTimeET"].(string),
						tvStation: stats["tvStation"].(string),
						winner: stats["winner"].(string),
					})
				}
			}
		}
	}

	fmt.Println(apiKey)
	return matchups, currentWeek.(string)
}

// Teams will return an unsorted slice of nflTeam structs
func (w Wideout) Teams() nflTeams {
	var teams nflTeams

	response := ffnGetRequest(serviceNflTeams, formatJson)
	teamMap := map[string]interface{}{}

	err := json.Unmarshal(response, &teamMap)
	if err != nil {
		log.Error().Err(err).Str("Function", "Teams").Msg("Failed to marshal HTTP response")
		return nil
	}

	if responseContent, ok := teamMap["NFLTeams"]; ok {
		switch responseContent.(type) {
		case []interface{}:
			responseSlice := responseContent.([]interface{})
			teams = make(nflTeams, len(responseSlice))

			for i, data := range responseSlice {
				team := data.(map[string]interface{})

				teams[i] = nflTeam{
					code: team["code"].(string),
					fullName: team["fullName"].(string),
					shortName: team["shortName"].(string),
				}
			}
		default:
			log.Error().Str("Function", "Teams").Msg("No NFL teams exist in the response")
		}

	}

	return teams
}

// WeeklyProjections will return an unsorted slice of weekly projections for NFL players
// binned by position ("QB", "RB", "WR", "TE", "K", "DEF")
func (w Wideout) WeeklyProjections(week string) nflPlayersWeeklyProjection{
	players := nflPlayersWeeklyProjection{}

	for _, pos := range positions {
		response := ffnGetRequest(serviceWeeklyProjections, formatJson, pos, week)
		playerMap := map[string]interface{}{}

		err := json.Unmarshal(response, &playerMap)
		if err != nil {
			log.Error().Err(err).Str("Function", "Teams").Msg("Failed to marshal HTTP response")
			return nil
		}

		for key, value := range playerMap {
			if key == "Projections" {
				for _, data := range value.([]interface{}) {
					stats := data.(map[string]interface{})
					players[pos] = append(players[pos], nflPlayerWeeklyProjection{
						week:          stats["week"].(string),
						playerId:      stats["playerId"].(string),
						position:      stats["position"].(string),
						passAtt:       stats["passAtt"].(string),
						passCmp:       stats["passCmp"].(string),
						passYds:       stats["passYds"].(string),
						passTD:        stats["passTD"].(string),
						passInt:       stats["passInt"].(string),
						rushAtt:       stats["rushAtt"].(string),
						rushYds:       stats["rushYds"].(string),
						rushTD:        stats["rushTD"].(string),
						fumblesLost:   stats["fumblesLost"].(string),
						receptions:    stats["receptions"].(string),
						recYds:        stats["recYds"].(string),
						recTD:         stats["recTD"].(string),
						fg:            stats["fg"].(string),
						fgAtt:         stats["fgAtt"].(string),
						xp:            stats["xp"].(string),
						defInt:        stats["defInt"].(string),
						defFR:         stats["defFR"].(string),
						defFF:         stats["defFF"].(string),
						defSack:       stats["defSack"].(string),
						defTD:         stats["defTD"].(string),
						defRetTD:      stats["defRetTD"].(string),
						defSafety:     stats["defSafety"].(string),
						defPA:         stats["defPA"].(string),
						defYdsAllowed: stats["defYdsAllowed"].(string),
						displayName:   stats["displayName"].(string),
						team:          stats["team"].(string),
					})

				}
			}
		}
	}
	return players
}

func (w Wideout) WeeklyRankings(week string) nflPlayersWeeklyRanking {
	players := nflPlayersWeeklyRanking{}

	for _, pos := range positions {
		response := ffnGetRequest(serviceWeeklyRankings, formatJson, pos, week)
		playerMap := map[string]interface{}{}

		err := json.Unmarshal(response, &playerMap)
		if err != nil {
			log.Error().Err(err).Str("Function", "Teams").Msg("Failed to marshal HTTP response")
			return nil
		}

		for key, value := range playerMap {
			if key == "Rankings" {
				for _, data := range value.([]interface{}) {
					stats := data.(map[string]interface{})
					players[pos] = append(players[pos], nflPlayerWeeklyRanking{
						week:          stats["week"].(string),
						playerId:      stats["playerId"].(string),
						position:      stats["position"].(string),
						name:       stats["name"].(string),
						team:       stats["team"].(string),
						standard:       stats["standard"].(string),
						standardLow:        stats["standardLow"].(string),
						standardHigh:       stats["standardHigh"].(string),
						ppr:       stats["ppr"].(string),
						pprLow:       stats["pprLow"].(string),
						pprHigh:        stats["pprHigh"].(string),

					})

				}
			}
		}
	}

	return players
}