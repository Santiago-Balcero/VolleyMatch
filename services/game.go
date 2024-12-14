package services

import (
	"fmt"
	"math"
	"strings"
)

// Game struct and methods
type Game struct {
	Team                string
	Opponent            string
	TeamSets            int
	OpponentSets        int
	Sets                []Set
	Attacks             int
	AttackNeutrals      int
	AttackErrors        int
	TotalAttacks        int
	AttackEffectiveness float64
	Blocks              int
	BlockNeutrals       int
	BlockErrors         int
	TotalBlocks         int
	BlockEffectiveness  float64
	Serves              int
	ServeNeutrals       int
	ServeErrors         int
	TotalServes         int
	ServeEffectiveness  float64
	OpponentErrors      int
	Points              int
	TotalActions        int
	TotalEffectiveness  float64
	OpponentAttacks     int
	OpponentBlocks      int
	OpponentServes      int
	Errors              int
	OpponentPoints      int
	Winner              string
}

func (g *Game) PrintGame() string {
	maxLen := 26
	var stSpace int = maxLen - 12
	var ptSpace int
	var atSpace int
	var blSpace int
	var seSpace int
	var erSpace int
	var toSpace int
	if g.Points < 10 {
		ptSpace = maxLen - 20
	} else {
		ptSpace = maxLen - 21
	}
	if g.Attacks < 10 {
		atSpace = maxLen - 15
	} else {
		atSpace = maxLen - 16
	}
	if g.Blocks < 10 {
		blSpace = maxLen - 14
	} else {
		blSpace = maxLen - 15
	}
	if g.Serves < 10 {
		seSpace = maxLen - 14
	} else {
		seSpace = maxLen - 15
	}
	if g.OpponentErrors < 10 {
		erSpace = 3
	} else {
		erSpace = 2
	}
	teamStats := fmt.Sprintf(
		"\n---- MATCH TEAM STATS ----\n%s%s - %s\nSets:%s%d - %d\nTotal points:%s%d - %d\nAttacks:%s%d - %d\nBlocks:%s%d - %d\nServes:%s%d - %d\nOpponent errors:%s%d - %d\n",
		strings.Repeat(" ", maxLen-9),
		g.Team,
		g.Opponent,
		strings.Repeat(" ", stSpace),
		g.TeamSets,
		g.OpponentSets,
		strings.Repeat(" ", ptSpace),
		g.Points,
		g.OpponentPoints,
		strings.Repeat(" ", atSpace),
		g.Attacks,
		g.OpponentAttacks,
		strings.Repeat(" ", blSpace),
		g.Blocks,
		g.OpponentBlocks,
		strings.Repeat(" ", seSpace),
		g.Serves,
		g.OpponentServes,
		strings.Repeat(" ", erSpace),
		g.OpponentErrors,
		g.Errors+g.AttackErrors+g.ServeErrors,
	)
	if g.AttackEffectiveness == 100 {
		atSpace = maxLen - 15
	} else if g.AttackEffectiveness >= 10 && g.AttackEffectiveness < 100 {
		atSpace = maxLen - 14
	} else if g.AttackEffectiveness < 10 {
		atSpace = maxLen - 13
	}
	if g.BlockEffectiveness == 100 {
		blSpace = maxLen - 14
	} else if g.BlockEffectiveness >= 10 && g.BlockEffectiveness < 100 {
		blSpace = maxLen - 13
	} else if g.BlockEffectiveness < 10 {
		blSpace = maxLen - 12
	}
	if g.ServeEffectiveness == 100 {
		seSpace = maxLen - 14
	} else if g.ServeEffectiveness >= 10 && g.ServeEffectiveness < 100 {
		seSpace = maxLen - 13
	} else if g.ServeEffectiveness < 10 {
		seSpace = maxLen - 12
	}
	if g.TotalEffectiveness == 100 {
		toSpace = maxLen - 14
	} else if g.TotalEffectiveness >= 10 && g.TotalEffectiveness < 100 {
		toSpace = maxLen - 13
	} else if g.TotalEffectiveness < 10 {
		toSpace = maxLen - 12
	}
	teamPerformance := fmt.Sprintf(
		"\n- TEAM MATCH EFFECTIVENESS\nAttack: %s%.2f%%\nBlock: %s%.2f%%\nServe: %s%.2f%%\nTotal: %s%.2f%%\n\nTotal attacks: %d\nTotal blocks: %d\nTotal serves: %d\n",
		strings.Repeat(" ", atSpace),
		math.Round(g.AttackEffectiveness*100)/100,
		strings.Repeat(" ", blSpace),
		math.Round(g.BlockEffectiveness*100)/100,
		strings.Repeat(" ", seSpace),
		math.Round(g.ServeEffectiveness*100)/100,
		strings.Repeat(" ", toSpace),
		math.Round(g.TotalEffectiveness*100)/100,
		g.TotalAttacks,
		g.TotalBlocks,
		g.TotalServes,
	)
	return teamStats + teamPerformance
}

