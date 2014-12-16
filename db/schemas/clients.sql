-- CLIENTS
DROP TABLE clients;

CREATE TABLE clients (
  Id            serial primary key,
  Name          varchar(255),
  ClientID      varchar(255),
  ClientSecret  varchar(255),
  RedirectURI   varchar(255),
  CreatedAt     timestamp not null default now(),
  UpdatedAt     timestamp not null default now()
);

