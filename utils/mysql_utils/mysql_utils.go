package mysql_utils

import (
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result"
)
func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NewNotFoundErr("no record matching given id")
		}
		return errors.NewInternalServerErr("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestErr("invalid data")
	}
	return errors.NewInternalServerErr("error processing request")
}