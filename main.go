package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lcycug/go-config-service/controller"
	"github.com/lcycug/go-config-service/domain"
)

func main() {
	config := domain.Config{}

	data, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}

	cfg, err := config.Get("service.registry.device")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	ctl := controller.Controller{
		Config: &config,
	}

	r := mux.NewRouter()
	r.HandleFunc("/read/{serviceName}", ctl.ReadConfig).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
