package main

import "fmt"

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

func main() {

}

func GameMenu() {
	for {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("0 - Sair")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opcao Invalida")
		}
	}
}
