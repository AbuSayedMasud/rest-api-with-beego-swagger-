package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id    int    `orm:"column(Id);auto"`
	Name  string `orm:"column(Name);null"`
	Age   int    `orm:"column(Age);null"`
	Email string `orm:"column(Email);null"`
	Image string `orm:"column(Image);null"`
}
type PatchOpp struct{
	Path string 
	Value string
}

func (t *Users) TableName() string {
	return "Users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	fmt.Println("gret")
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users. Returns empty list if
// no records exist
func GetAllUsers() (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	var l []Users
	if _, err = qs.All(&l); err == nil {

				for _, v := range l {
					ml = append(ml, v)
				}

		return ml, nil
	}
	return nil, err
}

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
func PatchUsersById(q *PatchOpp,id int)(error, *Users){
	o := orm.NewOrm()
	v := Users{Id: id}
    var err error
	if err = o.Read(&v); err == nil {
		
		l:=q.Path
		l=strings.Title(l)
		if(l=="Email"){
			v.Email=q.Value
		}else if(l=="Name"){
		 v.Name = q.Value
		}else if(l=="Age"){
			age, _ := strconv.Atoi(q.Value)
			v.Age = age
		}else if(l=="Image"){
			v.Image=q.Value
		}else{
			return err,&v
		}
		var num int64
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return  err,&v
}
// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
