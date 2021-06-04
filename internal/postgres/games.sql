-- name: GetGame :one
SELECT *
FROM games
WHERE id = $1
LIMIT 1;
--
-- name: UpsertGame :one
INSERT INTO games (id, turn, result, board_side1, board_side2)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO
UPDATE
SET id = $1,
  turn = $2,
  result = $3,
  board_side1 = $4,
  board_side2 = $5
RETURNING *;
--