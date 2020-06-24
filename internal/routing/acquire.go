package routing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoPlayAndFun/Distributed-File-System/internal/lockservice"
)

func acquire(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reached acquire")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = lockservice.Acquire(string(body))

	if err != nil {
		fmt.Println(err)
	}
}

func checkAcquire(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reached check acquire")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/groceries.go/GetItemsFromCart")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = lockservice.CheckAcquire(string(body))

	if err != nil {
		fmt.Println(err)
	}
}
