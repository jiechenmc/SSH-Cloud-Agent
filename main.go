package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"ssh_cloud_agent/core"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	file, err := os.Open("./servers")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		user := strings.Split(line, "@")[0]
		ip := strings.Split(line, "@")[1]

		ssh, err := core.NewSshClient(
			user,
			ip,
			22,
			os.Getenv("HOME")+"/.ssh/id_rsa",
			"")

		check(err)

		files, err := os.ReadDir("./files")
		check(err)
		var wg sync.WaitGroup

		for _, file := range files {
			wg.Add(1)
			go func(file fs.DirEntry) {
				srcFilePath := "./files/" + file.Name()
				remoteFilePath := file.Name()
				_, err = ssh.CopyFile(srcFilePath, remoteFilePath)
				if err != nil {
					fmt.Printf("%s@%s %s\n", user, ip, err)
				} else {
					fmt.Printf("%s@%s\t%s -> %s GOOD\n", user, ip, srcFilePath, remoteFilePath)
				}

				defer wg.Done()
			}(file)
		}
		wg.Wait()
		wg.Add(1)
		go func() {
			ssh.RunCommand("sudo ./up.sh")
			fmt.Printf("UP RAN\n")
			defer wg.Done()
		}()
		wg.Wait()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	check(err)
}
