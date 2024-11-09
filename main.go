package main
/*
Add imports
*/
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//create a fileServer (type :Handler) and init
	fileServer := http.FileServer(http.Dir("./static"))
	//set the static folder as root
	http.Handle("/", fileServer)
	//Print starting
	fmt.Println("Starting server on port : 8080")
	//start server 
	err := http.ListenAndServe(":8080", nil)
	//check the error , if contains then log
	if err != nil {
		log.Fatal(err)
	}
}
