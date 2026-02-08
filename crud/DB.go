package main

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {

	dsn := "host=localhost user=ellenchung password=postgre dbname=first_db port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

}

func create(up *User) error {

    return db.Create(up).Error


}

// 注意，err != nil的时候返回值不可信！
func read(userID uuid.UUID) (*User, error) {
	up := &User{}

	err := db.Model(&User{}).
		Where("user_id = ?", userID).
		First(up).Error

	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // 找不到不是 DB 错误
	}
	return up, err

}

func update(up *User) (int, error) {
	res := db.Save(up)

	return int(res.RowsAffected), res.Error
}

func delete(userID uuid.UUID) (int, error) {
	res := db.Where("user_id = ?", userID).Delete(&User{})

	return int(res.RowsAffected), res.Error
}
