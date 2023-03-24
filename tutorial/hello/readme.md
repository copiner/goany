### Compile and install the application

go build

#### install

Discover the Go install path where the go command will install the currrent package

```
go list -f '{{.Target}}'
```

Add the Go install directory to your systems's shell path eg.
```
export PATH=$PATH:~/go/bin
```

Once you've updated the shell path, run the go install command to compile and install the package

```
go install
```

Run your application by simply typing its name.
```
hello
```

### module
    
```
cd greetings

go mod init example.com/greetings

cd hello

go mod init example.com/hello

go mod edit -replace example.com/greetings=../greetings

go mod tify
```