package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Card struct {
	ID      int64
	Name    string
	Type    string
	Attack  int64
	Defence int64
}

type Deck struct {
	Name  string
	Cards []Card //SLICE
}

var cards []Card
var deck Deck
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
		fmt.Println("8 - Deletar Deck")
		fmt.Println("9 - Salvar Deck em JSON")
		fmt.Println("==========")
		fmt.Println("0 - Sair")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			CreateCard()
		case 2:
			ListCard()
		case 3:
			EditCard()
		case 4:
			DeleteCard()
		case 5:
			CreateDeck()
		case 6:
			ListDeck()
		case 7:
			AddToDeck()
		case 8:
			DeleteDeck()
		case 9:
			SaveJson()
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opcao Invalida")
		}
	}
}

//SAVE JSON

func SaveJson() {
	data, err := json.MarshalIndent(deck, "", "  ")

	if err != nil {
		fmt.Println("Erro ao salvar")
		return
	}

	os.WriteFile("deck.json", data, 0644)
	fmt.Println("Deck Salvo em deck.json")
}

//CARD FUNC

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

	card := Card{
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

func CreateDeck() {
	var deckName string

	fmt.Println("Digite o nome do Deck:")
	fmt.Scan(&deckName)

	deck = Deck{
		Name: deckName,
	}

	fmt.Println("Deck criado")
}

//DECK FUNC

func AddToDeck() {
	var id int64

	fmt.Println("Vamos Listar todas as cartas disponiveis")
	ListCard()

	fmt.Print("Id da carta para adicionar: ")
	fmt.Scan(&id)

	for _, card := range cards {
		if card.ID == id {
			deck.Cards = append(deck.Cards, card)
			fmt.Println("Carta adicionada ao deck")
			return
		}
	}

	fmt.Println("Carta nao encontrada")
}

func ListDeck() {
	fmt.Println("Deck: ", deck.Name)
	for _, card := range cards {
		fmt.Printf("ID: %d | Nome: %s | Tipo: %s | Ataque: %d | Defesa: %d \n",
			nextID, card.Name, card.Type, card.Attack, card.Defence)
	}
}
func DeleteDeck() {}
