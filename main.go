package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Configurações do teste de stress
	numRequests := 100   // Número total de requisições a serem feitas
	concurrency := 10    // Número de requisições concorrentes
	targetURL := "http://exemplo.com"  // URL do alvo a ser testado

	// Criação do WaitGroup para sincronizar as goroutines
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// criação de um canal para receber os resultados
	results := make(chan int)

	// inicia as goroutines concorrentes
	for i := 0; i < concurrency; i++ {
		go makeRequest(targetURL, numRequests/concurrency, &wg, results)
	}

	// aguarda a conclusão de todas as goroutines
	go func() {
		wg.Wait()
		close(results)
	}()

	// processa os resultados recebidos
	totalRequests := 0
	for r := range results {
		totalRequests += r
	}

	// exibe os resultados finais
	fmt.Printf("Total de requisições feitas: %d\n", totalRequests)
}

func makeRequest(url string, numRequests int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	client := http.Client{
		Timeout: 10 * time.Second,  // Timeout para as requisições
	}

	for i := 0; i < numRequests; i++ {
		start := time.Now()

		resp, err := client.Get(url)
		if err != nil {
			log.Printf("Erro na requisição: %v\n", err)
			continue
		}

		// processa a resposta, se necessário

		resp.Body.Close()

		elapsed := time.Since(start)
		log.Printf("Requisição concluída em %s\n", elapsed)

		// envia o resultado para o canal
		results <- 1
	}
}
