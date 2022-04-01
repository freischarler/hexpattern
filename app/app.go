package app

import (
	"log"
	"net/http"
	"os"

	"github.com/freischarler/hexpattern/domain"
	"github.com/freischarler/hexpattern/service"
	"github.com/gorilla/mux"
)

//USER- HandlerAdapter -> IPortService -> (Domain LOGIC) -> IPortRepository -> DBAdapter -> DB

const defaultAddr = "localhost"
const defaultPort = "9000"

func Start() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}

	addr := os.Getenv("IP")
	if addr == "" {
		addr = defaultAddr
	}
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)
	router.HandleFunc("/customers", ch.getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe(addr+":"+serverPort, router))
}
