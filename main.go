package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/1995parham/loser-scraper/config"
	"github.com/1995parham/loser-scraper/parser"
	"github.com/1995parham/loser-scraper/scrap"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	cfg := config.New()

	c := cron.New()

	_, err := c.AddFunc(fmt.Sprintf("@every %s", cfg.Period), func() {
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
			logrus.Errorf("parser failed: %s", err)
			return
		}
		for _, t := range ts {
			logrus.Infof("Tweet %d from %s at %s: %s\n", t.Index, t.User, t.At, t.Content)
		}
	})
	if err != nil {
		logrus.Fatalf("failed to register scraper: %s", err)
	}

	c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("18.20 As always ... left me alone")

	c.Stop()
}
