package mail_test

import (
	"fmt"
	"sfw/pkg/utils/generator"
	"sfw/pkg/utils/mail"
	"testing"
	"time"
)

func TestMail(t *testing.T) {
	e := mail.NewEmailStation(mail.EmailStationConfig{
		Address:      "smtp.mxhichina.com",
		Port:         465,
		Username:     "fulifuli@sophisms.cn",
		Password:     "_222200316Cyk",
		ConnPoolSize: 4,
	})
	e.Run()
	defer e.Close()
	e.Send(&mail.Email{
		To:      []string{"2478096487@qq.com"},
		Subject: "Test",
		HTML:    fmt.Sprintf(mail.HTML, "fulifuli", generator.GenerateAlnumString(generator.AlnumGeneratorOption{Length: 6, UseNumber: true}), "fulifuli", "fulifuli"),
	})

	time.Sleep(20 * time.Second)
}
