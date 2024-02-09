package main

import (
	"fmt"
	"sap-1/sap1"
	"sap-1/scr"
	"time"
)

var steps int

func main() {

	getUserSteps()

	comp := sap1.NewSap1()
	comp.Cntl.Start()
	comp.Clk.StartClk()

	waitAbit(steps)

}

func getUserSteps() {
	fmt.Print(scr.ClearScreen + scr.CursorHome)
	fmt.Print("Enter number of steps: ")
	fmt.Scanln(&steps)
}

func waitAbit(steps int) {
	time.Sleep(
		time.Duration(steps)*
			time.Duration(1000)*
			time.Millisecond +
			3*time.Millisecond,
	)
}
