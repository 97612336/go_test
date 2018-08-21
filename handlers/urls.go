package handlers

import (
	"net/http"
	"go_test/handlers/company_pro"
)

func MyUrls() {
	http.HandleFunc("/v1/get_one_specs", company_pro.Get_one_specs)
}
