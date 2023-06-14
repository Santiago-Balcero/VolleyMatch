package services

import (
	"fmt"
	"math"
	"strings"
	"volleygame/constants"
	"volleygame/utils"
)

// Set struct and methods
type Set struct {
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
	lastAction          []string
	forward             bool
	SetCount            int
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

func (s *Set) AttackNeutral(fwd bool) {
	if fwd {
		s.AttackNeutrals += 1
	} else if !fwd {
		s.AttackNeutrals -= 1
	}
}

func (s *Set) AttackError(fwd bool) {
	if fwd {
		s.AttackErrors += 1
		s.OpponentPoints += 1
	} else if !fwd {
		s.AttackErrors -= 1
		s.OpponentPoints -= 1
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

func (s *Set) BlockNeutral(fwd bool) {
	if fwd {
		s.BlockNeutrals += 1
		s.TotalBlocks += 1
	} else {
		s.BlockNeutrals -= 1
		s.TotalBlocks -= 1
	}
}

func (s *Set) BlockError(fwd bool) {
	if fwd {
		s.BlockErrors += 1
		s.OpponentAttacks += 1
		s.OpponentPoints += 1
	} else {
		s.BlockErrors -= 1
		s.OpponentAttacks -= 1
		s.OpponentPoints -= 1
	}
}

func (s *Set) OpponentBlock(fwd bool) {
	if fwd {
		s.OpponentBlocks += 1
		s.OpponentPoints += 1
		s.AttackNeutrals += 1
	} else {
		s.OpponentBlocks -= 1
		s.OpponentPoints -= 1
		s.AttackNeutrals -= 1
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

func (s *Set) ServeNeutral(fwd bool) {
	if fwd {
		s.ServeNeutrals += 1
	} else {
		s.ServeNeutrals -= 1
	}
}

func (s *Set) ServeError(fwd bool) {
	if fwd {
		s.ServeErrors += 1
		s.OpponentPoints += 1
	} else {
		s.ServeErrors -= 1
		s.OpponentPoints -= 1
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

func (s *Set) UpdateStats() {
	// When doing rollback until set starting point (all stats in 0) effectiveness values are NaN
	// Ifs in this method fix it
	s.TotalAttacks = s.Attacks + s.AttackNeutrals + s.AttackErrors
	s.AttackEffectiveness = (float64(s.Attacks) / float64(s.TotalAttacks)) * 100
	if math.IsNaN(s.AttackEffectiveness) {
		s.AttackEffectiveness = 0.00
	}
	s.TotalBlocks = s.Blocks + s.BlockNeutrals + s.BlockErrors
	s.BlockEffectiveness = (float64(s.Blocks) / float64(s.TotalBlocks)) * 100
	if math.IsNaN(s.BlockEffectiveness) {
		s.BlockEffectiveness = 0.00
	}
	s.TotalServes = s.Serves + s.ServeNeutrals + s.ServeErrors
	s.ServeEffectiveness = (float64(s.Serves) / float64(s.TotalServes)) * 100
	if math.IsNaN(s.ServeEffectiveness) {
		s.ServeEffectiveness = 0.00
	}
	s.TotalActions = s.TotalAttacks + s.TotalBlocks + s.TotalServes + s.Errors
	s.TotalEffectiveness = (float64(s.Points-s.OpponentErrors) / float64(s.TotalActions)) * 100
	if math.IsNaN(s.TotalEffectiveness) {
		s.TotalEffectiveness = 0.00
	}
}

func (s *Set) PrintSet(team, opponent string) string {
	maxLen := 26
	var ptSpace int
	var atSpace int
	var blSpace int
	var seSpace int
	var erSpace int
	var toSpace int
	if s.Points < 10 {
		ptSpace = maxLen - 14
	} else {
		ptSpace = maxLen - 15
	}
	if s.Attacks < 10 {
		atSpace = maxLen - 15
	} else {
		atSpace = maxLen - 16
	}
	if s.Blocks < 10 {
		blSpace = maxLen - 14
	} else {
		blSpace = maxLen - 15
	}
	if s.Serves < 10 {
		seSpace = maxLen - 14
	} else {
		seSpace = maxLen - 15
	}
	if s.OpponentErrors < 10 {
		erSpace = 3
	} else {
		erSpace = 2
	}
	teamStats := fmt.Sprintf(
		"\n---- SET %d TEAM STATS ----\n%s%s - %s\nPoints:%s%d - %d\nAttacks:%s%d - %d\nBlocks:%s%d - %d\nServes:%s%d - %d\nOpponent errors:%s%d - %d\n",
		s.SetCount,
		strings.Repeat(" ", maxLen-9),
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
		s.Errors+s.AttackErrors+s.ServeErrors,
	)
	if s.AttackEffectiveness == 100 {
		atSpace = maxLen - 15
	} else if s.AttackEffectiveness >= 10 && s.AttackEffectiveness < 100 {
		atSpace = maxLen - 14
	} else if s.AttackEffectiveness < 10 {
		atSpace = maxLen - 13
	}
	if s.BlockEffectiveness == 100 {
		blSpace = maxLen - 14
	} else if s.BlockEffectiveness >= 10 && s.BlockEffectiveness < 100 {
		blSpace = maxLen - 13
	} else if s.BlockEffectiveness < 10 {
		blSpace = maxLen - 12
	}
	if s.ServeEffectiveness == 100 {
		seSpace = maxLen - 14
	} else if s.ServeEffectiveness >= 10 && s.ServeEffectiveness < 100 {
		seSpace = maxLen - 13
	} else if s.ServeEffectiveness < 10 {
		seSpace = maxLen - 12
	}
	if s.TotalEffectiveness == 100 {
		toSpace = maxLen - 14
	} else if s.TotalEffectiveness >= 10 && s.TotalEffectiveness < 100 {
		toSpace = maxLen - 13
	} else if s.TotalEffectiveness < 10 {
		toSpace = maxLen - 12
	}
	teamPerformance := fmt.Sprintf(
		"\n- TEAM SET %d EFFECTIVENESS\nAttack: %s%.2f%%\nBlock: %s%.2f%%\nServe: %s%.2f%%\nTotal: %s%.2f%%\n",
		s.SetCount,
		strings.Repeat(" ", atSpace),
		math.Round(s.AttackEffectiveness*100)/100,
		strings.Repeat(" ", blSpace),
		math.Round(s.BlockEffectiveness*100)/100,
		strings.Repeat(" ", seSpace),
		math.Round(s.ServeEffectiveness*100)/100,
		strings.Repeat(" ", toSpace),
		math.Round(s.TotalEffectiveness*100)/100,
	)
	return teamStats + teamPerformance
}

func (s *Set) PlaySet(team, opponent string) {
	fmt.Println()
	fmt.Println("New set!")
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
			if len(s.lastAction) > 0 {
				s.forward = false
				choice = s.lastAction[len(s.lastAction)-1]
				s.lastAction = s.lastAction[:len(s.lastAction)-1]
			} else {
				fmt.Println("No registered actions in set")
			}
		} else {
			s.lastAction = append(s.lastAction, choice)
		}
		switch choice {
		case constants.Attack:
			s.Attack(s.forward)
		case constants.AttackNeutral:
			s.AttackNeutral(s.forward)
		case constants.AttackError:
			s.AttackError(s.forward)
		case constants.OpponentAttack:
			s.OpponentAttack(s.forward)
		case constants.Block:
			s.Block(s.forward)
		case constants.BlockNeutral:
			s.BlockNeutral(s.forward)
		case constants.BlockError:
			s.BlockError(s.forward)
		case constants.OpponentBlock:
			s.OpponentBlock(s.forward)
		case constants.Serve:
			s.Serve(s.forward)
		case constants.ServeNeutral:
			s.ServeNeutral(s.forward)
		case constants.ServeError:
			s.ServeError(s.forward)
		case constants.OpponentService:
			s.OpponentServe(s.forward)
		case constants.Error:
			s.Error(s.forward)
		case constants.OpponentError:
			s.OpponentError(s.forward)
		}
		s.UpdateStats()
	}
	if s.Points > s.OpponentPoints {
		s.Winner = team
	} else {
		s.Winner = opponent
	}

}
