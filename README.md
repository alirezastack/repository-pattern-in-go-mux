# Service Name

### Project Structure
 + configs  # all dev/prod configuration
   + dev.sample.json
   + env.go
 + controllers  # core business logic for routes
   + user_controller.go
 + helpers  # utility functions like validators/formatters/encryptors/...
   + formatter.go
   + request.go
   + response.go
 + models  # all DB related models/structs
   + user_model.go
 + responses
   + user_response.go
 + routes  # correlate routes to controller methods 
   + user_route.go
 + store  # repositories implementation and store interface
   + mongo_store.go
   + store.go  # generic interface for the service store/repository
 + .gitignore
 + go.mod
 + main.go
 + README.md

To start the project you need to first initialize configuration file and set `FILE_CONFIG_NAME` to the name of your config file. For development set it to dev.
Also change your `dev.sample.json` file in configs folder to `dev.json`.

### Log
[`Zerolog`](https://github.com/rs/zerolog) API is designed to provide both a great developer experience and stunning performance. Its unique chaining API allows zerolog to write JSON (or CBOR) log events by avoiding allocations and reflection.

In order to set service logging level to debug you can pass -debug to `go run main.go` like below:

```shell
go run main.go -debug
```

*NOTE:* by default service log level is set to `INFO`, so you won't see debug logs.


### Graceful shutdown
Service is able to receive shutdown signals via channels including `syscall.SIGINT`, `syscall.TERM`. 
To increase graceful timeout the service provides a flag called `graceful-timeout` in seconds. So we could wait for threads to finish their job before exiting. 

You could provide this flag as a duration in Go i.e. 30s, 1m and so on.

```bash
go run main.go -graceful-timeout 30s
```