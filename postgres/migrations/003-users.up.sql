create table if not exists users (
  id serial primary key,
  name text not null unique,
  password text not null
);

alter table projects
  add column if not exists userId
  int not null
  default(1)
  references users(id);