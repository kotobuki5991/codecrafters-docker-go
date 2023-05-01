package main

import (
	"fmt"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage!
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// Uncomment this block to pass the first stage!
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Err: %v", err)
		errCode := getExitStatus(err)
		os.Exit(errCode)
	}
}

func getExitStatus(err error) int{
	errMsg := err.Error()
	msgAry := strings.Split(errMsg, " ")
	errCode, err := strconv.Atoi(msgAry[len(msgAry)-1])
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	return errCode
}
