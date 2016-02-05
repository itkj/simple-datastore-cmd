package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	if command != "get" && command != "set" && command != "wait" {
		printUsage()
		return
	}
	var value string
	if command == "set" || command == "wait" {
		if len(os.Args) != 6 {
			printUsage()
			return
		}
		value = os.Args[5]
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
			if err.Error() != "datastore: no such entity" {
				log.Fatal(err)
			}
		}
		fmt.Println(e.Value)
	} else if command == "set" {
		e.Value = value
		if _, err := client.Put(ctx, k, e); err != nil {
			log.Fatal(err)
		}
	} else if command == "wait" {
		errCount := 0
		for {
			if err := client.Get(ctx, k, e); err != nil {
				if err.Error() != "datastore: no such entity" {
					errCount++
					if errCount > 3 {
						log.Fatal(err)
					}
				}
			} else if e.Value == value {
				break
			} else {
				errCount = 0
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func printUsage() {
	fmt.Println("usage:")
	fmt.Println(os.Args[0] + " get <datasetID> <kind> <key>")
	fmt.Println(os.Args[0] + " set <datasetID> <kind> <key> <value>")
	fmt.Println(os.Args[0] + " wait <datasetID> <kind> <key> <value>")
}
