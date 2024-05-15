package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// grep -r Port /etc/
func sol1() {
	err := os.Chdir("/etc/")
	if err != nil {
		log.Fatal("can't cd to /etc dir")
	}

	cmd := exec.Command("sudo", "grep", "-r", "Port", "/etc/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("can't execute grep command, error: ", err)
	}

	fmt.Println("command executed successfully, output:")
	fmt.Println(string(out))

}