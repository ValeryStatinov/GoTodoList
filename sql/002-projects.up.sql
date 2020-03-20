create table if not exists projects (
  id serial primary key,
  name text not null
);

-- insert into projects (name) values ('First project');

alter table tasks
  add column if not exists projectId
  int not null
  default (1)
  references projects(id);
