package company_pro

import (
	"net/http"
	"go_test/util"
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
	data["code"] = 200
	data["name"] = names
	
	util.Return_json(w, data)
}
