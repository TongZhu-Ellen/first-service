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

func create(cp *UserCreation) (*User, error) {

	up := &User{
		UserID:   uuid.NewString(),
		Username: cp.Username,
		Password: cp.Password,
	}

	err := db.Create(up).Error
	if err != nil {
		return nil, err
	}

	up.Password = "******" // 6 *'s
	return up, nil

}

func read(userID string) (*User, error) {
	up := &User{}

	err := db.Model(&User{}).
		Where("user_id = ?", userID).
		First(up).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	up.Password = "******"
	return up, nil

}

func update(userID string, upp *UserUpdating) (int, error) {
	res := db.Model(&User{}).
		Where("user_id = ? AND password = ?", userID, upp.Password).
		Updates(map[string]interface{}{
			"username": upp.NewUsername,
			"password": upp.NewPassword,
		})

	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}

func delete(userID string, dp *UserDeletion) (int, error) {
	res := db.Where("user_id = ? AND password = ?", userID, dp.Password).Delete(&User{})
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
