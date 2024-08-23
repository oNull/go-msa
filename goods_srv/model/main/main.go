package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func genMd5Sale(code string) string {
	// 加密 https://github.com/anaskhan96/go-password-encoder
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(code, options)

	// 拼接密码到数据库
	newPassword := fmt.Sprintf("$pbkbf2-sha512$%s$%s", salt, encodedPwd)

	// 验证密码
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify(code, passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
	return encodedPwd
}

// 链接数据库 同步数据表
func main() {
	//log.Println("Database connected")
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,        // Don't include params in the SQL log
	//		Colorful:                  false,       // Disable color
	//	},
	//)
	//dsm := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&charset=utf8mb4,utf8&parseTime=true", "root", "888888", "127.0.0.1", "3306", "mxshop")
	//
	//Db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   "mx_",
	//		SingularTable: true,
	//	},
	//	Logger:      newLogger,
	//	PrepareStmt: true, // 启用预编译语句以提高性能
	//})
	//
	//if err != nil {
	//	log.Println("err:", err.Error())
	//}
	//
	//// 迁移 schema
	//_ = Db.AutoMigrate(&model.User{})

}
