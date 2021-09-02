package pachong

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Db数据库连接池
var DB *gorm.DB

//注意方法名大写，就是public
func InitDB()  {
	dsn := "host=xxx user=name password=password dbname=demo port=3432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err1 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err1 != nil {
		panic("数据库连接失败！！！！")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("数据库初始化失败！！！！")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	DB = db
	fmt.Println("数据库链接成功...........")
}
