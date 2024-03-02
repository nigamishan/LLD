package entities

import (
	"fmt"
	"log"
	"math/rand/v2"
)

type Dice interface {
	Roll() int
}

type SUM struct {
	Count int
}
type MIN struct {
	Count int
}
type MAX struct {
	Count int
}

func (self *SUM) Roll() int {
	var result, val int
	var roll string
	result = 0
	for {
		fmt.Println("Enter D to get Roll Dice")
		if _, err := fmt.Scanf("%s", &roll); err != nil {
			log.Println("Incorrect format")
			continue
		} else {
			if roll == "D" {
				for i := 0; i < self.Count; i++ {
					val = rand.IntN(6) + 1
					result += val
				}
				return result
			} else {
				log.Println("Incorrect Character")
				continue
			}
		}
	}
}

func (self *MAX) Roll() int {
	var result, val int
	var roll string
	result = 0
	for {
		fmt.Println("Enter D to get Roll Dice")
		if _, err := fmt.Scanf("%s", &roll); err != nil {
			log.Println("Incorrect format")
			continue
		} else {
			if roll == "D" {
				for i := 0; i < self.Count; i++ {
					val = rand.IntN(6) + 1
					if val > result {
						result = val
					}
				}
				return result
			} else {
				log.Println("Incorrect Character")
				continue
			}
		}
	}
}

func (self *MIN) Roll() int {
	var result, val int
	var roll string
	result = 0
	for {
		fmt.Println("Enter D to get Roll Dice")
		if _, err := fmt.Scanf("%s", &roll); err != nil {
			log.Println("Incorrect format")
			continue
		} else {
			if roll == "D" {
				for i := 0; i < self.Count; i++ {
					val = rand.IntN(6) + 1
					if val < result {
						result = val
					}
				}
				return result
			} else {
				log.Println("Incorrect Character")
				continue
			}
		}
	}
}
