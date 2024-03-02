package entities

import (
	"fmt"
	"log"
)

type BoardInterface interface {
	Move(int) int
	AddActors(start, end int, actorType string) bool
	HasWon(currentPosition int) bool
}

type Board struct {
	ID     int
	Size   int
	Actors map[string]Actor
	Dice   Dice
}

func CreateBoard(size int, diceCount int, moveStrategy string) *Board {
	board := &Board{
		Size:   size,
		ID:     1,
		Actors: make(map[string]Actor),
		Dice:   setMoveStrategy(diceCount, moveStrategy),
	}
	return board
}

func setMoveStrategy(diceCount int, moveStrategy string) Dice {
	switch moveStrategy {
	case "MAX":
		return &MAX{diceCount}
	case "MIN":
		return &MIN{diceCount}
	case "SUM":
		return &SUM{diceCount}
	default:
		return nil
	}
}

func (self *Board) HasWon(currentPosition int) bool {
	return self.Size == currentPosition
}
func (self *Board) Move(currentPosition int) int {
	steps := self.Dice.Roll()
	endPos := currentPosition + steps
	fmt.Printf("Attempting to Move %d steps from position %d\n", steps, currentPosition)

	if !self.isValidMove(endPos) {
		return currentPosition
	}
	for actorName, actor := range self.Actors {
		if newEndpos, found := actor.StartsMap[endPos]; found {
			log.Printf("%s is encountered. Moving to %d from %d", actorName, newEndpos, currentPosition)
			return newEndpos
		}
	}
	return endPos
}

func (self *Board) isValidMove(position int) bool {
	return position <= self.Size
}

func (self *Board) AddActors(start, end int, actorType string) bool {
	return CreateActor(start, end, actorType, self)
}
