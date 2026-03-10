# Deck Builder

Projeto em Go para CRUD de jogos de cartas e construção de Decks utilziando linha de comando.

Objetivo do projeto:

* Struct
* Comunicação de packages
* Logica de negócio
* Tratamento de erros
* Salvar arquivo JSON
* Separação de responsabilidade

---

## Funcionalidades 

1. Cartas
   1. Criar 
   2. Editar
   3. Listas
   4. Deletar
2. Deck
   1. Criar
   2. Inserir Cartas
   3. Deletar Cartas
   4. Listar Deck
3. JSON
   1. Salvar Deck em Arquivo JSON

## Estrutura do projeto

* data -> Storage.go
* model -> cardModel.go
      -> deckModel.go
* service -> cardService.go
        -> deckService.go
* main.go
* go.mod
* README.md

## Conceitos Praticados

1. Modelagem de domínio com Struct
2. Métodos associados a Struct
3. Encapsulamento
4. Tratamento de erros idiomático em Go
5. Salvar de arquivos JSON
6. Conversão de tipos
7. Separação entre regra de negócio e infraestrutura