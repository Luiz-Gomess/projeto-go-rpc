package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var snapshotFile = "snapshot.json"

func Snapshot(rl *RemoteList){
	
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

func LoadData(rl *RemoteList){
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