#!/bin/sh

./bin/goose -dir ./migrations postgres "$DATABASE_URL" up
./bin/ocp-problem-api -host ''