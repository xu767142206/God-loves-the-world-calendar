package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var Db *gorm.DB

func InitDb() {
	var err error
	Db, err = gorm.Open("sqlite3", "../shengjing.db")
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to connect database"))
	}
	log.Println("链接chenggong")
}
