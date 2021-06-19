# cl

`cl` clones git repositories into nested folders like [GOPATH](https://golang.org/doc/gopath_code#Workspaces) and outputs the path of the cloned directory. Example:

```
# clone to /src/github.com/example-org/example-repo
cl -dir /src https://github.com/example-org/example-repo
```

You can install `cl` as a [binary release](https://github.com/felixge/cl/releases) or using Go:

```
go install github.com/felixge/cl@latest
```

The best way to use `cl` is to configure a bash function that specifies your preferred installation `-dir` and changes into the cloned directory on success:

```
cl() {
  path=$(command cl -dir "$GOPATH/src" "$1")
  cd "$path"
}
```

After that any of the following commands can be used to clone a repo and change into its directory:

```
cl https://github.com/example-org/example-repo.git
cl https://github.com/example-org/example-repo
cl git@github.com:example-org/example-repo.git
```

# License

cl is licensed under the MIT license.
