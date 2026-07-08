export GOOSE_DRIVER := "sqlite"
export GOOSE_DBSTRING := "./database.db"
export GOOSE_MIGRATION_DIR := "./migrations"

[parallel]
dev: air sql-watch

air:
  go tool air

sql-watch:
  watchexec -e sql just sqlc generate

goose *ARGS:
  go tool goose {{ARGS}}

sqlc *ARGS:
  go tool sqlc {{ARGS}}
