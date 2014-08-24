package main

import (
	"fmt"
	"time"

	"github.com/karota-project/kcapture"
)

func main() {
	if err := kcapture.StartStreamer(nil); err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(5 * time.Second)

	if err := kcapture.StartServer(nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Media Streaming...")

	time.Sleep(60 * time.Second)

	if err := kcapture.StopStreamer(); err != nil {
		fmt.Println(err)
		return
	}

	if err := kcapture.StopServer(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}
