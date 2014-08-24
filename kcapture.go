package kcapture

import (
	"log"
	"os"
	"os/exec"
)

func StartServer(args []string) (err error) {
	if args == nil {
		/* ex) $ ffserver -d -f /etc/ffserver.conf & */
		args = []string{"-d", "-f", "/etc/ffserver.conf", "&"}
	}

	return run("ffserver", args)
}

func StartStreamer(args []string) (err error) {
	if args == nil {
		/* ex) ffmpeg -s 320x240  -f video4linux2 -i /dev/video0  http://127.0.0.1:8090/feed1.ffm */
		args = []string{"-s", "320x240", "-f", "video4linux2", "-i", "/dev/video0", "-i", "http://127.0.0.1:8090/feed1.ffm"}
	}

	return run("ffmpeg", args)
}

func StopServer() (err error) {
	return run("pkill", []string{"-f", "ffserver"})
}

func StopStreamer() (err error) {
	return run("pkill", []string{"-f", "ffmpeg"})
}

func run(command string, args []string) (err error) {
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Dir = "."

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	buf := make([]byte, 2048)

	for {
		n, err := stdout.Read(buf)

		if err != nil {
			break
		}

		log.Print(string(buf[0:n]))
	}

	return nil
}
