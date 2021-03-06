package models

import "time"

type User struct {
	Id   int     `json:"id" xorm:"id"`
	UserName string  `json:"name" xorm:"name"`
	Password  string  `json:"pwd" xorm:"pwd"`
}

type UpdateUserDto struct {
	NewPassword  string  `form:"new_password" json:"new_password"`
	ReNewPassword  string  `form:"re_new_password" json:"re_new_password"`
	DisplayName  string  `form:"display_name" json:"display_name"`
	Email  string  `form:"email" json:"email"`
	Phone  string  `form:"phone" json:"phone"`
}

type Session struct {
	User  string  `json:"user"`
	Token  string  `json:"token"`
	Expire  int64  `json:"expire"`
}

type UserTable struct {
	Id  int  `json:"id"`
	Name  string  `json:"name"`
	Passwd  string  `json:"passwd"`
	DisplayName  string  `json:"display_name"`
	Email  string  `json:"email"`
	Phone  string  `json:"phone"`
	ExtContactOne  string  `json:"ext_contact_one"`
	ExtContactTwo  string  `json:"ext_contact_two"`
	Creator  string  `json:"creator"`
	Created  time.Time  `json:"created"`
}

type RoleTable struct {
	Id  int  `json:"id"`
	Name  string  `json:"name"`
	DisplayName  string  `json:"display_name"`
	Creator  string  `json:"creator"`
	Created  time.Time  `json:"created"`
}

type SendAlertObj struct {
	Accept  []string
	Subject  string
	Content  string
}

type RelRoleUserTable struct {
	Id  int  `json:"id"`
	RoleId  int  `json:"role_id"`
	UserId  int  `json:"user_id"`
}