package adduser

type AddUserRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}
