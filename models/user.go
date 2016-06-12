package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id      int
	Name    string `form:"name" valid:"Required;MaxSize(20);MinSize(6)"`
	Pass    string `form:"pass" valid:"Required;MaxSize(20);MinSize(6)"`
	Email   string `form:"email" valid:"Required;Email"`
	Phone   string `form:"phone" valid:"Required;Mobile"`
	Image   string `form:"image" valid:"MaxSize(50);MinSize(6)"`
	Addr    string `form:"addr" valid:"MaxSize(30)" form:"name"`
	Regtime string
	Birth   string `form:"birth"`
	Remark  string `valid:"MaxSize(200)" form:"remark"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "t_user"
}

func (this *User) ToString() string {
	return fmt.Sprintf("Id:%d\tName:%s", this.Id, this.Name)
}

func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func GetUserList(page int64, pageSize int64, user *User) (userlist []*User, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	if user.Id != 0 {
		qs = qs.Filter("id__exact", user.Id)
	}
	if user.Phone != "" {
		qs = qs.Filter("phone__contains", user.Phone)
	}
	if user.Name != "" {
		qs = qs.Filter("name__contains", user.Name)
	}
	qs.Limit(pageSize, offset).OrderBy("-Id").All(&userlist)
	count, _ = qs.Count()
	return userlist, count
}

func DelUser(idstr string) error {
	ids := strings.Split(idstr, ",")
	num, err := orm.NewOrm().QueryTable("t_user").Filter("id__in", ids).Delete()
	if num > 0 && err == nil {
		return nil
	} else {
		return err
	}
}

func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Name = u.Name
	user.Email = u.Email
	user.Phone = u.Phone
	user.Image = u.Image
	user.Addr = u.Addr

	tm2, _ := time.Parse("2006-02-02", u.Birth) //日期字符串转为时间戳
	user.Birth = strconv.FormatInt(tm2.Unix(), 10)

	user.Remark = u.Remark
	user.Regtime = strconv.FormatInt(time.Now().Unix(), 10) //获取当前时间戳
	id, err := o.Insert(user)
	return id, err
}

func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Name) > 0 {
		user["Name"] = u.Name
	}
	if len(u.Phone) > 0 {
		user["Phone"] = u.Phone
	}
	if len(u.Addr) > 0 {
		user["Addr"] = u.Addr
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Birth) > 0 {
		tm, _ := time.Parse("2006-02-02", u.Birth) //日期字符串转为时间戳
		user["Birth"] = strconv.FormatInt(tm.Unix(), 10)
	}
	if len(u.Image) > 0 {
		user["Image"] = u.Image
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Name: username}
	o := orm.NewOrm()
	o.Read(&user, "Name")
	return user
}
