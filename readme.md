# Simple Datastore CommandLine Tool

work on Google Compute Engine

## Build

```
GOOS=linux GOARCH=amd64 go build -o dscmd main.go
```

## Usage

```
./dscmd get <datasetID> <kind> <key>
./dscmd set <datasetID> <kind> <key> <value>
```