-- name: GetGame :one
SELECT * FROM games
WHERE id = $1 LIMIT 1;

-- name: UpsertGame :one
INSERT INTO games (
  turn, result, board_side1, board_side2
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (id)
DO
  UPDATE SET turn = $1, result = $2, board_side1 = $3, board_side2 = $4
RETURNING *;