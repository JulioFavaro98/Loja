package main

import (
	"net/http"

	"github.com/JulioFavaro98/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
