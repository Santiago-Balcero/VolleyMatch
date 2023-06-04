package services

import (
	"fmt"
	"strings"
	"volleygame/constants"
	"volleygame/utils"
)

// Set struct and methods
type Set struct {
	Attacks         int
	Blocks          int
	Serves          int
	OpponentErrors  int
	Points          int
	OpponentAttacks int
	OpponentBlocks  int
	OpponentServes  int
	Errors          int
	OpponentPoints  int
	Winner          string
	lastAction      string
	forward         bool
}

func (s *Set) Attack(fwd bool) {
	if fwd {
		s.Attacks += 1
		s.Points += 1
	} else if !fwd {
		s.Attacks -= 1
		s.Points -= 1
	}
}

func (s *Set) OpponentAttack(fwd bool) {
	if fwd {
		s.OpponentAttacks += 1
		s.OpponentPoints += 1
	} else {
		s.OpponentAttacks -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) Block(fwd bool) {
	if fwd {
		s.Blocks += 1
		s.Points += 1
	} else {
		s.Blocks -= 1
		s.Points -= 1
	}
}

func (s *Set) OpponentBlock(fwd bool) {
	if fwd {
		s.OpponentBlocks += 1
		s.OpponentPoints += 1
	} else {
		s.OpponentBlocks -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) Serve(fwd bool) {
	if fwd {
		s.Serves += 1
		s.Points += 1
	} else {
		s.Serves -= 1
		s.Points -= 1
	}
}

func (s *Set) OpponentServe(fwd bool) {
	if fwd {
		s.OpponentServes += 1
		s.OpponentPoints += 1
	} else {
		s.OpponentServes -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentError(fwd bool) {
	if fwd {
		s.OpponentErrors += 1
		s.Points += 1
	} else {
		s.OpponentErrors -= 1
		s.Points -= 1
	}
}

func (s *Set) Error(fwd bool) {
	if fwd {
		s.Errors += 1
		s.OpponentPoints += 1
	} else {
		s.Errors -= 1
		s.OpponentPoints -= 1
	}
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
	if s.Serves < 10 {
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
		s.Serves,
		s.OpponentServes,
		strings.Repeat(" ", erSpace),
		s.OpponentErrors,
		s.Errors,
	)
}

func (s *Set) PlaySet(team, opponent string) {
	fmt.Println()
	fmt.Println("NEW SET!")
	for {
		fmt.Println(s.PrintSet(team, opponent))
		var choice string
		s.forward = true
		for !utils.CheckStringInArray(choice, constants.SetActions) {
			fmt.Print("Game action: ")
			fmt.Scan(&choice)
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
		}
		if choice == constants.Exit {
			fmt.Println("\nEnd of set!")
			fmt.Println()
			break
		}
		if choice == constants.RollBack {
			s.forward = false
			choice = s.lastAction
		} else {
			s.lastAction = choice
		}
		switch choice {
		case constants.Attack:
			s.Attack(s.forward)
		case constants.OpponentAttack:
			s.OpponentAttack(s.forward)
		case constants.Block:
			s.Block(s.forward)
		case constants.OpponentBlock:
			s.OpponentBlock(s.forward)
		case constants.Service:
			s.Serve(s.forward)
		case constants.OpponentService:
			s.OpponentServe(s.forward)
		case constants.Error:
			s.Error(s.forward)
		case constants.OpponentError:
			s.OpponentError(s.forward)
		}
	}
	if s.Points > s.OpponentPoints {
		s.Winner = team
	} else {
		s.Winner = opponent
	}
}
