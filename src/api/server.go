package api

import (
	"net/http"
	"fmt"
	"log"

	"github.com/anyric/bts/src/config"
	"github.com/anyric/bts/src/api/router"
)

//Run starts the app on port 8080
func Run()  {
	config.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)

}

func listen(port int)  {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}