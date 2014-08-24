package main

import (
	"./kcapture"
	"fmt"
	"time"
)

func main() {
	err := kcapture.Start("ffmpeg", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(5 * time.Second)

	err = kcapture.Start("ffserver", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Media Streaming...")

	time.Sleep(60 * time.Second)

	err = kcapture.Stop("ffmpeg")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = kcapture.Stop("ffserver")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}
