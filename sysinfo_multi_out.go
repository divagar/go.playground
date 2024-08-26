package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func runCommand(cmmd string) {
	var stdOutBuffer, stdErrBuffer bytes.Buffer
	cmd := exec.Command(cmmd)

	cmd.Stdout = io.MultiWriter(os.Stdout, &stdOutBuffer)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stdErrBuffer)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to execute the command - %s", err)
		os.Exit(1)
	}

	fmt.Printf("Output : \n%s\n", string(stdOutBuffer.Bytes()))
	fmt.Printf("Error : \n%s\n", string(stdErrBuffer.Bytes()))

}

func main() {
	fmt.Println("System info with two output stream")
	runCommand("tasklist")
}
