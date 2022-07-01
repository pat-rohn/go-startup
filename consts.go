package startup

const (
	TimestampFormat string = "2006-01-02 15:04:05.000"
	logPkg          string = "startup"

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

type ServiceDetails struct {
	ExecutablePath string
	ServiceName    string
	UserName       string
	Description    string
}
