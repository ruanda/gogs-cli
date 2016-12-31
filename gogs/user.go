package gogs

import (
	"encoding/json"
	"strconv"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

func (c *Client) UserSearch(q string, limit int) ([]User, error) {
	query := map[string]string{
		"q":     q,
		"limit": strconv.Itoa(limit),
	}

	data, err := c.getResponse("GET", "/users/search", query)
	if err != nil {
		return nil, err
	}

	var users []User

	if data.Ok {
		err = json.Unmarshal(data.Data, &users)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}
