package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

// Recupera os dados das listas salvos no snapshot.json e em seguida os dados em logs.txt
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

	fmt.Println("Content loaded!")
	fmt.Println("Loading Operations...")

	loadLogOperations(rl)

	//TODO: Limpar o arquivo de logs após as operações já terem sido executadas

}

// Executa as operações salvas no arquivos de logs para recriar o estado das listas
func loadLogOperations(rl *RemoteList) {
	file, err := os.Open(logs)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")

		listId, _ := strconv.Atoi(args[1])

		// args[0] -> Operação
		// args[1] -> Id da lista
		// args[2] -> Valor numérico (Ausente quando a operação é 'Remove') 

		if args[0] == "Append" {
			value, _ := strconv.Atoi(args[2])
			rl.listsMap[listId] = append(rl.listsMap[listId], value)

		} else if args[0] == "Remove" {
			list, _ := rl.getList(listId)
			lastIndex := len(list) - 1
			rl.listsMap[listId] = list[:lastIndex]
		} else {
			fmt.Println("Stopped!!")
			break
		}

	}
	fmt.Println("It worked!!")
	fmt.Println(rl.listsMap[1])

}


/* Registra operação executada em um arquivo de log
	- operation: "Append" ou "Remove"
	- listId: Id da lista que foi modificada
	- value: Valor da operação. Obs.: Deve ser do tipo "int" ou "nil"
*/
func RegisterLog(operation string, listId int, value any) {

	line := fmt.Sprintf("%s %d %v \n", operation, listId, value)

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
