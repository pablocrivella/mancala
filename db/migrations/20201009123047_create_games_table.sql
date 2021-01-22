-- migrate:up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS games (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  turn BIGINT NOT NULL,
  result BIGINT NOT NULL,
  board_side1 JSONB NOT NULL DEFAULT '{}'::jsonb,
  board_side2 JSONB NOT NULL DEFAULT '{}'::jsonb
);

-- migrate:down

DROP TABLE IF EXISTS games;
DROP EXTENSION IF EXISTS "uuid-ossp";
