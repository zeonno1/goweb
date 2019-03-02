package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
	Age  int
}

func init() {
	//orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:fibo112358@/test?charset=utf8", 30)
	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}
func main() {
	/*
		orm := orm.NewOrm()
		user := User{Name: "slene"}
		id, err := orm.Insert(&user)
		fmt.Printf("ID:%d,ERR:%v\n", id, err)
	*/
	/*
		orm := orm.NewOrm()
		user := User{Name: "slene"}
		user.Name = "astaxie"
		user.Id = 1
		num, err := orm.Update(&user)
		fmt.Printf("NUM:%d, ERR:$v", num, err)
	*/
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable(user)
	qs.Filter("id", 1)
	//users, _ := qs.Filter("id", 1).All(&user)
	qs.Filter("age", 18)
	fmt.Println(user)

}
