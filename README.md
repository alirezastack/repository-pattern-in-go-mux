# Service Name

To start the project you need to first initialize configuration file and set `FILE_CONFIG_NAME` to the name of your config file. For development set it to dev.
Also change your `dev.sample.json` file in configs folder to `dev.json`.

### Set log level
In order to set service logging level to debug you can pass -debug to `go run main.go` like below:

```shell
go run main.go -debug
```

*NOTE:* by default service log level is set to `INFO`, so you won't see debug logs.
