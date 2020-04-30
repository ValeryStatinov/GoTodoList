alter table users
  alter column password type bytea using password::bytea;
