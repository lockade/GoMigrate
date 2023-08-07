package main

import (
	"fmt"
	"lockade/migrate/escritor"
	"log"
	"os"
)

var path_default string = "./migrate"
var path_config string = path_default + "/config"
var file_config string = "/config.conf"
var path_migrate string = path_default

func main() {

	input_len := len(os.Args)

	if input_len <= 1 {
		log.Fatal("Digite algum comando.", 1)
	}

	first_param := os.Args[1]

	switch first_param {
	case "configurar":
		configurar()
	}
}

func configurar() {
	fmt.Print("Url: ")
	endereco := ler()

	fmt.Print("User: ")
	user := ler()

	fmt.Print("Pass: ")
	pass := ler()

	fmt.Print("Database: ")
	database := ler()

	erro := escritor.Escrever(path_config, file_config, "URL:"+endereco+"USER:"+user+"PASS:"+pass+"DATABASE:"+database)

	if !erro {
		log.Println("Não foi possível criar arquivo de configuração, tente novamente.")
		configurar()
	}

	//salvar input nas configurações

}

func ler() string {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Erro interno.")
	}

	return input
}
