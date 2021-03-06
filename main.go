package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1995parham/loser-scraper/config"
	"github.com/1995parham/loser-scraper/mail"
	"github.com/1995parham/loser-scraper/parser"
	"github.com/1995parham/loser-scraper/scrap"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("18.20 at Sep 07 2016 7:20 IR721")

	var latest time.Time

	cfg := config.New()

	ml, err := mail.New(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.Username, cfg.Mail.Password)
	if err != nil {
		logrus.Fatalf("mailer initiation failed: %s", err)
	}

	c := cron.New()

	if _, err := c.AddFunc(fmt.Sprintf("@every %s", cfg.Period), func() {
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

		toSend := make([]parser.Timeline, 0)
		for _, t := range ts {
			if t.At.After(latest) {
				latest = t.At
				logrus.Infof("Tweet %d from %s at %s: %s\n", t.Index, t.User, t.At, t.Content)
				toSend = append(toSend, t)
			}
		}

		if err := ml.Send(cfg.Target, toSend, cfg.You, cfg.Mail.Username); err != nil {
			logrus.Errorf("send failed: %s", err)
		}
	}); err != nil {
		logrus.Fatalf("failed to register scraper: %s", err)
	}

	c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("18.20 As always ... left me alone")

	c.Stop()
}
