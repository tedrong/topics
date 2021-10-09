package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type DashboardController struct{}

func (ctrl DashboardController) SystemInfo(c *gin.Context) {
	type percentage struct {
		CPU    string `json:"CPU"`
		Memory string `json:"Memory"`
		Disk   string `json:"Disk"`
	}
	type response struct {
		BootTime   int64      `json:"bootTime"`
		Percentage percentage `json:"percentage"`
	}
	res := response{}

	now := time.Now().UTC()
	cpu, _ := cpu.Percent(2*time.Second, false)
	mem, _ := mem.VirtualMemory()
	disk, _ := disk.Usage("/")
	host, _ := host.BootTime()
	duration := (now.Unix() - int64(host))

	res.Percentage.CPU = fmt.Sprintf("%.2f", cpu[0])
	res.Percentage.Memory = fmt.Sprintf("%.2f", mem.UsedPercent)
	res.Percentage.Disk = fmt.Sprintf("%.2f", disk.UsedPercent)
	res.BootTime = duration
	c.JSON(http.StatusOK, res)
}
