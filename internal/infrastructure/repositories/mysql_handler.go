package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	CannotConnectToDatabaseError = errors.New("Cannot connect to database")
)

func ThrowCannotConnectToDatabaseError(actualErr error) error {
	return errors.New(fmt.Sprintf("%s ==> %s", CannotConnectToDatabaseError.Error(), actualErr.Error()))
}

func NewMySQLDB(conString string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true, // sonraki sorgular için cache
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, ThrowCannotConnectToDatabaseError(err)
	}

	return db, nil
}
