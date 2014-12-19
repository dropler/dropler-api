
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS clients (
  Id            serial primary key,
  Name          varchar(255),
  ClientID      varchar(255),
  ClientSecret  varchar(255),
  RedirectURI   varchar(255),
  CreatedAt     timestamp not null default now(),
  UpdatedAt     timestamp not null default now()
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS clients;
