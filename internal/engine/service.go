package engine

import (
	"context"
)

type gameStore interface {
	Find(context.Context, string) (*Game, error)
	Save(context.Context, Game) (*Game, error)
}

// Service is a games context domain service
type Service struct {
	GameStore gameStore
}

func (s *Service) CreateGame(ctx context.Context, player1, player2 string) (*Game, error) {
	g := NewGame(player1, player2)
	sg, err := s.GameStore.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s *Service) FindGame(ctx context.Context, id string) (*Game, error) {
	g, err := s.GameStore.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *Service) ExecutePlay(ctx context.Context, gameID string, pitIndex int64) (*Game, error) {
	g, err := s.FindGame(ctx, gameID)
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
