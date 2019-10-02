/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 03-10-2019
 * |
 * | File Name:     parser.go
 * +===============================================
 */

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func main() {
	f, err := os.Open("creep.html")
	if err != nil {
		return
	}
	timelines, err := ExtractTimeline(f)
	if err != nil {
		return
	}

	for _, t := range timelines {
		fmt.Printf("Tweet %d from %s at %s: %s\n", t.Index, t.User, t.At, t.Content)
	}
}

// Timeline represents an element on twitter timeline
type Timeline struct {
	User    string
	Index   int
	Content string
	At      time.Time
}

// ExtractTimeline extracts the timeline from the given user page
func ExtractTimeline(input io.Reader) ([]Timeline, error) {
	var timelines []Timeline

	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return timelines, err
	}

	doc.Find("li.stream-item").Each(func(i int, s *goquery.Selection) {
		var timeline Timeline

		// username
		timeline.User = s.Find("strong.fullname").Text()

		// index
		timeline.Index = i

		// content in html
		content, err := s.Find("p").Html()
		if err != nil {
			logrus.Errorf("tweet content fetch error: %s", err)
		}
		timeline.Content = content

		// timestamp
		secs, err := strconv.Atoi(s.Find("span._timestamp").AttrOr("data-time", "0"))
		if err != nil {
			logrus.Errorf("tweet timestamp parse error: %s", err)
		}
		timeline.At = time.Unix(int64(secs), 0)

		timelines = append(timelines, timeline)
	})

	return timelines, nil
}
