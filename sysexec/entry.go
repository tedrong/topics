package sysexec

import (
	"fmt"
	"log"
	"os/exec"
	"topics/config"
)

func FindWebDriverPID() []byte {
	cfg := config.Get()
	cmd := exec.Command("lsof", "-t", "-i:"+fmt.Sprintf("%d", cfg.Crawler.Port))
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
