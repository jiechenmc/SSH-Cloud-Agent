package main

import (
	"io/fs"
	"log"
	"os"
	"ssh_cloud_agent/core"
	"sync"
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

		var wg sync.WaitGroup

		for _, file := range files {
			wg.Add(1)
			go func(file fs.DirEntry) {
				ssh.CopyFile("./files/"+file.Name(), file.Name())
				defer wg.Done()
			}(file)
		}
		wg.Wait()
		if err != nil {
			log.Printf("SSH run command error %v", err)
		}
	}

}
