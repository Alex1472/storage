package main

import (
	"fmt"
	"log"
)
import "github.com/Alex1472/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	f, err := st.Upload("test.txt", []byte("Hello!"))
	if err != nil {
		log.Fatal(err)
	}

	f2, err := st.GetByID(f.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f2)
}
