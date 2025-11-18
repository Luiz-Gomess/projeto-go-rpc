package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var snapshotFile = "snapshot.json"
var logs = "logs.txt"

func Snapshot(rl *RemoteList) {

	jsonData, err := json.MarshalIndent(rl.listsMap, "", " ")

	if err != nil {
		panic(err)
	}
	err = os.WriteFile(snapshotFile, jsonData, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Snapshot created at ", time.Now())
}

func LoadData(rl *RemoteList) {
	jsonFile, err := os.Open(snapshotFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(byteValue, &rl.listsMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Content loaded")
	fmt.Println(rl.listsMap[1])

}

func RegisterLog(operation string, listId int, value any) {

	line := fmt.Sprintf("%s %d %d \n", operation, listId, value)

	file, err := os.OpenFile(logs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file for appending: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(line)

	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}
	fmt.Printf("[INFO] Operation %s saved on %s at %s \n", operation, logs, time.Now().Format(time.RFC3339))

}
