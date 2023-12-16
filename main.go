package main

import (
	"fmt"
	"log"
	"ssh_cloud_agent/core"
)

func main() {
	ssh, err := core.NewSshClient(
		"azureuser",
		"52.151.255.24",
		22,
		"/home/chenj7728/.ssh/id_rsa",
		"")

	if err != nil {
		log.Printf("SSH init error %v", err)
	} else {
		output, err := ssh.RunCommand("sudo docker ps -a")
		fmt.Println(output)
		if err != nil {
			log.Printf("SSH run command error %v", err)
		}
	}

}
