package handlers

import (
	"go_test/handlers/company_pro"
	"go_test/handlers/email"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/v1/get_one_specs", company_pro.Get_one_specs)
	http.HandleFunc("/v1/send", email.Send_one_email)
	//go语言设置静态文件目录
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/home/wangkun/static_dir"))))
}
