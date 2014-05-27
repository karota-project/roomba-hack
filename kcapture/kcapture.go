package kcapture

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*  run a command */
func run(command string, args []string) (isSuccessed bool, err error) {
	exec.LookPath(command)

	cmd := exec.Command(command)
	cmd.Env = os.Environ()
	cmd.Args = args
	cmd.Dir = "."
	stdout, err := cmd.StdoutPipe()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return false, err
	}

	var n int
	buf := make([]byte, 2048)

	for {
		if n, err = stdout.Read(buf); err != nil {
			break
		}
		log.Print(string(buf[0:n]))
	}

	return true, nil
}

/* start */
func Start(cmd string, args []string) (isSuccessed bool, err error) {
	_args := []string{}

	if cmd == "ffmpeg" {
		if args == nil {
			/* ex) ffmpeg -s 320x240  -f video4linux2 -i /dev/video0  http://127.0.0.1:8090/feed1.ffm */
			_args = []string{cmd, "-s", "320x240", "-f", "video4linux2", "-i", "/dev/video0", "-i", "http://127.0.0.1:8090/feed1.ffm"}
		} else {
			tmp := strings.Join(args, " ")
			_args := strings.Split(cmd+tmp, " ")
			fmt.Println(_args)
		}
	} else if cmd == "ffserver" {
		if args == nil {
			/* ex) $ ffserver -d -f /etc/ffserver.conf & */
			_args = []string{cmd, "-d", "-f", "/etc/ffserver.conf &"}
		} else {
			tmp := strings.Join(args, " ")
			_args := strings.Split(cmd+tmp, " ")
			fmt.Println(_args)
		}
	} else {
		fmt.Println("Not support command.")
		return false, err
	}

	isSuccessed, err = run(cmd, _args)
	if err != nil {
		log.Println(err)
		return isSuccessed, err
	}

	if cmd == "ffmpeg" {
		fmt.Println("Starting ffmpeg")
	} else if cmd == "ffserver" {
		fmt.Println("Starting ffserver")

	}

	return isSuccessed, nil
}

/* stop */
func Stop(proc string) (isSuccessed bool, err error) {
	// use pkill
	cmd := "pkill"
	_args := []string{cmd, "-f", proc}

	isSuccessed, err = run(cmd, _args)
	if err != nil {
		log.Println(err)
		return isSuccessed, err
	}

	fmt.Println("Stopping videoCapture")

	return isSuccessed, nil
}
