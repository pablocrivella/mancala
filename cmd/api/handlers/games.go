package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pablocrivella/mancala/api/openapi-v2/models"
	"github.com/pablocrivella/mancala/api/openapi-v2/restapi/operations"
	openapi "github.com/pablocrivella/mancala/api/openapi-v2/restapi/operations/games"
	"github.com/pablocrivella/mancala/internal/games"
)

// Games represens a restful API resource
type Games struct {
	GameService games.Service
}

// Create handles the POST /v1/games endpoint
func (g Games) Create(params openapi.CreateGameParams) middleware.Responder {
	game, err := g.GameService.CreateGame(params.Body.Player1, params.Body.Player2)
	if err != nil {
		return &openapi.CreateGameBadRequest{
			Payload: &models.Problem{
				Type: "/problems/bad-request",
			},
		}
	}
	return &openapi.CreateGameCreated{
		Payload: &models.Game{
			ID:     game.ID.String(),
			Result: models.Result(game.Result),
			Turn:   models.Turn(game.Turn),
		},
	}
}

// Update handles the PUT /v1/games/:id endpoint
func (g Games) Update(params openapi.UpdateGameParams) middleware.Responder {
	return &openapi.UpdateGameOK{}
}

// Show handles the GET /v1/games/:id endpoint
func (g Games) Show(params openapi.ShowGameParams) middleware.Responder {
	return &openapi.ShowGameOK{}
}

// Register wires the handlers with the go-swagger generated API
func (g Games) Register(api *operations.MancalaAPI) {
	api.GamesCreateGameHandler = openapi.CreateGameHandlerFunc(g.Create)
	api.GamesUpdateGameHandler = openapi.UpdateGameHandlerFunc(g.Update)
	api.GamesShowGameHandler = openapi.ShowGameHandlerFunc(g.Show)
}
