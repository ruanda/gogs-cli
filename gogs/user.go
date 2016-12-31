package gogs

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func (c *Client) UserShow(username string) (*User, error) {
	data, err := c.getResponse("GET", fmt.Sprintf("/users/%s", username), nil)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
	var res genericResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	var users []User
	if res.Ok {
		err = json.Unmarshal(res.Data, &users)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
