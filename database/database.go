package database

import (
	"gorm.io/gorm"
)

func DBInit() (*gorm.DB, error) {
    //db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    //if err != nil {
    //    return nil, err
    //}

    //err = AutoMigrate(db)
    //if err != nil {
    //    return nil, err
    //}

    //return db, nil
    return nil, nil
}