module github.com/ridwanulhoquejr/go-crashcourse/hello

go 1.22.3

replace github.com/ridwanulhoquejr/go-crashcourse/greetings => ../greetings

require (
	github.com/ridwanulhoquejr/go-crashcourse/greetings v0.0.0-00010101000000-000000000000
	github.com/ridwanulhoquejr/go-crashcourse/talking v0.0.0-00010101000000-000000000000
)

replace github.com/ridwanulhoquejr/go-crashcourse/talking => ../talking
