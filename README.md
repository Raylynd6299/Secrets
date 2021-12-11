# Secrets

A Backend Application

### Manage Dependency to go.mod

Update dependencies

```
	go mod tidy
```
### Autoindent and format code

For autoformat our sources, we only need to run one line for formating our code.

We only need to run the next line

```bash
	gofmt -w source.go
```

or

```bash
	gofmt -w path/to/your/application
```

## Conventions in this project

* Functions -> Use Pascal Case i.e.
```go
func HelloWorld()...
```
* Variables Use -> Camel Case i.e.
```go
routerHandler := nil
```
* Folders and init file have the same name i.e.

```bash
ls user/
# Output
	user.go
	...
	..
	.
```
