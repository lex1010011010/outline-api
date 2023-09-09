package main

import (
	"encoding/json"
	"fmt"
	"github.com/lex1010011010/outline-api/outline"
	"log"
)

//https://myserver/SecretPath/server/hostname-for-access-keys

func main() {
	apiURl := "https://45.10.247.116:58498/uLrpv-Jsi9OAu3Mg79iTmQ"
	manager := outline.NewManager(apiURl, "apiCrt")

	err := manager.UpdateMetricsStatus(true)
	if err != nil {
		fmt.Println(err)
	}

	severInfo, err := manager.GetServerInfo()
	if err != nil {
		return
	}

	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(severInfo, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
