module test_server

go 1.22.1

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
	products v0.0.0-00010101000000-000000000000
)

replace products => ./products
