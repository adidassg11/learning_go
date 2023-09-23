module github.com/adidassg11/hello

go 1.21.1

require (
	github.com/adidassg11/greetings v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.9
)

// go.work can be used instead of adding replace directives to work across multiple modules
replace github.com/adidassg11/greetings => ../greetings
