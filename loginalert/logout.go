package main

import (
	"github.com/gen2brain/beeep"
	"github.com/sqweek/dialog"
)

func main() {
	dialog.Message("%s", "Hey Akshay s if you are logging out please checkout from rayzr pay").Title("Login Alert").Info()
	err := beeep.Beep(10, 300)
	if err != nil {
		panic(err)
	}

}
