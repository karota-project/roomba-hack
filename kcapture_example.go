package main

import (
	"./kcapture"
	"fmt"
	"time"
)

func main() {
	m, err := kcapture.Start("ffmpeg", nil)

	time.Sleep(5 * time.Second)

	s, err2 := kcapture.Start("ffserver", nil)

	if err != nil || err2 != nil {
		fmt.Println(m, err, s, err2)
		return
	}

	fmt.Println("media streaming...")

	time.Sleep(60 * time.Second)

	m, err = kcapture.Stop("ffmpeg")
	s, err2 = kcapture.Stop("ffserver")

	if err != nil || err2 != nil {
		fmt.Println(m, err, s, err2)
		return
	}

	fmt.Println("finished example")
}
