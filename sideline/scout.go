package sideline

// scout.go defines all API data formatting as well as controls encapsulation for such formats

import (
	"github.com/rs/zerolog/log"
	"strconv"
)

var positions = []string{"QB", "RB", "WR", "TE", "K", "DEF"}

type nflTeams []nflTeam

type nflTeam struct {
	code string
	fullName string
	shortName string
}

type nflPlayersWeeklyProjection map[string][]nflPlayerWeeklyProjection

type nflPlayerWeeklyProjection struct {
	week	string
	playerId	string
	position	string
	passAtt	string
	passCmp	string
	passYds	string
	passTD	string
	passInt	string
	rushAtt	string
	rushYds	string
	rushTD	string
	fumblesLost	string
	receptions	string
	recYds	string
	recTD	string
	fg	string
	fgAtt	string
	xp	string
	defInt	string
	defFR	string
	defFF	string
	defSack	string
	defTD	string
	defRetTD	string
	defSafety	string
	defPA	string
	defYdsAllowed	string
	displayName	string
	team	string
}

type nflPlayersWeeklyRanking map[string][]nflPlayerWeeklyRanking

type nflPlayerWeeklyRanking struct {
	week	string
	playerId	string
	position	string
	name	string
	team	string
	standard	string
	standardLow	string
	standardHigh	string
	ppr	string
	pprLow	string
	pprHigh	string

}

type nflMatchups []nflMatchup

type nflMatchup struct {
	gameId	string
	gameWeek	string
	gameDate	string
	awayTeam	string
	homeTeam	string
	gameTimeET	string
	tvStation	string
	winner	string
}

func (matchup nflMatchup) Home() string {
	return matchup.homeTeam
}

func (matchup nflMatchup) Away() string {
	return matchup.awayTeam
}

func (player nflPlayerWeeklyProjection) Name() string {
	return player.displayName
}

func (player nflPlayerWeeklyProjection) Team() string {
	return player.team
}

func (player nflPlayerWeeklyProjection) DetermineProjectionScore() float64 {
	score := 0.0

	passAtt, err := strconv.ParseFloat(player.passAtt, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert passAtt")
		return 0.0
	}
	score += passAtt

	passCmp, err := strconv.ParseFloat(player.passCmp, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert passCmp")
		return 0.0
	}
	score += passCmp

	passYds, err := strconv.ParseFloat(player.passYds, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert passYds")
		return 0.0
	}
	score += passYds

	passTD, err := strconv.ParseFloat(player.passTD, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert passTD")
		return 0.0
	}
	score += passTD

	passInt, err := strconv.ParseFloat(player.passInt, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert passInt")
		return 0.0
	}
	score -= passInt

	rushAtt, err := strconv.ParseFloat(player.rushAtt, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert rushAtt")
		return 0.0
	}
	score += rushAtt

	rushYds, err := strconv.ParseFloat(player.rushYds, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert rushYds")
		return 0.0
	}
	score += rushYds

	rushTD, err := strconv.ParseFloat(player.rushTD, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert rushTD")
		return 0.0
	}
	score += rushTD

	fumblesLost, err := strconv.ParseFloat(player.fumblesLost, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert fumblesLost")
		return 0.0
	}
	score -= fumblesLost

	receptions, err := strconv.ParseFloat(player.receptions, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert receptions")
		return 0.0
	}
	score += receptions

	recYds, err := strconv.ParseFloat(player.recYds, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert recYds")
		return 0.0
	}
	score += recYds

	recTD, err := strconv.ParseFloat(player.recTD, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert recTD")
		return 0.0
	}
	score += recTD

	fg, err := strconv.ParseFloat(player.fg, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert fg")
		return 0.0
	}
	score += fg

	fgAtt, err := strconv.ParseFloat(player.fgAtt, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert fgAtt")
		return 0.0
	}
	score += fgAtt

	xp, err := strconv.ParseFloat(player.xp, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert xp")
		return 0.0
	}
	score += xp

	defInt, err := strconv.ParseFloat(player.defInt, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defInt")
		return 0.0
	}
	score += defInt

	defFR, err := strconv.ParseFloat(player.defFR, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defFR")
		return 0.0
	}
	score += defFR

	defFF, err := strconv.ParseFloat(player.defFF, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defFF")
		return 0.0
	}
	score += defFF

	defSack, err := strconv.ParseFloat(player.defSack, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defSack")
		return 0.0
	}
	score += defSack

	defTD, err := strconv.ParseFloat(player.defTD, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defTD")
		return 0.0
	}
	score += defTD

	defRetTD, err := strconv.ParseFloat(player.defRetTD, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defRetTD")
		return 0.0
	}
	score += defRetTD

	defSafety, err := strconv.ParseFloat(player.defSafety, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defSafety")
		return 0.0
	}
	score += defSafety

	defPA, err := strconv.ParseFloat(player.defPA, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defPA")
		return 0.0
	}
	score -= defPA

	defYdsAllowed, err := strconv.ParseFloat(player.defPA, 64)
	if err != nil {
		log.Error().Err(err).Str("Function", "getPlayers").Msg("Failed to convert defYdsAllowed")
		return 0.0
	}
	score -= defYdsAllowed

	return score
}