package app

import (
	"github.com/KanDevArg/yourdreamhome/go-backend/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/listings/postal-code", controllers.GetListings)

	http.HandleFunc("/listings", controllers.GetAllListings)

	http.HandleFunc("/ping",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Ping here!!"))
		})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
