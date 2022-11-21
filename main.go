package main

import (
	"bufio"
	"fmt"
	"os"
	"stateengine/vendingmachine"
	"strconv"
	"strings"
)

var stateEng vendingmachine.VendingMachine

func main() {
	stateEng = vendingmachine.NewStateEngine(dispenseDrink, coinRejected)
	stateEng.PerformAction(vendingmachine.AddBottle)
	stateEng.PerformAction(vendingmachine.AddBottle)
	stateEng.PerformAction(vendingmachine.AddBottle)

	for {
		showStatus()
		a := askAction()
		if a < 0 {
			return
		}
		fmt.Printf("Action: %s\n", a)
		stateEng.PerformAction(a)
		fmt.Println()
	}
}

func dispenseDrink() {
	fmt.Printf("Dispense Drink\n")
}
func coinRejected() {
	fmt.Printf("Coin rejected!! No bottles left\n")
}

func showStatus() {
	fmt.Printf("Current State: %s\t", stateEng.State())
	fmt.Printf("Bottles: %d\t", stateEng.Bottles())
	fmt.Printf("Credit: %d\n", stateEng.Credit())
	if stateEng.Bottles() == 0 {
		fmt.Println("Warning: Machine Empty")
	}
}

func askAction() vendingmachine.Action {
	for i, a := range vendingmachine.AllActions {
		fmt.Printf("%d - %s\n", i+1, a)
	}
	fmt.Printf("0 - Exit\n")
	fmt.Print("Select action:")
	return readInput(len(vendingmachine.AllActions))
}

func readInput(max int) vendingmachine.Action {
	for {
		buf := bufio.NewReader(os.Stdin)
		line, _ := buf.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if line == "" {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if i < 0 || i > max {
			continue
		}
		return vendingmachine.Action(i - 1)
	}
}
