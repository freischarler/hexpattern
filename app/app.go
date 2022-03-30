package app

import (
	"log"
	"net/http"

	"github.com/freischarler/hexpattern/domain"
	"github.com/freischarler/hexpattern/service"
	"github.com/gorilla/mux"
)

//USER- HandlerAdapter -> IPortService -> (Domain LOGIC) -> IPortRepository -> DBAdapter -> DB

func Start() {

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)
	router.HandleFunc("/customers", ch.getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
