# DEV - Commands
## MUST CD TO SQL FOLDER FIRST
`goose postgres "host=localhost port=5432 user=postgres password=%s dbname=go_learning"`
## ONLY IN CMD PROPT
sqlc   `docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate`