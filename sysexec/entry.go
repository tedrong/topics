package sysexec

import (
	"log"
	"os"
	"os/exec"
)

func FindWebDriverPID() []byte {
	cmd := exec.Command("lsof", "-t", "-i:"+os.Getenv("CRAWLER_PORT"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("No WebDriver process running, ready to start: %s", err)
		return nil
	}
	return out[:len(out)-1]
}

func KillWebDriver(pid []byte) {
	cmd := exec.Command("kill", string(pid))
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
