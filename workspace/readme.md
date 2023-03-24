### go work

```
go work init ./hello
```

The `go work init` command tells `go` to create a `go.work` file for a workspace
containing the modules in the `./hello` directory

The `go` command produces a go.work file that looks like this:
```
go 1.18

use ./hello
```

The `use` directive tells Go that the module in the `hello` directory should be
main modules when doing a build

### release

module release workflow documentation