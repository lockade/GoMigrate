package escritor

import (
	"log"
	"os"
)

// Escreve em um arquivo :)
func Escrever(path string, filename string, texto string) bool {

	if criarDiretorio(path) != true {
		return false
	}

	file_name_full := path + "/" + filename

	file, erro := os.Create(file_name_full)

	if erro != nil {
		log.Println("Erro ao criar arquivo.", file_name_full)
		return false
	}

	defer file.Close()

	_, erro2 := file.WriteString(texto)

	if erro2 != nil {
		log.Println("Erro ao escrever em arquivo.")
		return false
	}

	return true
}

func criarDiretorio(path string) bool {
	if erro := os.MkdirAll(path, os.ModePerm); erro != nil {
		log.Println("Falha ao criar diretório. ", path)
		return false
	}

	return true
}
