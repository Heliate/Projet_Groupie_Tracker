package main

import (
	"Api/config"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Printf("Serveur démarré sur http://%s\n", net.JoinHostPort(config.App.Server.URL, config.App.Server.Port))
	log.Fatal(http.ListenAndServe(net.JoinHostPort(config.App.Server.URL, config.App.Server.Port), nil))
}
