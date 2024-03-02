package entities

type Config struct {
	PlayerCount  int    `yaml:"player_count"`
	BoardSize    int    `yaml:"board_size"`
	SnakeCount   int    `yaml:"snake_count"`
	LadderCount  int    `yaml:"ladder_count"`
	DieCount     int    `yaml:"die_count"`
	MoveStrategy string `yaml:"move_strategy"`
}
