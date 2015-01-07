
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS access_tokens (
  Id            serial primary key,
  Code          varchar(255),
  ExpiresIn     integer,
  Scope         varchar(255),
  RedirectUri   varchar(255),
  State         varchar(255),
  CreatedAt     timestamp not null default now(),
  UpdatedAt     timestamp not null default now()
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS access_tokens;
