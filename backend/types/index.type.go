package types

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ZitadelResponse struct {
    Details []UserDetail `json:"result"`
    Total   int         `json:"totalResult"`
}

type UserDetail struct {
    ID          string `json:"id"`
    UserName    string `json:"userName"`
    Email       string `json:"email"`
    FirstName   string `json:"firstName"`
    LastName    string `json:"lastName"`
    DisplayName string `json:"displayName"`
}

type UserProfile struct {
    ID          string   `json:"id"`
    Username    string   `json:"username"`
    Email       string   `json:"email"`
    FirstName   string   `json:"firstName"`
    LastName    string   `json:"lastName"`
    DisplayName string   `json:"displayName"`
    Roles       []string `json:"roles"`
}

type APISucessResponse[T any] struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Results    T `json:"results,omitempty"`
}

type IntrospectionResponse struct {
    Active    bool     `json:"active"`
    Scope     string   `json:"scope"`
    ClientID  string   `json:"client_id"`
    Username  string   `json:"username"`
    TokenType string   `json:"token_type"`
    Exp       int64    `json:"exp"`
    Iat       int64    `json:"iat"`
    Nbf       int64    `json:"nbf"`
    Sub       string   `json:"sub"`
    Aud       []string `json:"aud"`
    Iss       string   `json:"iss"`
    Jti       string   `json:"jti"`
}