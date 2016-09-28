package main

import (
	"flag"
	"fmt"
	"github.com/matthpl/goscrape"
	"os"
)

const usage = `cmd: scrap sites with given pattern.

Usage: %s
`

var (
	sourcesPath = flag.String("sources", "../sources.json", "path of url targets json config file")
	matchesPath = flag.String("matches", "../matches.json", "path of regex pattern json config file")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	sourcesConfig, _ := goscrape.GetSourcesConfig(*sourcesPath)
	matchesConfig, _ := goscrape.GetMatchesConfig(*matchesPath)
	matcher, _ := goscrape.NewScrapeMatcher()
	strs, _ := matcher.FindMatch(sourcesConfig, matchesConfig)
	fmt.Println(strs)
}
