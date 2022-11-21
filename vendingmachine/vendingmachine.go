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
		// Only accept if there are the bottles available
		if s.coins < s.bottles {
			s.coins++
		} else {
			// out of stock, reject and remain in same state
			s.rejectFunc()
			ns = s.state
		}
	case PushButton:
		if s.state == Unlocked {
			s.coins--
			s.bottles--
			// If coins still left, remain in unlocked
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
