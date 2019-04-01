package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/fuyao/service"
)

type Application struct {
	Setup *service.ServiceSetup
}

// 添加信息
func (app *Application) AddInfo(w http.ResponseWriter, r *http.Request) {
	var fy = service.Fuyao{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &fy)

	_, err := app.Setup.SaveInfo(fy)

	if err != nil {
		var str string
		str = "添加失败,失败原因为:" + err.Error()
		io.WriteString(w, str)
	}
	io.WriteString(w, "添加成功")
}

// 根据身份证号码查询信息
func (app *Application) FindByID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form["id"])
	result, err := app.Setup.FindFyInfoByCodeId(r.Form["id"][0])
	if err != nil {
		var str string
		str = "添加失败,失败原因为:" + err.Error()
		io.WriteString(w, str)
	}
	var edu = service.Fuyao{}
	json.Unmarshal(result, &edu)
	data, _ := json.Marshal(edu)
	w.Write(data)
}

//根据编码号删除信息
func (app *Application) DeleteInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form["id"])
	_, err := app.Setup.DelInfo(r.Form["id"][0])
	if err != nil {
		var str string
		str = "删除失败,失败原因为:" + err.Error()
		io.WriteString(w, str)
	}
	io.WriteString(w, "删除成功")
}
