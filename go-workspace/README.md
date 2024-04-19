# Go Workspace

This tutorial is for how to add multiple module in a single workspace module

#### Folder Structure

```bash
└── go-workspace/
    ├── main/
    │   ├── go.mod
    │   └── main.go 
    ├── module1/
    │   ├── go.mod
    │   └── text.go
    ├── module2/
    │   ├── go.mod
    │   ├── number.go
    │   └── module3/
    │       ├── go.mod
    │       └── time.go
    └── go.work
```

#### Create new module in go 

`cd main && go mod init github.com/arun-maersk/go-platform/go-workspace/main`

`cd module1 && go mod init github.com/arun-maersk/go-platform/go-workspace/module1`

`cd module2 && go mod init github.com/arun-maersk/go-platform/go-workspace/module2`

`cd module3 && go mod init github.com/arun-maersk/go-platform/go-workspace/module2/module3`

#### Problem Statement
Normally, we cannot able to import and call a function from one module to another module. 
<b>Example:</b> `text.go` having a `func Text(){}` which has not been call by `main.go`, since these two are different modules

#### Solution
To solve this we can go for `go workspace`. To create go workspace need to run the following command

`go work init ./main ./module1 ./module2 ./module2/module3`

(or)

`go work init ./main`
`go work use ./module1`
`go work use ./module2`
`go work use ./module2/module3`


This will create `go.work` file

```bash
go 1.22.1

use (
	./main
	./module1
	./module2
	./module2/module3
)
```

Now, we can able to import and call any functions between these modules

