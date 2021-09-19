package startup

const (
	Version         string = "1.0.2"
	LogPath         string = "/var/log/mylogs/"
	TimestampFormat string = "2006-01-02 15:04:05.000"

	serviceDescription string = `
Service-file was stored in /etc/systemd/system/
Run following commands to activate service:
	sudo systemctl --system daemon-reload
	sudo systemctl enable <myservicename>
	sudo systemctl start <myservicename>
To show status
	sudo systemctl status <myservicename>
Adapt arguments in services.conf (/home/$USER/bin/)
	
	`
)

type SimpleObserver struct {
	ScanRate  int    `json:"ScanRate"`
	TableName string `json:"TableName"`
}

type ServiceDetails struct {
	ExecutableName string `json:"string"`
	ServiceName    string `json:"string"`
	UserName       string `json:"string"`
	Description    string `json:"string"`
}
