package services

import (
	"fmt"
	"strings"
	"volleygame/constants"
	"volleygame/utils"
)

// Set struct and methods
type Set struct {
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

func (s *Set) Attack() {
	s.Attacks += 1
	s.Points += 1
}

func (s *Set) OpponentAttack() {
	s.OpponentAttacks += 1
	s.OpponentPoints += 1
}

func (s *Set) Block() {
	s.Blocks += 1
	s.Points += 1
}

func (s *Set) OpponentBlock() {
	s.OpponentBlocks += 1
	s.OpponentPoints += 1
}

func (s *Set) Service() {
	s.Services += 1
	s.Points += 1
}

func (s *Set) OpponentService() {
	s.OpponentServices += 1
	s.OpponentPoints += 1
}

func (s *Set) OpponentError() {
	s.OpponentErrors += 1
	s.Points += 1
}

func (s *Set) Error() {
	s.Errors += 1
	s.OpponentPoints += 1
}

func (s *Set) StartSet() {

}

func (s *Set) PrintSet(team, opponent string) string {
	maxLen := 17
	var ptSpace int
	var atSpace int
	var blSpace int
	var seSpace int
	var erSpace int
	if s.Points < 10 {
		ptSpace = maxLen - 5
	} else {
		ptSpace = maxLen - 6
	}
	if s.Attacks < 10 {
		atSpace = maxLen - 6
	} else {
		atSpace = maxLen - 7
	}
	if s.Blocks < 10 {
		blSpace = maxLen - 5
	} else {
		blSpace = maxLen - 6
	}
	if s.Services < 10 {
		seSpace = maxLen - 5
	} else {
		seSpace = maxLen - 6
	}
	if s.OpponentErrors < 10 {
		erSpace = 3
	} else {
		erSpace = 2
	}
	return fmt.Sprintf(
		"\n--- SET TEAM STATS ---\n%s%s - %s\nPoints:%s%d - %d\nAttacks:%s%d - %d\nBlocks:%s%d - %d\nServes:%s%d - %d\nOpponent errors:%s%d - %d\n",
		strings.Repeat(" ", maxLen),
		team,
		opponent,
		strings.Repeat(" ", ptSpace),
		s.Points,
		s.OpponentPoints,
		strings.Repeat(" ", atSpace),
		s.Attacks,
		s.OpponentAttacks,
		strings.Repeat(" ", blSpace),
		s.Blocks,
		s.OpponentBlocks,
		strings.Repeat(" ", seSpace),
		s.Services,
		s.OpponentServices,
		strings.Repeat(" ", erSpace),
		s.OpponentErrors,
		s.Errors,
	)
}

func (s *Set) PlaySet(team, opponent string) {
	fmt.Println()
	fmt.Println("NEW SET!")
	fmt.Println()
	for {
		fmt.Println(s.PrintSet(team, opponent))
		var choice string
		for !utils.CheckStringInArray(choice, constants.SetActions) {
			fmt.Println("Game action: ")
			fmt.Scan(&choice)
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
		}
		if choice == constants.Exit {
			fmt.Println("\nEnd of set!")
			fmt.Println()
			break
		}
		switch choice {
		case constants.Attack:
			s.Attack()
		case constants.OpponentAttack:
			s.OpponentAttack()
		case constants.Block:
			s.Block()
		case constants.OpponentBlock:
			s.OpponentBlock()
		case constants.Service:
			s.Service()
		case constants.OpponentService:
			s.OpponentService()
		case constants.Error:
			s.Error()
		case constants.OpponentError:
			s.OpponentError()
		}
	}
	if s.Points > s.OpponentPoints {
		s.Winner = team
	} else {
		s.Winner = opponent
	}
}
