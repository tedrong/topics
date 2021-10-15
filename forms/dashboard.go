package forms

type SystemInfo struct {
	CPU      string `json:"cpu"`
	Memory   string `json:"memory"`
	Disk     string `json:"disk"`
	BootTime int64  `json:"bootTime"`
}

type SystemLog struct {
	Level   string `json:"level"`
	Time    int64  `json:"time"`
	Msg     string `json:"message"`
	Type    string `json:"type"`
	Server  string `json:"server"`
	CronJob string `json:"cronjob"`
}
