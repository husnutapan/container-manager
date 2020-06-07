package main

import (
	"container-manager/configs"
	"container-manager/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	clientConfig := configs.ClientConfig{}
	clientConfig.CreateClient()

	namespaceBase := controller.NewBaseHandler(&clientConfig)

	router := &mux.Router{}
	router.HandleFunc("/namespaces", namespaceBase.GetNamespaces).Methods("GET")
	router.HandleFunc("/pods/{namespace}", namespaceBase.GetPods).Methods("GET")

	http.ListenAndServe(":9090", router)
}
