-- USERS
DROP TABLE users;

CREATE TABLE users (
  Id              serial primary key,
  Name            varchar(255),
  Email           varchar(255),
  HashedPassword  varchar(255),
  CreatedAt       timestamp not null default now(),
  UpdatedAt       timestamp not null default now()
);

