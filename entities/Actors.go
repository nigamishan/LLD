package entities

import (
	"log"
	"math"
)

type EndpointMap map[int]int

type Actor struct {
	Name      string
	StartsMap EndpointMap
	EndsMap   EndpointMap
}

type Snake struct {
	Start, End int
}

type Ladder struct {
	Start, End int
}

func CreateActor(start, end int, actorType string, board *Board) bool {
	switch actorType {
	case "Snake":
		snake := Snake{start, end}
		if snake.IsValidActor(*board) {
			actor, found := board.Actors["Snake"]
			if !found {
				actor.Name = "Snake"
				actor.EndsMap = make(map[int]int)
				actor.StartsMap = make(map[int]int)
			}
			actor.StartsMap[start] = end
			actor.EndsMap[end] = start
			(*board).Actors["Snake"] = actor
			return true
		} else {
			return false
		}
	case "Ladder":
		ladder := Ladder{start, end}
		if ladder.IsValidActor(*board) {
			actor, found := board.Actors["Ladder"]
			if !found {
				actor.Name = "Ladder"
				actor.EndsMap = make(map[int]int)
				actor.StartsMap = make(map[int]int)
			}
			actor.StartsMap[start] = end
			actor.EndsMap[end] = start
			(*board).Actors["Ladder"] = actor
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func GeneralValidation(start, end int, endpoints map[string]Actor) bool {
	for _, actor := range endpoints {
		// start point of any existing should not match our current end point
		if _, found := actor.StartsMap[end]; found {
			log.Println("start point of any existing actor should not match our current end point")
			return false
		}
		// end point should not match any start point
		if _, found := actor.EndsMap[start]; found {
			log.Println("end point should not match any start point")
			return false
		}
		// start point should not match any other start point
		if _, found := actor.StartsMap[start]; found {
			log.Println("start point should not match any other start point")
			return false
		}
	}
	return true
}

func (self *Snake) IsValidActor(board Board) bool {
	if self.Start <= self.End {
		return false
	}

	if _, found := board.Actors["Snake"]; !found {
		log.Println("No snake found. Adding a new one")
	}
	if !GeneralValidation(self.Start, self.End, board.Actors) {
		return false
	}

	return true
}

func (self *Ladder) IsValidActor(board Board) bool {
	boardLen := int(math.Sqrt(float64(board.Size)))
	// Should be strictly an upper cell
	if self.Start/boardLen >= self.End/boardLen {
		return false
	}
	if _, found := board.Actors["Ladder"]; !found {
		log.Println("No Ladder found. Adding a new one")
	}
	if !GeneralValidation(self.Start, self.End, board.Actors) {
		return false
	}
	return true
}
