package Manager

import (
	"SnakesAndLadder/common"
	"SnakesAndLadder/entities"
	"log"
)

type Users []entities.User

type Game struct {
	Id           int
	PlayingUsers Users
	WonUsers     Users
	Board        entities.BoardInterface
	State        common.GameState
}

type GameInterface interface {
	StartGame()
}

func CreateGame(board *entities.Board) *Game {
	return &Game{
		Id:           1,
		PlayingUsers: make(Users, 0),
		WonUsers:     make(Users, 0),
		Board:        board,
		State:        common.LIVE,
	}
}

func (self *Game) AddPlayers(player entities.User) {
	self.PlayingUsers = append(self.PlayingUsers, player)
}
func (self *Game) DisplacePlayers(userName string, currentPos int) {
	for i, player := range self.PlayingUsers {
		if player.Name != userName && player.CurrentPosition == currentPos {
			log.Println("Displacing Player: ", player.Name)
			self.PlayingUsers[i].CurrentPosition = 1
		}
	}
}
func (self *Game) StartGame() {
	for self.State == common.LIVE {
		user := self.PlayingUsers.Dequeue()
		log.Println("Roll Dice for player: ", user.Name)
		user.CurrentPosition = self.Board.Move(user.CurrentPosition)
		self.DisplacePlayers(user.Name, user.CurrentPosition)
		log.Printf("Moved Player %s to position %d\n", user.Name, user.CurrentPosition)
		if self.Board.HasWon(user.CurrentPosition) {
			log.Printf("Player %s WON!\n", user.Name)
			user.State = common.WON
			self.WonUsers.Enqueue(user)
		} else {
			self.PlayingUsers.Enqueue(user)
		}
		log.Println("Here are Live Players", self.PlayingUsers)
		log.Println("Here are Winners", self.WonUsers)
		if len(self.PlayingUsers) == 0 {
			self.State = common.ENDED
		}

	}
}

func (self *Users) Enqueue(user entities.User) {
	*self = append(*self, user)
}

func (self *Users) Dequeue() (user entities.User) {
	user = (*self)[0]
	*self = (*self)[1:]
	return user
}
