package main

import (
	"fmt"

	"github.com/1995parham/loser-scraper/config"
	"github.com/1995parham/loser-scraper/parser"
	"github.com/1995parham/loser-scraper/scrap"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	cfg := config.New()

	sc := scrap.New(cfg.Target)
	rd, err := sc.Scrap()
	if err != nil {
		logrus.Fatalf("scrap failed: %s", err)
		return
	}
	defer func() {
		if err := rd.Close(); err != nil {
			logrus.Errorf("reader closed: %s", err)
		}
	}()
	ts, err := parser.ExtractTimeline(rd)
	if err != nil {
		logrus.Fatalf("parser failed: %s", err)
		return
	}
	for _, t := range ts {
		fmt.Printf("Tweet %d from %s at %s: %s\n", t.Index, t.User, t.At, t.Content)
	}
}
