# thebrag-api
Easy management of your work log for year end reviews

## Golang Version:
1.20

## Install dependency packages:
```
go mod
```

## Run on Local:

```
go run main.go
```

## Run on VM:
1. Run the SCP command to copy the folder to the server.
2. `cd thebrag-api`
3. `go run main.go&`

#### To check if the application is listening to 8080 port:

`netstat -lntu`

To get the PID of the application:

`lsof -wni tcp:8080`

To kill the application:
`kill -9 <PID>`

