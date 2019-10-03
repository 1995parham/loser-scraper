package scrap

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Scrapper scrapes twitter page and return its html
type Scrapper struct {
	r *resty.Client
}

// New creates a new scapper for given user
func New(u string) Scrapper {
	r := resty.New()
	r.HostURL = fmt.Sprintf("https://twitter.com/%s", u)
	return Scrapper{r}
}

// Scrap scraps twitter page
func (s Scrapper) Scrap() (io.Reader, error) {
	logrus.Infof("sending request for fetching twitter timeline")
	resp, err := s.r.R().Get("")
	if err != nil {
		return nil, err
	}
	logrus.Infof("fetch was done successfully")

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("twitter returns %d", resp.StatusCode())
	}

	return resp.RawBody(), nil
}
