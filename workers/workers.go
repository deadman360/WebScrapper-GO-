package workers

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var Wg sync.WaitGroup

// LIMPA LISTA E FAZ A LOGICA PARA CRIAR WORKERS
func RecebeLista(lista string) {

	Slice := strings.Split(lista, ",")

	for _, v := range Slice {
		strings.TrimSpace(v)
	}
	for _, v := range Slice {
		Wg.Add(1)
		go CreateWorker(v)
	}
}

// CRIA SCRAPPER CONCORRENTE

func CreateWorker(url string) {
	c := colly.NewCollector(colly.AllowedDomains(url))
	c.OnHTML("div.WBVL_7", func(h *colly.HTMLElement) {

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.Method)
	})

	Wg.Done()

}
