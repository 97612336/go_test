package email

import (
	"fmt"
	"go_test/util"
	"gopkg.in/gomail.v2"
	"net/http"
)

func Send_one_email(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		to_mail := util.Get_argument(r, "mail")
		subject := util.Get_argument(r, "subject")
		text := util.Get_argument(r, "text")
		secret := util.Get_argument(r, "secret")
		//获取邮箱配置信息
		email_info := util.Get_email_info()
		if secret == email_info.Secret {
			m := gomail.NewMessage()
			m.SetAddressHeader("From", email_info.Account, "小小卡") // 发件人
			m.SetHeader("To", // 收件人
				m.FormatAddress(to_mail, "亲"),
			)
			m.SetHeader("Subject", subject) // 主题
			m.SetBody("text/html", text)    // 正文
			d := gomail.NewDialer(email_info.Server, email_info.Port, email_info.Account, email_info.Passwd)
			if err := d.DialAndSend(m); err != nil {
				fmt.Println(err)
			}
		}
	}
}
