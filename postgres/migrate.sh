#!/bin/bash
for file in ./migrations/*
do
  psql -U todolist -d todolist -f $file
done