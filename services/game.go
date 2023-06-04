package services

import (
	"fmt"
	"strings"
)

// Game struct and methods
type Game struct {
	Team             string
	Opponent         string
	TeamSets         int
	OpponentSets     int
	Sets             []Set
	Attacks          int
	Blocks           int
	Services         int
	OpponentErrors   int
	Points           int
	OpponentAttacks  int
	OpponentBlocks   int
	OpponentServices int
	Errors           int
	OpponentPoints   int
	Winner           string
}

func (g *Game) PrintGame() string {
	maxLen := 17
	var stSpace int = maxLen - 3
	var ptSpace int
	var atSpace int
	var blSpace int
	var seSpace int
	var erSpace int
	if g.Points < 10 {
		ptSpace = maxLen - 11
	} else {
		ptSpace = maxLen - 12
	}
	if g.Attacks < 10 {
		atSpace = maxLen - 6
	} else {
		atSpace = maxLen - 7
	}
	if g.Blocks < 10 {
		blSpace = maxLen - 5
	} else {
		blSpace = maxLen - 6
	}
	if g.Services < 10 {
		seSpace = maxLen - 5
	} else {
		seSpace = maxLen - 6
	}
	if g.OpponentErrors < 10 {
		erSpace = 3
	} else {
		erSpace = 2
	}
	return fmt.Sprintf(
		"\n--- MATCH TEAM STATS ---\n%s%s - %s\nSets:%s%d - %d\nTotal points:%s%d - %d\nAttacks:%s%d - %d\nBlocks:%s%d - %d\nServes:%s%d - %d\nOpponent errors:%s%d - %d\n",
		strings.Repeat(" ", maxLen),
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
		g.Services,
		g.OpponentServices,
		strings.Repeat(" ", erSpace),
		g.OpponentErrors,
		g.Errors,
	)
}

func (g *Game) Play() {
	for {
		var choice string
		for choice != "y" && choice != "n" {
			fmt.Print("Play set? [y/n]: ")
			fmt.Scan(&choice)
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
		}
		if choice == "y" {
			newSet := Set{}
			newSet.PlaySet(g.Team, g.Opponent)
			g.UpdateGame((newSet))
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
	g.Blocks += set.Blocks
	g.Services += set.Services
	g.OpponentErrors += set.OpponentErrors
	g.Points += set.Points
	g.OpponentAttacks += set.OpponentAttacks
	g.OpponentBlocks += set.OpponentBlocks
	g.OpponentServices += set.OpponentServices
	g.Errors += set.Errors
	g.OpponentPoints += set.OpponentPoints
}
