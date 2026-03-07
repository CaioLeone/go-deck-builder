package deck

import "fmt"

func CreateDeck() {
	var deckName string

	fmt.Println("Digite o nome do Deck:")
	fmt.Scan(&deckName)

	deck = Deck{
		Name: deckName,
	}

	fmt.Println("Deck criado")
}

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
