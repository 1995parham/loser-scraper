package mail

import (
	"bytes"
	"html/template"

	"github.com/1995parham/loser-scraper/parser"
	gomail "github.com/go-mail/mail"
)

// Mailer sends emails.
type Mailer struct {
	d    *gomail.Dialer
	tmpl *template.Template
}

// info contains the required information for creating an eamil message
type info struct {
	Timelines []parser.Timeline
	Target    string
}

// New creates a new mailer
func New(host string, port int, username string, password string) (*Mailer, error) {
	d := gomail.NewDialer(host, port, username, password)
	d.StartTLSPolicy = gomail.MandatoryStartTLS

	tmpl, err := template.New("mail").Parse(tmplString)
	if err != nil {
		return nil, err
	}

	return &Mailer{
		d:    d,
		tmpl: tmpl,
	}, nil
}

// Send sends given timelines based on configured mail server
func (m *Mailer) Send(t string, ts []parser.Timeline, to string, from string) error {
	buf := bytes.NewBufferString("")
	if err := m.tmpl.Execute(buf, info{
		Timelines: ts,
		Target:    t,
	}); err != nil {
		return err
	}

	ms := gomail.NewMessage()
	ms.SetHeader("From", from)
	ms.SetHeader("To", to)
	ms.SetHeader("Subject", "Scraper for Losers")
	ms.SetBody("text/html", buf.String())

	if err := m.d.DialAndSend(ms); err != nil {
		return err
	}
	return nil
}
