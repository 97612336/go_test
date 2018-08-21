package company_pro

import (
	"net/http"
	"go_test/util"
	"io/ioutil"
)

func Get_one_specs(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})
	sql_str := "select i1 from specs where id=3112;"
	var names []string
	var name string
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	//测试网页请求
	response, err := http.Get("https://www.baidu.com")
	util.CheckErr(err)
	res, err := ioutil.ReadAll(response.Body)
	util.CheckErr(err)
	data["msg"] = res
	data["code"] = 200
	data["name"] = names
	util.Return_json(w, data)
}

type Type int

func Select_sql_rows(sql_str string, lists []Type, one_obj Type) {
	rows, err := util.DB.Query(sql_str)
	util.CheckErr(err)
	for rows.Next() {
		err := rows.Scan(&one_obj)
		util.CheckErr(err)
		lists = append(lists, one_obj)
	}
}
