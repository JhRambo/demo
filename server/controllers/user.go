package controllers

import (
	"demo/server/models"
	"demo/server/tools"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

type User struct {
}

type Res struct {
	Success bool
	Message string
}

type Cmd struct {
	Id    int
	Phone string
	Code  string
}

/*
	结合 msgpack json

查询指定id的数据
*/
func (this *User) GetById(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read failed:", err)
	}
	cmd := &Cmd{}
	err = msgpack.Unmarshal(b, cmd)
	if err != nil {
		log.Println("json format error:", err)
	}
	user := models.User{}
	// id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	id := cmd.Id
	tools.DB.First(&user, id)
	json, _ := json.Marshal(user)
	w.Write(json)
}

// 查询所有数据
func (this User) List(w http.ResponseWriter, r *http.Request) {
	user := []models.User{}
	tools.DB.Find(&user)
	json, _ := json.Marshal(user)
	w.Write(json)
}

// 删除
func (this *User) Del(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body) //数据流
	if err != nil {
		log.Println("err=>", err)
	}
	defer r.Body.Close()
	cmd := &Cmd{}
	json.Unmarshal(b, cmd)
	// id, _ := strconv.Atoi(r.PostFormValue("id"))
	id := cmd.Id
	if id == 0 {
		res := &Res{
			Success: false,
			Message: "删除失败",
		}
		json, _ := json.Marshal(res) //[]byte类型
		w.Write(json)
		return
	}
	user := models.User{}
	tools.DB.Where("id = ?", id).Delete(&user)
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			res := &Res{
				Success: true,
				Message: "删除成功",
			}
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}()
}

// 新增
func (this User) Add(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.Atoi(r.PostFormValue("age"))
	user := models.User{
		Name:     r.PostFormValue("name"),
		Phone:    r.PostFormValue("name"),
		Age:      age,
		Password: r.PostFormValue("password"),
	}
	tools.DB.Create(&user)
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			res := &Res{
				Success: true,
				Message: "添加成功",
			}
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}()
}

// 更新
func (this User) Upd(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	if id == 0 {
		res := &Res{
			Success: false,
			Message: "更新失败",
		}
		json, _ := json.Marshal(res) //[]byte类型
		w.Write(json)
		return
	}
	age, _ := strconv.Atoi(r.PostFormValue("age"))
	name := r.PostFormValue("name")
	phone := r.PostFormValue("phone")
	password := r.PostFormValue("password")
	user := models.User{
		Name:     name,
		Phone:    phone,
		Age:      age,
		Password: password,
	}
	tools.DB.Model(&user).Where("id = ?", id).Updates(user)
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			res := Res{
				Success: true,
				Message: "更新成功",
			}
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}()
}
