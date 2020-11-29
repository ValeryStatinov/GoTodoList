create table if not exists tasks (
  id serial primary key,
  name text not null,
  description text not null,
  priority int not null,
  completed boolean not null
);
