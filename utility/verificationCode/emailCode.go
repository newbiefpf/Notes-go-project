package verificationCode

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

func SendCode(emailStr string) (string, error) {
	e := email.NewEmail()
	mailUserName := "newbie0714@163.com" //邮箱账号
	mailPassword := "BMZKAWKHZSRMSVEY"   //邮箱授权码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000)) //发送的验证码
	Subject := "Note的验证码"                            //发送的主题
	e.From = "Note <newbie0714@163.com>"
	e.To = []string{emailStr}
	e.Subject = Subject
	t := time.Now().Format("2006-01-02 15:04:05")
	e.HTML = []byte(`<div>
	<div>
		尊敬的菜鸟，您好！
	</div>
	<div style="padding: 8px 40px 8px 50px;">
	<p>您于 ` + t + ` 提交的邮箱验证，本次验证码为<u><strong>` + code + `</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
	</div>
	<div>
	<p>此邮箱为系统邮箱，请勿回复。</p>
	</div>
	</div>`)
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	return code, err
}