func (g *Game) Play() {
	var setCount int
	for {
		setCount++
		var choice string
		for choice != "y" && choice != "n" {
			fmt.Print("Play set? [y/n]: ")
			fmt.Scan(&choice)
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
		}
		if choice == "y" {
			newSet := Set{SetCount: setCount}
			newSet.PlaySet(g.Team, g.Opponent)
			g.UpdateGame(newSet)
		} else if choice == "n" {
			break
		}
	}
	// When end game check for team with more sets
	if g.TeamSets > g.OpponentSets {
		g.Winner = g.Team
	} else {
		g.Winner = g.Opponent
	}
	fmt.Println(g.PrintGame())
}

func (g *Game) UpdateGame(set Set) {
	if set.Winner == g.Team {
		g.TeamSets += 1
	} else {
		g.OpponentSets += 1
	}
	g.Sets = append(g.Sets, set)
	g.Attacks += set.Attacks
	g.AttackNeutrals += set.AttackNeutrals
	g.AttackErrors += set.AttackErrors
	g.TotalAttacks += set.TotalAttacks
	g.Blocks += set.Blocks
	g.BlockNeutrals += set.BlockNeutrals
	g.BlockErrors += set.BlockErrors
	g.TotalBlocks += set.TotalBlocks
	g.Serves += set.Serves
	g.ServeNeutrals += set.ServeNeutrals
	g.ServeErrors += set.ServeErrors
	g.TotalServes += set.TotalServes
	g.OpponentErrors += set.OpponentErrors
	g.Points += set.Points
	g.OpponentAttacks += set.OpponentAttacks
	g.OpponentBlocks += set.OpponentBlocks
	g.OpponentServes += set.OpponentServes
	g.Errors += set.Errors
	g.OpponentPoints += set.OpponentPoints
	g.TotalActions += set.TotalActions
	g.AttackEffectiveness = (float64(g.Attacks) / float64(g.TotalAttacks)) * 100
	if math.IsNaN(g.AttackEffectiveness) {
		g.AttackEffectiveness = 0.00
	}
	g.TotalBlocks = g.Blocks + g.BlockNeutrals + g.BlockErrors
	g.BlockEffectiveness = (float64(g.Blocks) / float64(g.TotalBlocks)) * 100
	if math.IsNaN(g.BlockEffectiveness) {
		g.BlockEffectiveness = 0.00
	}
	g.TotalServes = g.Serves + g.ServeNeutrals + g.ServeErrors
	g.ServeEffectiveness = (float64(g.Serves) / float64(g.TotalServes)) * 100
	if math.IsNaN(g.ServeEffectiveness) {
		g.ServeEffectiveness = 0.00
	}
	g.TotalEffectiveness = (float64(g.Points-g.OpponentErrors) / float64(g.TotalActions)) * 100
	if math.IsNaN(g.TotalEffectiveness) {
		g.TotalEffectiveness = 0.00
	}
}
