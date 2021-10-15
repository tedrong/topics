package crontab

import (
	"github.com/topics/database"
	"github.com/topics/logging"
	"github.com/topics/models"
)

type Internal struct {
	BasicCron
}

var dashboardModel = new(models.DashboardModel)

func (m *Internal) Period() string {
	return "@daily"
}

func (m *Internal) Do() {
	info, err := dashboardModel.SystemInfo()
	if err != nil {
		m.LogJob(logging.Get().Panic(), CJ_Internal).Err(err)
	}
	dashboardModel.SystemInfoStore(&database.Consumption{
		CPU:    info.CPU,
		Memory: info.Memory,
		Disk:   info.Disk,
	})
}
