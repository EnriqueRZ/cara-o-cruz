package main

import (
	"fmt"
	"cara-cruz/utils"
);

func main() {
	random := float64(utils.RandIntn(1<<53)) / (1 << 53)
	fmt.Print("RANDOM = ", random);
}