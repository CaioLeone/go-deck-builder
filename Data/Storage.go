package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caioleone/go-deck-builder/Deck"
)

func SaveJson(d Deck.Deck) {
	data, err := json.MarshalIndent(d, "", "  ")

	if err != nil {
		fmt.Println("Erro ao salvar")
		return
	}

	err = os.WriteFile("/Data/deck.json", data, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo")
	}
	fmt.Println("Deck Salvo em deck.json")
}
