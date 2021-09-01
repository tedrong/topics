package sysexec

import (
	"log"
	"os/exec"
)

func FindWebDriverPID(port string) []byte {
	cmd := exec.Command("lsof", "-t", "-i:"+port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("No WebDriver process running, ready to start: %s", string(output))
		return nil
	}
	return output[:len(output)-1]
}

func KillWebDriver(pid []byte) {
	cmd := exec.Command("kill", string(pid))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", string(output))
	}
}
