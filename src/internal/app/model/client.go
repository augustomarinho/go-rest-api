package model

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) *Client {
	client := new(Client)
	client.Email = email
	client.Name = name

	return client
}
