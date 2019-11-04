for file in ./database/migrations/*
do
  psql -d todolist -f $file
done
