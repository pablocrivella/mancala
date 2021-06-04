package engine

import (
	"errors"
	"fmt"
)

var (
	// ErrGameNotFound happens when the game cannot be found
	ErrGameNotFound = errors.New("game not found")

	// ErrSelectedEmptyPit happens when the selected pit is empty
	ErrSelectedEmptyPit = fmt.Errorf("%w: selected pit is empty", ErrInvalidPlay)

	// ErrGameIsDone happens when a play is placed after the game is finished
	ErrGameIsDone = fmt.Errorf("%w: game is already done", ErrInvalidPlay)

	// ErrSelectedInvalidPit happens when the selected pit is out of bounds (0-5)
	ErrSelectedInvalidPit = fmt.Errorf("%w: selected pit is invalid", ErrInvalidPlay)

	// ErrInvalidPlay happends when a play cannot be performed due to a specific error
	ErrInvalidPlay = errors.New("invalid play")
)
