create table if not exists tasks (
  id serial primary key,
  name text not null,
  description text not null,
  project_id int not null references projects(id),
  priority int not null
  -- user_id int not null references users(id)
);

-- INSERT INTO tasks (NAME, DESCRIPTION, PROJECT_ID, USER_ID) VALUES ('Do server', 'Write web server on GO', 1, 1);
