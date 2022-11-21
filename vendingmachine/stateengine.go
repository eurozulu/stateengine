package vendingmachine

type StateEventFunc func()

// VendingMachine represents the current status as vending machine
type VendingMachine interface {
	State() State
	Credit() int
	Bottles() int
	PerformAction(action Action) State
}

type vendingMachine struct {
	state        State
	coins        int
	bottles      int
	dispenseFunc StateEventFunc
	rejectFunc   StateEventFunc
}

func (s vendingMachine) State() State {
	return s.state
}

func (s vendingMachine) Credit() int {
	return s.coins
}

func (s vendingMachine) Bottles() int {
	return s.bottles
}

func (s *vendingMachine) PerformAction(action Action) State {
	ns := s.state.NextState(action)
	switch action {
	case AddBottle:
		s.bottles++
	case InsertCoin:
		if s.coins < s.bottles {
			s.coins++
		} else {
			s.rejectFunc()
			ns = s.state
		}
	case PushButton:
		if s.state == Unlocked {
			s.coins--
			s.bottles--
			if s.coins > 0 {
				ns = Unlocked
			}
			if s.bottles < 1 {
				ns = Empty
			}
			s.dispenseFunc()
		}
	}
	s.state = ns
	return ns
}

func NewStateEngine(dispense, reject StateEventFunc) VendingMachine {
	return &vendingMachine{dispenseFunc: dispense, rejectFunc: reject}
}
