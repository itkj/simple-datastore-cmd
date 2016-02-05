# Simple Datastore CommandLine Tool

work on Google Compute Engine

## Scope

* https://www.googleapis.com/auth/datastore
* https://www.googleapis.com/auth/userinfo.email

## Build

```
GOOS=linux GOARCH=amd64 go build -o dscmd main.go
```

## Usage

```
./dscmd get <datasetID> <kind> <key>
./dscmd set <datasetID> <kind> <key> <value>
./dscmd wait <datasetID> <kind> <key> <value>  # waiting for stored value is <value>
```