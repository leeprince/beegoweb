package models

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

// 用户基本信息
type User struct {
	Id   int
	Name string
	
	/********设置一对一*/
	Profile *Profile `orm:"rel(one)"`
	
	/********** 设置一对多的反向关系
	Post 结构体设置一对多关系: User  *User `orm:"rel(fk)"` 之后，User 模型的 Post 可以不设置
	*/
	Post []*Post `orm:"reverse(many)"`
}

// 用户信息
type Profile struct {
	Id  int
	Age int16
	
	/******* 设置一对一反向关系(可选)
	// User 结构体设置一对一关系: Profile *Profile `orm:"rel(one)"`之后，下面存在时测试用户未打印数据
	*/
	// User        *User   `orm:"reverse(one)"`
}

// // 文章
type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` // 设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

// 标签
type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` // 设置多对多反向关系
}

func init() {
	orm.RegisterModel(
		new(User),
		new(Profile),
		new(Post),
		new(Tag),
	)
}

/*
User 和 Profile 是 OneToOne 的关系
*/
func UserProfile(uid int) {
	o := orm.NewOrm()
	
	/*********************** read(): 复杂的单个对象查询参见 One
	1. SELECT `id`, `name`, `profile_id` FROM `user` WHERE `id` = ?
	2. SELECT `id`, `age` FROM `profile` WHERE `id` = ?
	*/
	/*user := &User{Id: uid}
	o.Read(user)
	if user.Profile != nil {
		o.Read(user.Profile)
	}*/
	
	user := &User{}
	/*********************** RelatedSel(): 设置关系模型一起查询
	SELECT
		T0.`id`,
		T0.`name`,
		T0.`profile_id`,
		T1.`id`,
		T1.`age`
	FROM
		`user` T0
		INNER JOIN `profile` T1 ON T1.`id` = T0.`profile_id`
	WHERE
		T0.`id` = ?
		LIMIT 1
	
	o.QueryTable(user).Filter("id", uid).RelatedSel("Profile").One(user) == o.QueryTable(user).Filter("id", uid).RelatedSel("Profile").One(user)
	*/
	o.QueryTable(user).Filter("id", uid).RelatedSel().One(user)
	// o.QueryTable(user).Filter("id", uid).RelatedSel("Profile").One(user)
	
	/************************ LoadRelated(md, ...) 将相关模型加载到 md 模型
	1. SELECT T0.`id`, T0.`name`, T0.`profile_id` FROM `user` T0 WHERE T0.`id` = ? LIMIT 1;
	
	2. SELECT
		T0.`id`,
		T0.`age`
	FROM
		`profile` T0
		INNER JOIN `user` T1 ON T1.`profile_id` = T0.`id`
	WHERE
		T1.`id` = ?
		LIMIT 1
	*/
	/*o.QueryTable(user).Filter("id", uid).One(user)
	o.LoadRelated(user, "Profile")*/
	
	b, _ := json.MarshalIndent(user, "", "		")
	fmt.Println(string(b))
}

/*
一对多
*/
func UserPost(uid int) {
	o := orm.NewOrm()
	
	var posts []*Post
	num, err := o.QueryTable("post").Filter("User", 1).RelatedSel().All(&posts)
	
	/**** One 只能查到一个用户一个文章（post）的信息, 而不是一个用户对应多个文章的信息。不符合关系
	1. SELECT T0.`id`, T0.`title`, T0.`user_id` FROM `post` T0 INNER JOIN `user` T1 ON T1.`id` = T0.`user_id` WHERE T0.`user_id` = ? LIMIT 1
	2. SELECT T0.`id`, T0.`name`, T0.`profile_id` FROM `user` T0 INNER JOIN `post` T1 ON T1.`user_id` = T0.`id` WHERE T1.`id` = ? LIMIT 1
	*/
	/*var posts = Post{}
	err := o.QueryTable("post").Filter("User", 1).One(&posts)
	o.LoadRelated(&posts, "User")*/
	
	if err == nil {
		fmt.Println("UserPost: num, posts", num, posts)
		
		b, _ := json.MarshalIndent(posts, "", "		")
		fmt.Println(string(b))
		
		for _, post := range posts {
			fmt.Printf(">>>>>>>>> Id: %d, Name: %s, Title: %s\n", post.Id, post.User.Name, post.Title)
		}
	}
}

/**
多对多
*/
func PostTags(name string) {
	dORM := orm.NewOrm()
	
	/** 通过标签查找文章标题
	 */
	/*var posts []*Post
	dORM.QueryTable("post").Filter("Tags__Tag__Name", name).All(&posts)
	j, _ := json.MarshalIndent(posts, "", "		")
	fmt.Println(string(j))*/
	
	/*** 通过文章标题查找标签
	 */
	var tags []*Tag
	dORM.QueryTable("tag").Filter("Posts__Post__Title", name).All(&tags)
	b, _ := json.MarshalIndent(tags, "", "		")
	fmt.Println(string(b))
}
