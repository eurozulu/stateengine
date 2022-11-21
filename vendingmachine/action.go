package vendingmachine

const (
	PushButton Action = iota
	InsertCoin
	AddBottle
)

// The actions which can change states
type Action int

var AllActions = [...]Action{PushButton, InsertCoin, AddBottle}
var actionNames = [...]string{"Push Button", "Inset Coin", "Add Bottle"}

func (a Action) String() string {
	return actionNames[a]
}
