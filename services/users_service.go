package services

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)


// GetUser ユーザー情報の取得
func GetUser(userID int64)(*users.User, *errors.RestErr){
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil{
		return nil, err
	}
	return result, nil
}

// CreateUser ユーザー情報の新規追加
func CreateUser(user users.User)(*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil{
		return nil, err
	}
	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil
}

// UpdateUser ユーザー情報の変更
func UpdateUser(isPartial bool, user users.User)(*users.User, *errors.RestErr){
	current, err := GetUser(user.ID)
	if err != nil{
		return nil, err
	}

	if isPartial {
		if user.FirstName != ""{
			current.FirstName = user.FirstName
		}
		if user.LastName != ""{
			current.LastName = user.LastName
		}
		if user.NickName != ""{
			current.NickName = user.NickName
		}
		if user.Email != ""{
			current.Email = user.Email
		}
		if user.Age != 0{
			current.Age = user.Age
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.NickName = user.NickName
		current.Email = user.Email
		current.Age = user.Age
	}

	if err := current.Update(); err != nil{
		return nil, err
	}
	return current, nil
}

// DeleteUser ユーザー情報の削除
func DeleteUser(userID int64) *errors.RestErr{
	user := &users.User{ID: userID}
	return user.Delete()
}