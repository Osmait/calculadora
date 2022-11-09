package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	count := 0

	for {
		if count == 3 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		count++

	}

	tiempoPaso := time.Since(inicio)

	fmt.Printf("tiempo de ejecucion %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {

		canal <- servidor + "No se encuentra disponible"
	}

	canal <- servidor + "esta funcionando"

}
