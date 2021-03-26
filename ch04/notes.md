- custom packages are shared by putting them in `~/go/src` which is kinda weird...
- working around committing stuff outside the workspace by using symbolic links
```
ln -s greeting ~/go/src/greeting
```
- if you create sub-dirs in your custom libs, their import paths will match
```
~/go/src/greeting/dansk/danske.go       # import path greeting/dansk
~/go/src/greeting/deutsch/deutsch.go    # import path greeting/deutsch
```

- format for constants
```
const MyConst int = 5
```
- because the constant is Capitalized, it'll be exported from it's package. Also weird.

- `go install $BIN` compiles source code in `$GOPATH/src/$BIN` and puts the binary in `$GOPATH/bin`
- `go get github.com/user/lib` will download the code to `$GOPATH/src/github/user/lib`, making it available for local use in imports: 
    `import github.com/user/lib` and use as `lib.ExportedFunction()`

- `go doc $PACKAGE_NAME` for in-line docs, `go doc $PACKAGE_NAME.function` also works
- docstrings! the line right above package declaration and/or function declaration is read by `go doc`

