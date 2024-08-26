package main

import (
	"fmt"
	"io"
	"log"
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"
	"encoding/json"
)

func runCommand(cmmd string, args []string) {
	// execute and get a pipe
	cmd := exec.Command(cmmd, args[0:]...)

	// stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return
	}
	// stderr
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		return
	}
	
	if err := cmd.Start(); err != nil {
		log.Println(err)
		return
	}
	
	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		postTaskUpdate("inprogress", s.Bytes())
	}
	
	if err := cmd.Wait(); err != nil {
		log.Println(err)
		return
	}
	postTaskUpdate("completed", s.Bytes())
}

func postTaskUpdate(status string, data []byte) {
	url := "http://localhost:3000/sysintegration/test/update"

	postBody, _ := json.Marshal(map[string]string{
		"status":  status,
		"data": string(data),
	 })
	// fmt.Println("\n-- postTaskUpdate --")
	// fmt.Println(string(postBody))
	// fmt.Println("\n")
	fmt.Printf("Sending update to %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	//fmt.Printf("Requst : %s\n", req)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	cmd := "systeminfo"
	args := []string{"/FO", "TABLE"}
	runCommand(cmd, args)
}