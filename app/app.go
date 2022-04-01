package app

import (
	"log"
	"net/http"

	"github.com/Jestinepa/doctor/domain"
	"github.com/Jestinepa/doctor/service"
	"github.com/gorilla/mux"
)

func Start() {

	myRouter := mux.NewRouter()

	//wiring
	dh := DoctorHandlers{service.NewDoctorService(domain.NewDoctorRespositoryStub())}

	//define routes
	myRouter.HandleFunc("/getavailability", dh.getavailability)

	//starting server

	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}
