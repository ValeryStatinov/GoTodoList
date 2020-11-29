FROM postgres:13

ENV PG_DATA=/pgdata:/var/lib/postgresql/data

ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_USER=todolist

COPY migrate.sh .
RUN chmod +x ./migrate.sh

COPY migrations/* /migrations/