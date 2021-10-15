package models

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/topics/database"
	"github.com/topics/forms"
	"gorm.io/gorm"
)

type DashboardModel struct{}

func (m DashboardModel) SystemInfoStore(info *database.Consumption) {
	db := database.GetPG(database.DBInternal)
	db.Create(info)
}

func (m DashboardModel) ClientTypeStore(header string) error {
	var client string
	re := regexp.MustCompile(`Android|iPhone`)
	if re.MatchString(header) {
		client = database.ClientIndex()[database.Mobile]
	} else {
		client = database.ClientIndex()[database.Desktop]
	}

	db := database.GetPG(database.DBInternal)
	err := db.Transaction(func(tx *gorm.DB) error {
		var feature database.Client
		if err := tx.Where("type = ?", client).First(&feature).Error; err != nil {
			return err
		}
		if err := tx.Model(&feature).Update("login_counter", feature.LoginCounter+1).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m DashboardModel) ClientTypePercentage() (map[string]int, error) {
	db := database.GetPG(database.DBInternal)
	var total int64
	result := db.Model(&database.Client{}).Select("sum(login_counter)").Scan(&total)
	if result.Error != nil {
		return nil, result.Error
	}
	percentage := make(map[string]int)
	for _, element := range database.ClientIndex() {
		var client database.Client
		result := db.Where("type = ?", element).Find(&client)
		if result.Error != nil {
			return nil, result.Error
		}
		percentage[element] = int(math.Round(float64(client.LoginCounter) * 100 / float64(total)))
	}
	return percentage, nil
}

func (m DashboardModel) SystemInfo() (*forms.SystemInfo, error) {
	var info forms.SystemInfo

	now := time.Now().UTC()
	cpu, err := cpu.Percent(2*time.Second, false)
	if err != nil {
		return nil, err
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	disk, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	host, err := host.BootTime()
	if err != nil {
		return nil, err
	}
	duration := (now.Unix() - int64(host))

	info.CPU = fmt.Sprintf("%.2f", cpu[0])
	info.Memory = fmt.Sprintf("%.2f", mem.UsedPercent)
	info.Disk = fmt.Sprintf("%.2f", disk.UsedPercent)
	info.BootTime = duration
	return &info, nil
}

func (m DashboardModel) SystemInfoHistory(timestamp time.Time) (map[string][]string, error) {
	db := database.GetPG(database.DBInternal)
	consumptions := []database.Consumption{}
	result := db.Where("created_at > ?", timestamp).Order("created_at").Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	content := map[string][]string{}
	for _, consumption := range consumptions {
		content["cpu"] = append(content["cpu"], consumption.CPU)
		content["memory"] = append(content["memory"], consumption.Memory)
		content["disk"] = append(content["disk"], consumption.Disk)
		content["label"] = append(content["label"], consumption.CreatedAt.String())
	}
	return content, nil
}

func (m DashboardModel) TailLog(lines int) (*[]*forms.SystemLog, error) {
	fileHandle, err := os.Open("./tmp/topics.log")
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	var logs []*forms.SystemLog
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()

	for i := 0; i < lines; i++ {
		var lineByte []byte
		var systemLog forms.SystemLog
		for {
			cursor -= 1
			fileHandle.Seek(cursor, io.SeekEnd)

			char := make([]byte, 1)
			fileHandle.Read(char)
			// stop if we find a line
			if cursor != -1 && (char[0] == 10 || char[0] == 13) {
				break
			}
			lineByte = append(char, lineByte...)
			// stop if we are at the begining
			if cursor == -filesize {
				break
			}
		}
		json.Unmarshal(lineByte, &systemLog)
		logs = append(logs, &systemLog)
	}

	return &logs, nil
}
