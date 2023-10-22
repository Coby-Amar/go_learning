# DEV - Commands
## debug app on vs code with `F5` after runing the build command below
### build and move app `go build -C ./main  -v && mv ./main/main.exe ./go_learning.exe`
## run the app `./go_learning.exe`

## MUST CD TO SQL FOLDER FIRST
`goose postgres "host=localhost port=5432 user=postgres password=%s dbname=go_learning"`
## ONLY IN CMD PROPT
sqlc   `docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate`