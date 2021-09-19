package startup

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InstallService(details ServiceDetails) error {

	f, err := os.OpenFile("/etc/systemd/system/"+details.ServiceName+".service", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.WithField("package", logPkg).Errorf("open failed: %v", err)
		return err
	}
	cmd := `ExecStart=/home/` +
		details.UserName + `/bin/SimpleObserver $` +
		details.ServiceName + `1 $` + details.ServiceName + `2 $` +
		details.ServiceName + `3 $` + details.ServiceName + `4`
	serviceText := `
[Unit]
Description=` + details.Description + `
Requires=network.target

[Service]
Type=simple
Restart=always
EnvironmentFile=/home/` + details.UserName + `/bin/services.conf
WorkingDirectory=/home/` + details.UserName + `/bin/` + details.ServiceName + `
` + cmd + `
User=` + details.UserName + `

[Install]
WantedBy=multi-user.target
`
	if _, err = f.WriteString(serviceText); err != nil {
		log.WithField("package", logPkg).Errorf("write failed: %v", err)
	}
	f.Close()
	descr := strings.Replace(serviceDescription, "<myservicename>", details.ServiceName, -1)
	descr = strings.Replace(descr, "<myuserame>", details.UserName, -1)
	fmt.Println(descr)
	serviceFile, err := os.OpenFile("/home/"+details.UserName+"/bin/services.conf",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.WithField("package", logPkg).Errorf("write servicesConf failed: %v", err)
		return err
	}
	servicesConf := `
	` + details.ServiceName + `1=-f
	` + details.ServiceName + `2=-e
	` + details.ServiceName + `3=` + details.ServiceName + `
	` + details.ServiceName + `4=
`
	defer serviceFile.Close()
	if _, err = serviceFile.WriteString(servicesConf); err != nil {
		log.WithField("package", logPkg).Errorf("write servicesConf failed: %v", err)
		return err
	}

	logPath := LogPath
	err = os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		log.WithField("package", logPkg).Errorf("failed to create directory: %v %v", err, logPath)
		return err
	}
	myUser, err := user.Lookup(details.UserName)
	if err != nil {
		log.WithField("package", logPkg).Errorf("User lookup: %v", err)
	}
	myUserID, _ := strconv.Atoi(myUser.Uid)
	myGID, _ := strconv.Atoi(myUser.Gid)
	err = os.Chown(logPath, myUserID, myGID)
	if err != nil {
		log.WithField("package", logPkg).Errorf("Chown: %v", err)
	}

	return nil
}
