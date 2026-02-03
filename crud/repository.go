package main

import (
	"errors"
	"log"


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

	
	return up, nil

}

func update(userID string, infp *UserInfo) (int, error) {
	res := db.Model(&User{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"username": infp.Username,
			"password": infp.Password,
		})

	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}

func delete(userID string) (int, error) {
	res := db.Where("user_id = ?", userID).Delete(&User{})
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
