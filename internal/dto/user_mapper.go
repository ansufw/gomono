// internal/dto/user_mapper.go
package dto

import (
	"encoding/json"

	"github.com/ansufw/gomono/internal/infrastructure/database/sqlc"
)

func MapUserToResponse(dbUser sqlc.GetUserRow) UserResponse {
	// Parse the roles JSON from interface{}
	var roles []Role
	if dbUser.Roles != nil {
		rolesBytes, err := json.Marshal(dbUser.Roles)
		if err == nil {
			json.Unmarshal(rolesBytes, &roles)
		}
	}

	return UserResponse{
		ID:        dbUser.ID.String(),
		FirstName: dbUser.FirstName.String,
		LastName:  dbUser.LastName.String,
		Email:     dbUser.Email,
		Roles:     roles,
		CreatedAt: dbUser.CreatedAt.Time,
	}
}

func MapUsersToResponse(dbUsers []sqlc.ListUsersRow) UserListResponse {
	users := make([]UserResponse, len(dbUsers))
	for i, user := range dbUsers {
		// Parse the roles JSON from interface{}
		var roles []Role
		if user.Roles != nil {
			rolesBytes, err := json.Marshal(user.Roles)
			if err == nil {
				json.Unmarshal(rolesBytes, &roles)
			}
		}

		users[i] = UserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
			Email:     user.Email,
			Roles:     roles,
			CreatedAt: user.CreatedAt.Time,
		}
	}
	return UserListResponse{
		Users: users,
		Total: len(dbUsers),
	}
}
