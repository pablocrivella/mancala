package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/go-openapi/loads"

	"github.com/pablocrivella/mancala/api/openapi-v2/restapi"
	"github.com/pablocrivella/mancala/api/openapi-v2/restapi/operations"
	"github.com/pablocrivella/mancala/cmd/api/handlers"
	"github.com/pablocrivella/mancala/internal/games"
	"github.com/pablocrivella/mancala/internal/postgres"
)

const exitFail = 1

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run() error {
	db, err := postgres.OpenDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}
	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  mancala-server [OPTIONS]\n\n")

		title := "Mancala API"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := swaggerSpec.Spec().Info.Description
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		flag.CommandLine.SetOutput(os.Stderr)
		flag.PrintDefaults()
	}
	// parse the CLI flags
	flag.Parse()
	api := operations.NewMancalaAPI(swaggerSpec)

	gameStore := postgres.NewGameStore(db)
	gameService := games.Service{GameStore: gameStore}
	g := handlers.Games{GameService: gameService}
	g.Register(api)
	// get server with flag values filled out
	server = restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		return err
	}
	return nil
}
