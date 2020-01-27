## The `go` CLI

The `go <action>` cmd is quite consistent about the parameters. For example:

    go fix ...
    go build ...
    go fmt ...
    go test...

They all take `packages` as parameters to refer the package in a module
(assuming we are not using `GOPATH` anymore).

Some tips:

- All these cmds are module aware. Use `go list all` to find the package names.
- To understand what's `packages`, run `go help packages`.

Assuming you have the following set up

    # The folder is ~/Workspace/myrepro
    # The github package is mapped as github.come/xiejw/myrepro
    # The files under myrepro is

    myrepro
      -> lib  (no source files)
         -> lexer
            -> lexer.go
            -> lexer_test.go
         -> parser
            -> parser.go
            -> parser_test.go

Then you can do this

    go <action> github.com/xiejw/lunar/lib/...
    go <action> github.com/xiejw/lunar/lib/lexer
    go <action> github.com/xiejw/lunar/lib/parser

    # But the following is not allowed, given no source file under `lib`.
    go <action> github.com/xiejw/lunar/lib


