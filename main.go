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

	//res, err := manager.CreateNewAccessKey()
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Printf("%#v", res)

	accKeys, err := manager.GetAccessKeys()
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(accKeys, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
