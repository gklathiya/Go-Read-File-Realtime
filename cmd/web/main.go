package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
)

type PlugData struct {
	StationID         string `json:"stationID"`
	StationName       string `json:"stationName"`
	DeviceID          string `json:"deviceID"`
	DeviceName        string `json:"deviceName"`
	IsAvailable       string `json:"isAvailable"`
	CurrentTime       string `json:"currentTime"`
	CurrentDate       string `json:"currentDate"`
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	Voltages          string `json:"voltages"`
	TotalTransfer     string `json:"totalTransfer"`
	TransferSpeed     string `json:"transferSpeed"`
	TotalTime         string `json:"totalTime"`
	FrequencyInterval string `json:"frequencyInterval"`
}

func main() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event.Op)
				file, _ := ioutil.ReadFile("./data.json")
				var data PlugData
				err := json.Unmarshal(file, &data)
				if err != nil {
					fmt.Println("ERROR", err)
				} else {
					fmt.Println(data.StationID)
				}

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR in Watching :", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("./data.json"); err != nil {
		fmt.Println("ERROR in Adding : ", err)
	}

	<-done
}
