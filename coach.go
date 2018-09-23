package main

// coach.go coordinates the requested data based on user input

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/trashbo4t/sideline/sideline"
	"math"
	"sort"
)

type rank struct {
	name string
	score float64
}

type lineup struct {
	QB rank
	RB1 rank
	RB2 rank
	WR1 rank
	WR2 rank
	TE rank
	RWT rank
	K rank
	Def rank
}

type ByScore []rank
func (s ByScore) Len() int {return len(s)}
func (s ByScore) Less(i, j int) bool { return s[i].score > s[j].score }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func init() {
	apiKey := flag.String("api-key", "test", "ApiKey for the fantasyfootballnerds.com REST API")
	flag.Parse()
	sideline.APIKey(*apiKey)
}

func showOptions() {
	fmt.Println("Sideline: A fantasy Football engine written in Go")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Select an option:")
	fmt.Println("1: Predict game winners for this week")
	fmt.Println("---------------------------------------------------")
}

func main() {

	getInput:
	for {
		showOptions()

		input := ""
		fmt.Scanln(&input)

		switch input {
		case "1":
			predictWinners()
			continue getInput
		default:
			fmt.Println("Please select a valid option...")
			continue getInput
		}
	}
	zerolog.TimeFieldFormat = ""
	/*	lineups := top5Teams()

		for i, lineup := range lineups {
			println("Team", i+1)
			fmt.Println("QB:", lineup.QB)
			fmt.Println("WR1:", lineup.WR1)
			fmt.Println("WR2:", lineup.WR2)
			fmt.Println("RB1:", lineup.RB1)
			fmt.Println("RB2:", lineup.RB2)
			fmt.Println("TE:", lineup.TE)
			fmt.Println("K:", lineup.K)
			fmt.Println("DEF:", lineup.Def)
			println("------------------------")
		}
	*/

}

func predictWinners() {
	ranks := rankTeams()
	rankMap := map[string]float64{}

	for _, team := range ranks {
		rankMap[team.name] = team.score
	}
	winners := []string{}

	w := sideline.Wideout{}
	games, currentWeek := w.Games()
	for _, game := range games {
		fmt.Println(game.Home(), "vs", game.Away())

		fmt.Println("Predicted winners for week", currentWeek)
		fmt.Println(game.Home(),"->", rankMap[game.Home()])
		fmt.Println(game.Away(),"->", rankMap[game.Away()])

		fmt.Println("Differential")
		if rankMap[game.Home()] > rankMap[game.Away()] {
			fmt.Println(rankMap[game.Home()]-rankMap[game.Away()])
			winners = append(winners, game.Home())
		} else {
			fmt.Println(rankMap[game.Away()]-rankMap[game.Home()])
			winners = append(winners, game.Away())
		}

		fmt.Println("Differential (%)")
		fmt.Println((math.Abs(rankMap[game.Home()]-rankMap[game.Away()])/((rankMap[game.Away()]+rankMap[game.Home()])/2)) * 100)

		winnerMap := map[bool]string{true: game.Home(), false: game.Away()}
		fmt.Println("Predicted Winner:", winnerMap[rankMap[game.Home()] > rankMap[game.Away()]])
		fmt.Println("-------------")
	}
	fmt.Println("Condensed winners for week", currentWeek)
	fmt.Println(winners, "\n")
}

func rankTeams() []rank {
	ranks := []rank{}

	w := sideline.Wideout{}

	weeklyProjections := w.WeeklyProjections("3")
	topPos := map[string]float64{}

	for _, players := range weeklyProjections {
		for _, player := range players {
			topPos[player.Team()] += player.DetermineProjectionScore()
		}
	}

	for team, score := range topPos {
		ranks = append(ranks, rank{name: team, score: score})
	}
	sort.Sort(ByScore(ranks))

	return ranks
}

func top5Teams() []lineup {
	w := sideline.Wideout{}

	weeklyProjections := w.WeeklyProjections("3")
	topPos := map[string]ByScore{}

	for pos, players := range weeklyProjections {
		top := ByScore{}
		for _, player := range players {
			top = append(top, rank{name: player.Name(), score: player.DetermineProjectionScore()})
		}
		sort.Sort(ByScore(top))
		topPos[pos] = top[:20]
	}

	// build 5 lineups based on top players
	lineups := make([]lineup, 5)
	j := 1
	for i := range lineups {
		lineups[i].QB = topPos["QB"][i]
		lineups[i].TE = topPos["TE"][i]
		lineups[i].K = topPos["K"][i]
		lineups[i].Def = topPos["DEF"][i]
		lineups[i].RB1 = topPos["RB"][i]
		lineups[i].RB2 = topPos["RB"][j]
		lineups[i].WR1 = topPos["WR"][i]
		lineups[i].WR2 = topPos["WR"][j]
		j+=i+2
	}

	return lineups
}


