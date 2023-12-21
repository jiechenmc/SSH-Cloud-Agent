package main

import (
	"fmt"
	"log"
	"os"
	"ssh_cloud_agent/core"
)

func main() {

	ssh, err := core.NewSshClient(
		"azureuser",
		"52.151.255.24",
		22,
		os.Getenv("HOME")+"/.ssh/id_rsa",
		"")

	if err != nil {
		log.Printf("SSH init error %v", err)
	} else {

		files, err := os.ReadDir("./files")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			go ssh.CopyFile("./files/"+file.Name(), file.Name())
		}

		output, err := ssh.RunCommand("./echo.sh")
		fmt.Println(output)

		if err != nil {
			log.Printf("SSH run command error %v", err)
		}
	}

}
