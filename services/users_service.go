package services

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)


// GetUser ユーザーの取得
func GetUser(userID int64)(*users.User, *errors.RestErr){
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil{
		return nil, err
	}
	return result, nil
}



// CreateUser ユーザーの新規追加
func CreateUser(user users.User)(*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil{
		return nil, err
	}
	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil
}

