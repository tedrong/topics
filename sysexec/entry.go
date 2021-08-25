package sysexec

import (
	"log"
	"os/exec"
)

func FindWebDriverPID(port string) []byte {
	cmd := exec.Command("lsof", "-t", "-i:"+port)
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
