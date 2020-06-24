package routing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoPlayAndFun/Distributed-File-System/internal/lockservice"
)

func release(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reached release")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = lockservice.Release(string(body))

	if err != nil {
		fmt.Println(err)
	}
}

func checkRelease(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reached check release")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = lockservice.CheckRelease(string(body))

	if err != nil {
		fmt.Println(err)
	}
}
