package database

import (
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func (c *Connection) Close() error {
	if c.DB == nil {
		return nil
	}
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
