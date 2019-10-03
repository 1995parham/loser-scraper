package main

import (
	"fmt"

	"github.com/1995parham/loser-scraper/config"
	"github.com/1995parham/loser-scraper/parser"
	"github.com/1995parham/loser-scraper/scrap"
)

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	cfg := config.New()

	sc := scrap.New(cfg.Target)
	rd, err := sc.Scrap()
	if err != nil {
		fmt.Println(err)
		return
	}
	parser.ExtractTimeline(rd)
}
