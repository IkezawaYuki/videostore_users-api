package users

import "encoding/json"

type PublicAdminUser struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
}

type PrivateAdminUser struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NickName    string `json:"nick_name"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status 		string `json:"status"`
}

func (adminUsers AdminUsers) Marshall(isPublic bool) []interface{}{
	result := make([]interface{}, len(adminUsers))
	for index, adminUser := range adminUsers{
		result[index] = adminUser.Marshall(isPublic)
	}
	return result
}

func (adminUser *AdminUser) Marshall(isPublic bool) interface{}{
	if isPublic{
		return &PublicAdminUser{
			ID:          adminUser.ID,
			UserID:      adminUser.UserID,
			DateCreated: adminUser.DateCreated,
			Status:      adminUser.Status,
		}
	}
	userJson, _ := json.Marshal(adminUser)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}