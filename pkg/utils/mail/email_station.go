package mail

import (
	"context"
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

// Email represents an email message.
// To is a list of recipients.
// Subject is the subject of the email.
// HTML is the body of the email.
// Email 表示一封邮件。
// To 是收件人列表。
// Subject 是邮件的主题。
// HTML 是邮件的正文。
type Email struct {
	To      []string
	Subject string
	HTML    string
}

// EmailStationConfig represents the configuration of the email station.
// Address is the SMTP server address.
// Port is the SMTP server port.
// Username is the username of the SMTP server.
// Password is the password of the SMTP server.
// ConnPoolSize is the size of the connection pool.
// EmailStationConfig 表示邮件站的配置。
// Address 是 SMTP 服务器地址。
// Port 是 SMTP 服务器端口。
// Username 是 SMTP 服务器的用户名。
// Password 是 SMTP 服务器的密码。
// ConnPoolSize 是连接池的大小。
type EmailStationConfig struct {
	Address      string
	Port         int
	Username     string
	Password     string
	ConnPoolSize int
}

// EmailStation represents an email station.
// It sends emails concurrently.
// EmailStation 表示一个邮件站。
// 它并发地发送邮件。
type EmailStation struct {
	config *EmailStationConfig
	ctx    context.Context
	cancel context.CancelFunc
	ch     chan *Email
}

// NewEmailStation creates a new email station.
// config is the configuration of the email station.
// NewEmailStation 创建一个新的邮件站。
// config 是邮件站的配置。
func NewEmailStation(config EmailStationConfig) *EmailStation {
	ctx, cancel := context.WithCancel(context.Background())
	return &EmailStation{
		config: &config,
		ctx:    ctx,
		cancel: cancel,
		ch:     make(chan *Email),
	}
}

// Run starts the email station.
// Run 启动邮件站。
func (e *EmailStation) Run() {
	for i := 0; i < e.config.ConnPoolSize; i++ {
		go e.worker()
	}
}

// worker is the worker function of the email station.
// It sends emails.
// worker 是邮件站的工作函数。
// 它发送邮件。
func (e *EmailStation) worker() {
	for {
		select {
		case email := <-e.ch:
			m := gomail.NewMessage()
			m.SetHeader("From", e.config.Username)
			m.SetHeader("To", email.To...)
			m.SetHeader("Subject", email.Subject)
			m.SetBody("text/html", email.HTML)
			d := gomail.NewDialer(e.config.Address, e.config.Port, e.config.Username, e.config.Password)
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
			if err := d.DialAndSend(m); err != nil {
				log.Println(err)
			}
		case <-e.ctx.Done():
			return
		}
	}
}

// Send sends an email.
// Send 发送一封邮件。
func (e *EmailStation) Send(email *Email) {
	e.ch <- email
}

// Close closes the email station.
// Close 关闭邮件站。
func (e *EmailStation) Close() {
	e.cancel()
	close(e.ch)
}
