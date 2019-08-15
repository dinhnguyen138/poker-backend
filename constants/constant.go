package constants

const JOIN string = "JOIN"
const LEAVE string = "LEAVE"
const DEAL string = "DEAL"
const RAISE string = "RAISE"
const FOLD string = "FOLD"
const ALLIN string = "ALLIN"
const DOUBLE string = "DOUBLE"
const FOLLOW string = "FOLLOW"

const PLAYERS string = "PLAYERS"
const NEWPLAYER string = "NEWPLAYER"
const ELIMINATED string = "ELIMINATED"
const START string = "START"
const SHOWBACK string = "SHOWBACK"

const MAXPLAYER int = 4

var CardOrder map[string]int = map[string]int{
	"7":  0,
	"8":  1,
	"9":  2,
	"10": 3,
	"J":  4,
	"Q":  5,
	"K":  6,
	"A":  7,
}

var Types = []string{"7", "8", "9", "10", "J", "Q", "K", "A"}

var Suits = []string{"H", "D", "C", "S"}
