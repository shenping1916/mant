package log

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"testing"
	"time"
)

type User struct {
	ID           int64      `gorm:"primary_key; -" comment:"用户表(user),id" json:",omitempty"`
	JobNumber    string     `gorm:"column:job_number; type:varchar(10);unique_index:idx_job_number_name; not null" comment:"用户表(user),工号" json:"jobnumber"`
	Name         string     `gorm:"column:name; type:varchar(50);unique_index:idx_job_number_name; not null" comment:"用户表(user)," json:"realName"`
	Email        string     `gorm:"column:email; type:varchar(100);unique_index:idx_email; not null" json:"email"`
	DepartmentL1 string     `gorm:"column:department_level1; type:varchar(100);not null" comment:"用户表(user),一级部门" json:"department_level1"`
	DepartmentL2 string     `gorm:"column:department_level2; type:varchar(100);not null" comment:"用户表(user),二级部门" json:"department_level2"`
	DepartmentL3 string     `gorm:"column:department_level3; type:varchar(100);not null" comment:"用户表(user),三级部门" json:"department_level3"`
	Position     string     `gorm:"column:position; type:varchar(100); not null" comment:"用户表(user),职位" json:"position"`
	Token        string     `gorm:"column:token; type:varchar(20)" comment:"用户表(user),gitlab访问令牌" json:"token"`
	Org          [][]string `gorm:"-" json:",omitempty"`
	Description  string     `gorm:"column:description; type:text; default: null" comment:"用户表(user),描述" json:"description"`
	Time
	DeletedAt *time.Time `gorm:"column:delete_time; type: timestamp without time zone" comment:"删除时间" json:"delete_time,omitempty"`
}

type Time struct {
	CreatedAt time.Time `gorm:"column:create_time; type: timestamp without time zone" comment:"创建时间" json:"create_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:update_time; type: timestamp without time zone" comment:"更新时间" json:"update_time,omitempty"`
}

func (User) TableName() string {
	return "t_jcc_user"
}

func TestLogFormatGorm(t *testing.T) {
	var (
		host = "127.0.0.1"
		port = 3625
		username = "appcooltesttissot"
		passwd = "appcooltesttissot"
		dbname = "appcooltesttissot"
		sslmode = "disbale"
		connect_time = 5
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s connect_timeout=%d",host, port, username, passwd, dbname, sslmode, connect_time)
	db, _ := gorm.Open("postgres", dsn)
	defer db.Close()

	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetOutput(CONSOLE, nil)
	logger.SetAsynChronous()
	logger.SetColour()

	db.LogMode(true)
	db.SetLogger(logger)

	var user User
	err := db.Model(&user).Find(&user).Error
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(user)
}