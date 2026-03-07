package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caioleone/go-deck-builder/deck"
)

func SaveJson() {
	data, err := json.MarshalIndent(deck, "", "  ")

	if err != nil {
		fmt.Println("Erro ao salvar")
		return
	}

	os.WriteFile("/Data/deck.json", data, 0644)
	fmt.Println("Deck Salvo em deck.json")
}
