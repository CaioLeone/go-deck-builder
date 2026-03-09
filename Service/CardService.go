package Service

import (
	"fmt"

	modelRef "github.com/caioleone/go-deck-builder/model"
)

var cards []modelRef.Card
var nextID int64 = 1

func CreateCard() {
	var name string
	var cardType string
	var attack int64
	var defence int64

	fmt.Println("Nome da carta:")
	fmt.Scan(&name)
	fmt.Println("Tipo da carta:")
	fmt.Scan(&cardType)
	fmt.Println("Ataque da carta:")
	fmt.Scan(&attack)
	fmt.Println("Defesa da carta:")
	fmt.Scan(&defence)

	card := modelRef.Card{
		ID:      nextID,
		Name:    name,
		Type:    cardType,
		Attack:  attack,
		Defence: defence,
	}

	nextID++
	cards = append(cards, card)
	fmt.Println("Carta Forjada Com Sucesso")
}

func ListCard() {
	for _, card := range cards {
		fmt.Printf("ID: %d | Nome: %s | Tipo: %s | Ataque: %d | Defesa: %d \n",
			nextID, card.Name, card.Type, card.Attack, card.Defence)
	}
}

func DeleteCard() {
	var id int64
	fmt.Print("Qual Id da carta que sera deletada? ")
	fmt.Scan(&id)

	for i, card := range cards {
		if card.ID == id {
			cards = append(cards[:i], cards[i+1:]...)
			fmt.Println("Carta deletada!")
			return
		}
	}
	fmt.Println("Carta nao encontrada")
}

func EditCard() {
	var id int64

	fmt.Println("Listar todas as cartas:")
	ListCard()

	fmt.Println("Qual Id da carta que deseja editar: ")
	fmt.Scan(&id)

	for i, card := range cards {
		if card.ID == id {
			fmt.Println("Novo Nome:")
			fmt.Scan(&cards[i].Name)

			fmt.Println("Novo Tipo:")
			fmt.Scan(&cards[i].Type)

			fmt.Println("Novo Ataque:")
			fmt.Scan(&cards[i].Attack)

			fmt.Println("Nova Defesa:")
			fmt.Scan(&cards[i].Defence)

			fmt.Println("Carta editada com sucesso")
			return
		}
	}
	fmt.Println("Carta nao encontrada")
}
