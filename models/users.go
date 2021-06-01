package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

/*
字段结构体必须是大写可访问
*/
type Users struct {
	Id       int
	Username string `json:"user_name_string"`
	Password string
	Sex      int
	Role     *UserRole `orm:"rel(fk)"` // 一对多：Role(一个角色)对应多个User(多个用户)。users 表的外键：role_id 对应 user_role 表的 id
}

type UserRole struct {
	Id    int      `orm:"pk"`
	Name  string   `orm:"column(role_name)"`
	Desc  string   `orm:"column(role_desc)"`
	Users []*Users `orm:"reverse(many)"` // 一对多(反向)
}

func FindUserByUserName(Username string) (*Users, error) {
	o := orm.NewOrm()
	
	/*
		查询的其中两种方法
	*/
	// Read() 默认根据主键查询
	/*user := &Users{Username:Username}
	err := o.Read(user, "Username") // 默认根据主键查询
	return user, err*/
	
	// One()
	user := &Users{}
	// o.QueryTable("Users").Filter("username", Username).One(user)
	err := o.QueryTable(user).Filter("username", Username).One(user)
	
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Println("Returned Multi Rows Not One")
		return nil, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Println("Not row found")
		return nil, err
	}
	return user, nil
}

func AllUser() ([]*Users, error) {
	o := orm.NewOrm()
	var u []*Users

	// _, err := o.QueryTable("Users").RelatedSel().All(&u) // 方式1
	_, err := o.QueryTable(&Users{}).RelatedSel().All(&u) // 方式2

	return u, err
}

func AllUserRole() (*UserRole, error) {
	o := orm.NewOrm()
	var u *UserRole
	
	// _, err := o.QueryTable("UserRole").RelatedSel().All(&u) // 方式1
	_, err := o.QueryTable(&UserRole{}).RelatedSel().All(&u) // 方式2

	return u, err
}

func FindUserRoleByUserId(Id int) (*UserRole, error) {
	o := orm.NewOrm()
	r := &UserRole{}
	err := o.QueryTable(r).Filter("Id", Id).One(r)
	if err != nil {
		fmt.Println("FindUserRoleByUserId>One", err)
		return nil, err
	}
	_, err = o.LoadRelated(r, "Users")
	if err != nil {
		fmt.Println("FindUserRoleByUserId>LoadRelated", err)
		return nil, err
	}
	return r, nil
}
