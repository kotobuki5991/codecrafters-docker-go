package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"syscall"

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

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}

	setupMyDockerApp()

	err := cmd.Run()
	// exitステータスの確認
	if err != nil {
		fmt.Printf("Err: %v", err)
		errCode := getExitStatus(err)
		os.Exit(errCode)
	}

	// 標準出力の有無と内容確認
	if stdout.String() == "" {
		fmt.Print("No such file or directory")
		os.Exit(2)
	}else {
		fmt.Println(strings.Trim(stdout.String(), "\n"))
		os.Exit(1)
	}
}

func setupMyDockerApp()  {
	syscall.Chroot("/chroot")
	// os.Chdir("/chroot")
	// cpCmd := exec.Command("cp", "/usr/local/bin/docker-explorer", "/chroot");
	// err := cpCmd.Run()
	// if err != nil {
	// 	fmt.Printf("Err: %v", err)
	// 	errCode := getExitStatus(err)
	// 	os.Exit(errCode)
	// }
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
