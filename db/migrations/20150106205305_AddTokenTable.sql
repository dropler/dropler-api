
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS access_tokens (
  Id            serial primary key,
  Token         varchar(255),
  UserID        integer,
  ClientID      varchar,
  ExpiresIn     integer,
  Scope         varchar(255),
  CreatedAt     timestamp not null default now(),
  UpdatedAt     timestamp not null default now()
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS access_tokens;
