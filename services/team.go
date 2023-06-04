package services

import (
	"fmt"
	"strings"
)

// Team struct and method
type Team struct {
	Name string
}

func CreateTeam() Team {
	var teamName string
	for len(teamName) != 3 {
		fmt.Println("\nEnter team's name: ")
		fmt.Scan(&teamName)
		teamName = strings.TrimSpace(teamName)
		teamName = strings.ToUpper(teamName)
	}
	return Team{Name: teamName}
}

func (t *Team) StartGame() Game {
	var opponent string
	for len(opponent) != 3 || opponent == t.Name {
		fmt.Println("\nEnter opponent team's name: ")
		fmt.Scan(&opponent)
		opponent = strings.TrimSpace(opponent)
		opponent = strings.ToUpper(opponent)
	}
	return Game{
		Team:     t.Name,
		Opponent: opponent,
	}
}
