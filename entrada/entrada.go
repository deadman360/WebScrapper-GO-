package entrada

import "fmt"

func PedeLista() string {
	var listaURL string
	fmt.Println("Porfavor envie uma lista de URL separadas por vírgulas: ")
	fmt.Scan(&listaURL)
	fmt.Println("Começando o Scrapping...")
	return listaURL
}
