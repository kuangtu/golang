package main

import(
	"fmt"
	"net/http"
	_ "strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)

	fmt.Fprintf(w, "hello")

}

func main() {
	http.HandleFunc("/", sayHelloName)

	err := http.ListenAndServe(":9190", nil)

	if err != nil {

		log.Fatal("ListenAndServe", err)
	}


}
