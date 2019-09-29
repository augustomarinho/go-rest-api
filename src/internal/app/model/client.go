package model

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) *Client {
	return &Client{
		Email: email,
		Name:  name,
	}
}
