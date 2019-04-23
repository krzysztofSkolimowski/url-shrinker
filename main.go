package main

import (
	"context"
	"fmt"
	"github.com/krzysztofSkolimowski/url-shrinker/common"
	"github.com/krzysztofSkolimowski/url-shrinker/pkg/app"
	"github.com/krzysztofSkolimowski/url-shrinker/pkg/infrastructure"
	"github.com/krzysztofSkolimowski/url-shrinker/pkg/interfaces"
	"log"
	"net/http"
)

func main() {
	c := common.Config{}
	c.LoadConfig()

	log.Fatal(runServer(c))
}

type modules struct {
	//app
	sh app.Shrinker
	r  app.Retriever

	//infra
	s infrastructure.Storage
}

func buildModules(config common.Config) modules {
	//todo implement
	return modules{
	}
}

func runServer(config common.Config) error {
	fmt.Println("Server running")
	fmt.Println("actually not, todo!")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	modules := buildModules(config)
	http.Handle("/save", interfaces.SaveHandler(modules.sh))
	http.Handle("/key", interfaces.KeyHandler(modules.r))
	http.Handle("/Decode", interfaces.DecodeHandler(modules.r))
	http.Handle("/", interfaces.RedirectHandler(modules.r))

	return http.ListenAndServe(":"+config.Port(), nil)
}
