package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade string `json:"idade"`
}

func main() {

	//Conversão para struct
	cachorroEmJSON := `{"nome":"teca","raca":"viralata","idade":"8"}`

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	//Conversão para map
	cachorroEmJSON2 := `{"idade":"8","nome":"teca","raca":"viralata"}`

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(cachorroEmJSON2), &c2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)

}
