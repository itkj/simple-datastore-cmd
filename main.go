package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/datastore"
)

type Entity struct {
	Value string
}

func main() {
	if len(os.Args) < 5 {
		printUsage()
		return
	}

	command := strings.ToLower(os.Args[1])
	if command != "get" && command != "set" {
		printUsage()
		return
	}

	datasetID := os.Args[2]
	kind := os.Args[3]
	key := os.Args[4]

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, datasetID, cloud.WithTokenSource(google.ComputeTokenSource("")))
	if err != nil {
		log.Fatal(err)
	}

	k := datastore.NewKey(ctx, kind, key, 0, nil)
	e := new(Entity)

	if command == "get" {
		if err := client.Get(ctx, k, e); err != nil {
			log.Fatal(err)
		}
		fmt.Println(e.Value)
	} else if command == "set" {
		if len(os.Args) != 6 {
			printUsage()
			return
		}
		value := os.Args[5]
		e.Value = value
		if _, err := client.Put(ctx, k, e); err != nil {
			log.Fatal(err)
		}
	}
}

func printUsage() {
	fmt.Println("usage:")
	fmt.Println(os.Args[0] + " get <datasetID> <kind> <key>")
	fmt.Println(os.Args[0] + " set <datasetID> <kind> <key> <value>")
}
