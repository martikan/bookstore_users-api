package user

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

// Function for Marshall request/response for User struct.
// If `public` is true, than it will use the public user 'schema'
// If `public` is false, it will use the private user 'schema'
func (u *User) Marshall(public bool) interface{} {

	if public {
		// We have to explicitly declare it because the ids in the jsons are different
		return PublicUser{
			Id:        u.Id,
			CreatedAt: u.CreatedAt,
			Status:    u.Status,
		}
	}

	userJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)

	return privateUser
}

func (usrs Users) Marshall(public bool) []interface{} {
	resultset := make([]interface{}, len(usrs))
	for i, u := range usrs {
		resultset[i] = u.Marshall(public)
	}

	return resultset
}
