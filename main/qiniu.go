package main

import (
	"go_test/util"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type Qiniu_key struct {
	Ak string
	Sk string
}

func Get_qiniiu_keys() Qiniu_key {
	home_paht := util.Get_home_path()
	config_path := home_paht + "/conf/qiniu_conf"
	data, err := ioutil.ReadFile(config_path)
	util.CheckErr(err)
	var keys Qiniu_key
	json.Unmarshal(data, &keys)
	fmt.Println(keys)
	return keys
}

func main() {
	keys := Get_qiniiu_keys()
	accesskey := keys.Ak
	secretkey := keys.Sk
	bucket := "webfile"
	putPkicy := storage.PutPolicy{
		Scope: bucket,
	}
	//获取上传凭证
	mac := qbox.NewMac(accesskey, secretkey)
	upToken := putPkicy.UploadToken(mac)
	fmt.Println(upToken)
}
