package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"teca", "viralata", 8}
	fmt.Println(c)

	cachorroEmJSON, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome":  "teca",
		"idade": "8",
		"raca":  "viralata",
	}

	dogEmJSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogEmJSON)
	fmt.Println(bytes.NewBuffer(dogEmJSON))

}
