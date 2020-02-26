package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anyric/bts/src/api/router"
	"github.com/anyric/bts/src/config"
)

//Run starts the app on port 8080
func Run() {
	fmt.Printf("\n\tListening [::]:%d\n\n", config.PORT)
	listen(config.PORT)

}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
