package vendingmachine

import "log"

const (
	Empty State = iota
	Locked
	Unlocked
)

// The finite states the engine can hold
type State int

// actionMaps represents which map is used by which state.
var actionMaps = map[State]map[Action]State{
	Empty:    actionMapEmpty,
	Locked:   actionMapLocked,
	Unlocked: actionMapUnlocked,
}

// actionMap Maps each eaction to its respective State
var actionMapUnlocked = map[Action]State{
	PushButton: Locked,
	InsertCoin: Unlocked,
	AddBottle:  Unlocked,
}
var actionMapLocked = map[Action]State{
	PushButton: Locked,
	InsertCoin: Unlocked,
	AddBottle:  Locked,
}

var actionMapEmpty = map[Action]State{
	PushButton: Empty,
	InsertCoin: Empty,
	AddBottle:  Locked,
}

// NextState will return the next state based on the given action
func (s State) NextState(action Action) State {
	nm := actionMaps[s]
	ns, ok := nm[action]
	if !ok {
		log.Printf("%s is not a valid action in this state (%s)", action, s)
		return s
	}
	return ns
}

// helps to display states and actions
var stateNames = [...]string{"Empty", "Locked", "Unlocked"}

func (s State) String() string {
	return stateNames[s]
}
