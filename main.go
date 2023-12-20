package main

import (
	"fmt"
	"log"
	"os"
	"ssh_cloud_agent/core"
)

func main() {

	fmt.Println(os.Getenv("HOME"))
	ssh, err := core.NewSshClient(
		"azureuser",
		"52.151.255.24",
		22,
		os.Getenv("HOME")+"/.ssh/id_rsa",
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
