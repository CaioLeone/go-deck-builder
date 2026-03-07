package main

import (
	"fmt"

	"github.com/caioleone/go-deck-builder/card"
	"github.com/caioleone/go-deck-builder/data"
	"github.com/caioleone/go-deck-builder/deck"
)

var cards []card.Card
var decks deck.Deck
var nextID int64 = 1

func main() {

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
			card.CreateCard()
		case 2:
			card.ListCard()
		case 3:
			card.EditCard()
		case 4:
			card.DeleteCard()
		case 5:
			deck.CreateDeck()
		case 6:
			deck.ListDeck()
		case 7:
			deck.AddToDeck()
		case 8:
			data.SaveJson()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opcao Invalida")
		}
	}
}
