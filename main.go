package main

import (
	"github.com/deadman360/WbScrapper/entrada"
	"github.com/deadman360/WbScrapper/workers"
)

func main() {
	lista := entrada.PedeLista()
	workers.RecebeLista(lista)
	workers.Wg.Wait()

}
