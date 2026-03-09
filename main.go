package main

import (
	"fmt"

	dataRef "github.com/caioleone/go-deck-builder/Data"
	cardRef "github.com/caioleone/go-deck-builder/card"
	deckRef "github.com/caioleone/go-deck-builder/deck"
)

var cards []cardRef.Card
var decks deckRef.Deck
var nextID int64 = 1

func main() {
	GameMenu()
}

func GameMenu() {
	for {
		fmt.Println("")
		fmt.Println("===== MENU =====")
		fmt.Println("1 - Criar Carta")
		fmt.Println("2 - Listar Carta")
		fmt.Println("3 - Editar Carta")
		fmt.Println("4 - Deletar Carta")
		fmt.Println("===== DECK MENU =====")
		fmt.Println("5 - Criar Deck")
		fmt.Println("6 - Listar Deck")
		fmt.Println("7 - Editar Deck")
		fmt.Println("8 - Salvar Deck em JSON")
		fmt.Println("==========")
		fmt.Println("0 - Sair")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cardRef.CreateCard(&cards, &nextID)
		case 2:
			cardRef.ListCard(cards)
		case 3:
			cardRef.EditCard(cards)
		case 4:
			cardRef.DeleteCard(&cards)
		case 5:
			deckRef. CreateDeck(&decks)
		case 6:
			deckRef.ListDeck(decks)
		case 7:
			deckRef.AddToDeck(cards, &decks)
		case 8:
			dataRef.SaveJson(decks)
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opcao Invalida")
		}
	}
}
