# cl

`cl` clones git repositories into nested folders like [GOPATH](https://golang.org/doc/gopath_code#Workspaces) and outputs the path of the cloned directory. Example:


```
cl https://github.com/foo/bar
```

Is roughly equivalent to:

```
mkdir -p ~/src/github.com/foo/bar
cd ~/src/github.com/foo/bar
git clone https://github.com/foo/bar .
```

You can install `cl` as a [binary release](https://github.com/felixge/cl/releases) or using Go:

```
go install github.com/felixge/cl@latest
```

In order for the example above to work, you'll need to configure a bash function in your profile that specifies your preferred installation `-dir` and changes into the cloned directory on success:

```
cl() {
  cloned=$(command cl -dir "$GOPATH/src" "$1")
  cd "$cloned"
}
```

After that any of the following commands can be used to clone a repo and change into its directory:

```
cl https://github.com/foo/bar.git
cl https://github.com/foo/bar
cl git@github.com:foo/bar.git
```

Without the bash function, you'll have to use a slightly more verbose command:

```
cd $(cl -dir ~/src https://github.com/foo/bar)
```

# License

cl is licensed under the MIT license.
