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
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	//测试go语言模拟请求
	//req := curl.NewRequest()
	//resp, err := req.SetUrl("https://www.baidu.com").Get()
	//util.CheckErr(err)
	//fmt.Println(resp.Body)
	response, err := http.Get("https://www.baidu.com")
	util.CheckErr(err)
	res, err := ioutil.ReadAll(response.Body)
	util.CheckErr(err)
	data["msg"] = res
	data["code"] = 200
	data["name"] = names
	util.Return_json(w, data)
}
