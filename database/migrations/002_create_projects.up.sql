create table if not exists projects (
  id serial primary key,
  name text not null
);

-- insert into projects (name) values ('project 1');