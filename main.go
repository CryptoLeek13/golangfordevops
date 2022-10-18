package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step")
	} else {
		fmt.Println("Plan a better disguise next time")
	}
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault <= 70 {
		isHeistOn = false
		fmt.Println("The vault can;t be opened")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("The heist failed")
		case 1:
			isHeistOn = false
			fmt.Println("When did they start raising dogs in vaults??")
		case 2:
			isHeistOn = false
			fmt.Println("Looks like this fingerprint scanner won’t accept any fingerprint…")
		case 3:
			isHeistOn = false
			fmt.Println("Did I even pack the burlap bags?")
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			amtStolen := 1000 + rand.Intn(1000000)
			fmt.Printf("they stole %v$", amtStolen)
		}
	}

	fmt.Println("isHeistOn is currently:", isHeistOn)

}
