package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/parthverma/CRUDapplication/router"
)

func main() {
	router := router.Router()
	port := 9000
	fmt.Println("Server started listening at port " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
