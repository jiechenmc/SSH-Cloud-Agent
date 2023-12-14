package main

import "golang.org/x/crypto/ssh"

func main() {
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("pass"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, _ := ssh.Dial("tcp", "remotehost:22", config)
	defer client.Close()
}
