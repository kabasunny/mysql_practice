package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ユーザーデータを格納するテーブル用の構造体を定義
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"uniqueIndex;size:100"`
	Password string `gorm:"size:255"`
}

func main() {
	// MySQLへの接続情報
	dsn := "zeninvestor:zenpass@tcp(localhost:3306)/zeninv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// テーブルの自動マイグレーション
	db.AutoMigrate(&User{})

	// ユーザーデータを作成
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Password: "password1"},
		{Name: "Bob", Email: "bob@example.com", Password: "password2"},
		{Name: "Charlie", Email: "charlie@example.com", Password: "password3"},
	}

	// ユーザーデータをデータベースに追加
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatalf("Failed to insert user: %v", result.Error)
		}
	}

	log.Println("Database connected and User table migrated successfully. Users added.")
}
