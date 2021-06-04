package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/pablocrivella/mancala/internal/engine"
	"github.com/pablocrivella/mancala/internal/postgres/sqlc"
)

type GameStore struct {
	db *sql.DB
}

func NewGameStore(db *sql.DB) GameStore {
	return GameStore{
		db: db,
	}
}

func (s GameStore) Save(ctx context.Context, g engine.Game) (*engine.Game, error) {
	queries := sqlc.New(s.db)
	params := sqlc.UpsertGameParams{
		ID:     g.ID,
		Turn:   int64(g.Turn),
		Result: int64(g.Result),
	}
	side1, err := boardSideToJSONRawMessage(g.BoardSide1)
	if err != nil {
		return nil, err
	}
	side2, err := boardSideToJSONRawMessage(g.BoardSide2)
	if err != nil {
		return nil, err
	}
	params.BoardSide1 = side1
	params.BoardSide2 = side2
	record, err := queries.UpsertGame(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert game. %w", err)
	}
	game, err := dbRecordToGame(record)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (s GameStore) Find(ctx context.Context, id string) (*engine.Game, error) {
	queries := sqlc.New(s.db)
	record, err := queries.GetGame(ctx, uuid.MustParse(id))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, engine.ErrGameNotFound
		default:
			return nil, err
		}
	}
	game, err := dbRecordToGame(record)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func boardSideToJSONRawMessage(b engine.BoardSide) (json.RawMessage, error) {
	e, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json raw msg: %e", err)
	}
	return json.RawMessage(e), nil
}

func jsonRawMessageToBoardSide(m json.RawMessage) (engine.BoardSide, error) {
	var side engine.BoardSide
	err := json.Unmarshal(m, &side)
	if err != nil {
		return side, err
	}
	return side, nil
}

func dbRecordToGame(r sqlc.Game) (engine.Game, error) {
	var game engine.Game
	game.ID = r.ID
	game.Turn = engine.Turn(r.Turn)
	game.Result = engine.Result(r.Result)
	side1, err := jsonRawMessageToBoardSide(r.BoardSide1)
	if err != nil {
		return game, err
	}
	game.BoardSide1 = side1
	side2, err := jsonRawMessageToBoardSide(r.BoardSide2)
	if err != nil {
		return game, err
	}
	game.BoardSide2 = side2
	return game, nil
}
