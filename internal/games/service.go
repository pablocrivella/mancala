package games

import (
	"context"

	"github.com/pablocrivella/mancala/internal/engine"
)

// Service is a games context domain service
type Service struct {
	GameStore interface {
		Find(context.Context, string) (*engine.Game, error)
		Save(context.Context, engine.Game) (*engine.Game, error)
	}
}

func (s Service) CreateGame(player1, player2 string) (*engine.Game, error) {
	g := engine.NewGame(player1, player2)
	sg, err := s.GameStore.Save(context.Background(), g)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s Service) FindGame(id string) (*engine.Game, error) {
	g, err := s.GameStore.Find(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s Service) ExecutePlay(gameID string, pitIndex int64) (*engine.Game, error) {
	g, err := s.FindGame(gameID)
	if err != nil {
		return nil, err
	}

	err = g.PlayTurn(pitIndex)
	if err != nil {
		return nil, err
	}

	sg, err := s.GameStore.Save(context.Background(), *g)
	if err != nil {
		return nil, err
	}
	return sg, nil
}
