package main

import (
	"fmt"

	dataRef "github.com/caioleone/go-deck-builder/data"
	modelRef "github.com/caioleone/go-deck-builder/model"
	serviceRef "github.com/caioleone/go-deck-builder/service"
)

var decks modelRef.Deck

//var cards []modelRef.Card
//var nextID int64 = 1

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
		fmt.Println("7 - Adcionar Carta ao Deck")
		fmt.Println("8 - Salvar Deck em JSON")
		fmt.Println("==========")
		fmt.Println("0 - Sair")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			serviceRef.CreateCard()
		case 2:
			serviceRef.ListCard()
		case 3:
			serviceRef.EditCard()
		case 4:
			serviceRef.DeleteCard()
		case 5:
			serviceRef.CreateDeck(&decks)
		case 6:
			serviceRef.ListDeck(&decks)
		case 7:
			serviceRef.AddToDeck(&decks)
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
