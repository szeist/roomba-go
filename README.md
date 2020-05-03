# roomba-go

Unofficial SDK for iRobot Roomba models.
This is a re-implementation of a part of the https://github.com/koalazak/dorita980 library. Many Thanks!

## CLI App

### Obtain credentials

1. Discover roomba and save the `Blid` field and the `IP Address`
   ```sh
   roomba-cli discover
   ```
2. Get roomba password
   ```sh
   roomba-cli get-password -host [IP Address]
   ```
   Follow the instructions and save the printed password.

These credentials can be passwd with the `-host`, `-user` and the `-password` flag for each command or can be loaded from the `ROOMBA_HOST`, `ROOMBA_USER` and the `ROOMBA_PASSWORD` environment variables.

Usage information:
```sh
roomba-cli --help
```

## SDK

### Usage example

```go
cfg := &config.Config{
  Address:  "[ROOMBA HOST]"
  User:     "[ROOMBA BLID]"
  Password: "[ROOMBA PASSWORD]"
  Debug:    false
}
client := roomba.New(cfg)
if err = client.Connect(); err != nil {
  log.Fatalf("Unable to connect to roomba: %v", err)
}

if err = client.SendCommand("clean"); err != nil {
  log.Fatalf("Unable to send command: %v", err)
}

client.Disconnect()
```

### Configuration

Client configuration can be loaded from environment variables

```go
cfg := config.NewFromEnv("ROOMBA_")
```

In this case the roomba will load the matching config attributes from environment variables. (e.g.: `ROOMBA_ADDRESS` ->  `Address`)

#### State logging

The configuration accepts a `StateWriter` field. If this field is set all state message (coordinates, status messages, signal strength, ...) will written in JSON format.