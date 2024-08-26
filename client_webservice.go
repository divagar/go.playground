package main

import (
	"io"
	"fmt"
	"bufio"
	"bytes"
	"os/exec"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//Config struct
type ServiceConfig struct {
	CentralServerHost string
	CentralServerPort string
	CentralServerResource string
	HostCommand string
	HostCommandArgs []string
}
var serviceConfigObj ServiceConfig

func readConfig() {
	config, err := ioutil.ReadFile("./client_webservice.json")
    if err != nil {
      fmt.Println(err)
    }
	err = json.Unmarshal(config, &serviceConfigObj)
    if err != nil {
        fmt.Println(err)
    }
}

func runCommand() {
	// execute and get a pipe
	cmd := exec.Command(serviceConfigObj.HostCommand, serviceConfigObj.HostCommandArgs[0:]...)

	// stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	// stderr
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		postUpdate("inprogress", s.Bytes())
	}
	
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	postUpdate("completed", s.Bytes())
}

func postUpdate(status string, data []byte) {
	url := "http://" + serviceConfigObj.CentralServerHost + ":" + serviceConfigObj.CentralServerPort + "/" + serviceConfigObj.CentralServerResource
	postBody, _ := json.Marshal(map[string]string{
		"status":  status,
		"data": string(data),
	 })
	 
	fmt.Printf("Sending update to %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	resBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(resBody))
}

func main() {
	readConfig()
	runCommand()
}