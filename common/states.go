package common

type PlayerState string
type GameState string

const (
	ACTIVE PlayerState = "ACTIVE"
	WON    PlayerState = "WON"
)

const (
	LIVE  GameState = "ACTIVE"
	ENDED GameState = "ENDED"
)
