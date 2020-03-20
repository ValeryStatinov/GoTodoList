create table if not exists tasks (
  id serial primary key,
  name text not null,
  description text not null,
  priority int not null,
  completed boolean not null
);

-- INSERT INTO tasks (name, description, priority, completed) VALUES ('Do server', 'Write web server on GO', 1, false);
