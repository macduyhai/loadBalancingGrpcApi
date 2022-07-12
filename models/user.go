package models

import "time"

type User struct {
	ID         int64      `gorm:column:id;PRIMERY_KEY`
	Username   string     `gorm:column:username`
	Password   string     `gorm:column:password`
	Fullname   string     `gorm:column:fullname`
	Salary     int64      `gorm:column:salary`
	Permission int8       `gorm:column:permission`
	Delstatus  int8       `gorm:column:delstatus`
	Active     int8       `gorm:column:active`
	CreateTime *time.Time `gorm:column:create_time`
	UpdateTime *time.Time `gorm:column:update_time`
}
