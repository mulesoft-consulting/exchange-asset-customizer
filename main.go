package main

import (
	"context"
	"flag"
	"fmt"

	auth "github.com/mulesoft-consulting/cloudhub-client-go/authorization"
)

var (
	ANYPOINT_USERNAME string
	ANYPOINT_PWD      string
	ANYPOINT_ORG      string
	ANYPOINT_REGION   string
	CATEGORIES_FILE   string
	TAGS_FILE         string
	ATTRIBUTES_FILE   string
	SERVER_INDEX      int
	ANYPOINT_ORG_KEY  anypointContextKey
)

// func getenv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

func init() {
	// ANYPOINT_USERNAME = getenv("ANYPOINT_USERNAME", "")
	// ANYPOINT_PWD = getenv("ANYPOINT_PWD", "")
	// ANYPOINT_ORG = getenv("ANYPOINT_ORG", "")
	// ANYPOINT_REGION = getenv("ANYPOINT_REGION", "eu")
	// CATEGORIES_FILE = getenv("CATEGORIES_FILE", "")
	// TAGS_FILE = getenv("TAGS_FILE", "")

	flag.StringVar(&ANYPOINT_USERNAME, "u", "", "anypoint username")
	flag.StringVar(&ANYPOINT_PWD, "p", "", "anypoint password")
	flag.StringVar(&ANYPOINT_ORG, "o", "", "anypoint organization id")
	flag.StringVar(&ANYPOINT_REGION, "r", "eu", "anypoint region, by default eu (eu/us)")
	flag.StringVar(&CATEGORIES_FILE, "catfile", "", "path to categories csv file")
	flag.StringVar(&TAGS_FILE, "tagfile", "", "path to tags csv file")
	flag.StringVar(&ATTRIBUTES_FILE, "attrfile", "", "path to attributes csv file")
	flag.Parse()

	if ANYPOINT_ORG == "" {
		panic("org parameter is missing. --help for more information")
	} else if ANYPOINT_USERNAME == "" {
		panic("anypoint username parameter is missing. --help for more information")
	} else if ANYPOINT_PWD == "" {
		panic("anypoint password parameter is missing. --help for more information")
	}

	SERVER_INDEX = cplane2serverindex(ANYPOINT_REGION)
	ANYPOINT_ORG_KEY = anypointContextKey("orgId")
}

/*
	returns the server index depending on the control plane name
	if the control plane is not recognized, returns -1
*/
func cplane2serverindex(cplane string) int {
	if cplane == "eu" {
		return 1
	} else if cplane == "us" {
		return 0
	}
	return -1
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, auth.ContextServerIndex, SERVER_INDEX)
	ctx = context.WithValue(ctx, ANYPOINT_ORG_KEY, ANYPOINT_ORG)

	var client *ExchangeClient
	if TAGS_FILE != "" || CATEGORIES_FILE != "" || ATTRIBUTES_FILE != "" {
		client = newExchangeClient(ctx, ANYPOINT_USERNAME, ANYPOINT_PWD)
		if err := client.login(); err != nil {
			panic(err.Error())
		}
	} else {
		fmt.Println("No file provided. Use --help for more information")
	}

	if TAGS_FILE != "" {
		fmt.Printf("Process tags based on file %s \n", TAGS_FILE)
		if err := client.handleTags(TAGS_FILE); err != nil {
			panic(err.Error())
		}
	}

	if CATEGORIES_FILE != "" {
		fmt.Printf("Process categories based on file %s \n", CATEGORIES_FILE)
		if err := client.handleCategories(CATEGORIES_FILE); err != nil {
			panic(err.Error())
		}
	}

	if ATTRIBUTES_FILE != "" {
		fmt.Printf("Process attributes based on file %s \n", ATTRIBUTES_FILE)
		if err := client.handleAttributes(ATTRIBUTES_FILE); err != nil {
			panic(err.Error())
		}
	}

}
