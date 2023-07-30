package test

import (
	"fmt"
	"gin-gorm-oj/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

// A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

func TestGormTest(t *testing.T) {
	dsn := "root:903290aA+@tcp(192.168.91.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init error: ", err)
	}
	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problem ==> %v\n", v)
	}
}
