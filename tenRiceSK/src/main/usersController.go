package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResultData struct {
	Code int         `json:"code" bson:"code"`
	Msg  string      `json:"msg" bson:"msg"`
	Data interface{} `json:"data" bson:"data"`
}

type BaseUser struct {
	U_uid          int    `json:"u_uid" bson:"u_uid"`
	U_nick_name    string `json:"u_nick_name" bson:"u_nick_name"`
	U_phone_number string `json:"u_phone_number" bson:"u_phone_number"`
	U_email        string `json:"u_email" bson:"u_email"`
}

type usersController struct {
}

func (this *usersController) UsersAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println("this is format")
	//err := r.ParseForm()
	//if err != nil {
	//	OutputJson(w, 0, "params errors", nil)
	//	return
	//}
	//
	//admin_name := r.FormValue("admin_name")
	//admin_password := r.FormValue("admin_password")
	//f
	//if admin_name == "" || admin_password == "" {
	//	OutputJson(w, 0, "params errors", nil)
	//	return
	//}

	//db := mysql.New("tcp", "", "192.168.100.166", "root", "test", "webdemo")
	//if err := db.Connect(); err != nil {
	//	log.Println(err)
	//	OutputJson(w, 0, "数据库操作失败", nil)
	//	return
	//}
	//defer db.Close()
	//
	//rows, res, err := db.Query("select * from webdemo_admin where admin_name = '%s'", admin_name)
	//if err != nil {
	//	log.Println(err)
	//	OutputJson(w, 0, "数据库操作失败", nil)
	//	return
	//}
	//
	//name := res.Map("admin_password")
	//admin_password_db := rows[0].Str(name)

	//if admin_password_db != admin_password {
	//	OutputJson(w, 0, "密码输入错误", nil)
	//	return
	//}

	// 存入cookie,使用cookie存储
	//cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
	//cookie := http.Cookie{Name: "admin_name", Value: "alexluan", Path: "/"}
	//http.SetCookie(w, &cookie)

	var user = BaseUser{}
	user.U_uid = 317
	user.U_nick_name = "alexluan"
	user.U_email = "326083325@qq.com"
	user.U_phone_number = "18817322819"

	var users = []BaseUser{user, user} //定义数组的时候，直接初始化
	users = append(users, user)

	var response = ResultData{}
	response.Code = 0
	response.Msg = "success"
	response.Data = user
	fmt.Println("users:::", users)
	OutputUsersJson(w, 210, "successg", users)
	return
}

func OutputUsersJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &ResultData{ret, reason, i}
	b, err := json.Marshal(out) //json.Marshal(out)
	if err != nil {
		return
	}

	m, n := w.Write(b)
	println(m)
	println(n)
}
