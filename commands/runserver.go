package main

import "github.com/gorilla/mux"

//http://www.omdbapi.com/?i=tt3896198&apikey=786b39f5
///https://api.themoviedb.org/3/movie/550?api_key=9da8d8ded94a727b9995bed399b6c315

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup

		wg.Add(1)
		go requestApi(w, 1, &wg)
		wg.Add(1)
		go requestApi(w, 2, &wg)
		wg.Add(1)
		go requestApi(w, 3, &wg)
		wg.Add(1)
		go requestApi(w, 4, &wg)
		wg.Add(1)
		go requestApi(w, 5, &wg)

		wg.Wait()
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":99", r)
}

func requestApi(w http.ResponseWriter, n int, wg *sync.WaitGroup) {
	//resp, err := http.Get("http://www.omdbapi.com/?i=tt3896198&apikey=786b39f5")
	defer wg.Done()

	resp, err := http.Get("http://www.omdbapi.com/?i=tt3896198&apikey=786b39f5")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	//log.Println(result)
	fmt.Fprint(w, n)
}
