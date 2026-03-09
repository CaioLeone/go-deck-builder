package Service

import (
	"fmt"

	modelRef "github.com/caioleone/go-deck-builder/model"
)

func CreateDeck(d *modelRef.Deck) {
	var deckName string

	fmt.Println("Digite o nome do Deck:")
	fmt.Scan(&deckName)

	d.Name = deckName

	fmt.Println("Deck criado")
}

func AddToDeck(d *modelRef.Deck) {
	var id int64

	fmt.Println("Vamos Listar todas as cartas disponiveis")
	ListCard()

	fmt.Print("Id da carta para adicionar: ")
	fmt.Scan(&id)

	for _, card := range d.Cards {
		if card.ID == id {
			d.Cards = append(d.Cards, card)
			fmt.Println("Carta adicionada ao deck")
			return
		}
	}

	fmt.Println("Carta nao encontrada")
}

func ListDeck(d *modelRef.Deck) {
	fmt.Println("Deck: ", d.Name)
	for _, card := range d.Cards {
		fmt.Printf("ID: %d | Nome: %s | Tipo: %s | Ataque: %d | Defesa: %d \n",
			nextID, card.Name, card.Type, card.Attack, card.Defence)
	}
}
