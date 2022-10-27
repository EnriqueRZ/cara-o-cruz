package main

import (
	"cara-cruz/utils"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
);

func main() {
	random := float64(utils.RandIntn(1<<53)) / (1 << 53)
	fmt.Print("RANDOM = ", random);
	
	carasArray := utils.GetCaras()
	for _, s := range carasArray {
		fmt.Printf("%2s \r", s)
		time.Sleep(1 * time.Second/2)
		//if !(i == len(carasArray) - 1) {
		clearScreen()
		//} 
	}

	arrayCaraGanadora := utils.GetCaraGanadora()
	if random < 0.5 {
		fmt.Printf("%2s \r", arrayCaraGanadora[0])
	} else if random > 0.5 {
		fmt.Printf("%2s \r", arrayCaraGanadora[1])
	}
}

var clearScreen = func() {
	fmt.Print("os ", runtime.GOOS)
    cmd := exec.Command("clear") // clears the scrollback buffer as well as the screen.
    cmd.Stdout = os.Stdout
    cmd.Run()
}