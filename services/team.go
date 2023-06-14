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
	fmt.Println()
	for {
		for len(teamName) != 3 {
			fmt.Print("Enter team's name: ")
			fmt.Scan(&teamName)
			teamName = strings.TrimSpace(teamName)
			teamName = strings.ToUpper(teamName)
		}
		var confirm string
		for confirm != "y" && confirm != "n" {
			fmt.Printf("Confirm name %s? [y/n]: ", teamName)
			fmt.Scan(&confirm)
		}
		if confirm == "y" {
			break
		} else if confirm == "n" {
			teamName = ""
			fmt.Println()
			continue
		}
	}
	return Team{Name: teamName}
}

func (t *Team) StartGame() Game {
	var opponent string
	fmt.Println()
	for {
		for len(opponent) != 3 || opponent == t.Name {
			fmt.Print("Enter opponent team's name: ")
			fmt.Scan(&opponent)
			opponent = strings.TrimSpace(opponent)
			opponent = strings.ToUpper(opponent)
		}
		var confirm string
		for confirm != "y" && confirm != "n" {
			fmt.Printf("Confirm opponent team's name %s? [y/n]: ", opponent)
			fmt.Scan(&confirm)
		}
		if confirm == "y" {
			break
		} else if confirm == "n" {
			opponent = ""
			fmt.Println()
			continue
		}
	}
	return Game{
		Team:     t.Name,
		Opponent: opponent,
	}
}
