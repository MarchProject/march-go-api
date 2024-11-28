// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

type LoginInputParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Price struct {
	Price float64 `json:"price"`
}

type Query struct {
}

type ResponseCreateUser struct {
	Data   *User   `json:"data,omitempty"`
	Status *Status `json:"status,omitempty"`
}

type ResponseGetUser struct {
	Data   []*User `json:"data,omitempty"`
	Status *Status `json:"status,omitempty"`
}

type ResponseLogin struct {
	Data   *Token  `json:"data,omitempty"`
	Status *Status `json:"status,omitempty"`
}

type SignOutResponse struct {
	ID string `json:"id"`
}

type Status struct {
	Message *string `json:"message,omitempty"`
	Code    int     `json:"code"`
}

type Token struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Username     *string `json:"username,omitempty"`
	UserID       *string `json:"userId,omitempty"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserInputParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type VerifyAccessTokenResponse struct {
	Success *bool `json:"success,omitempty"`
}
