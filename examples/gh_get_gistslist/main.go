package main

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	ghgists "github.com/alcacoop/go-github-client/gists"
	"flag"
	"fmt"
)

func main() {
	help := flag.Bool("help", false, "Show usage")
	username := flag.String("u", "", "Specify Github user")
	password := flag.String("p", "", "Specify Github password")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || *username == "" || *password == ""{
		flag.Usage()
		return
	}

	ghc, err := ghclient.NewGithubClient(*username, *password, ghclient.AUTH_USER_PASSWORD)

	gistsc := ghgists.NewGists(ghc)

	res, err := gistsc.GetPublicGistsList()

	jr, err := res.Json()

	fmt.Printf("JSON: %v\nERROR: %v\n", jr, err)

	print("\n\nLOADING NEXT PAGE...\n\n")
	res, err = res.GetNextPage()

	jr, err = res.Json()

	fmt.Printf("JSON: %v\nERROR: %v\n", jr, err)
}