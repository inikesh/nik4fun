package main

import (
	"fmt"
	"log"
	"net/http"
)

func hel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey nikks here")
}

func main() {
	http.HandleFunc("/", hel)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

/*import "fmt"
func main() {
	fmt.Println("hellsdjjasdnflkasd");
}*/
