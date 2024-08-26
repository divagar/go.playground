package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runCommandCustomOutput(cmmd string) {
	var stdOut, stdError bytes.Buffer
	cmd := exec.Command(cmmd)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdError

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Commmand execution failed - %s\n", err)
		os.Exit(1)
	}

	outStr, errStr := string(stdOut.Bytes()), string(stdError.Bytes())
	fmt.Printf("Out: \n%s\n Err: \n%s\n", outStr, errStr)
}

func runCommandOutputToOsStd(cmmd string) {
	cmd := exec.Command(cmmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Commmand execution failed - %s\n", err)
		os.Exit(1)
	}
}

func main() {
	runCommandOutputToOsStd("systeminfo")
	runCommandCustomOutput("systeminfo")
}
