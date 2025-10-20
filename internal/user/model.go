package user

type User struct {
	Id       uint32 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Tasks    []taskService.Task `json:"tasks"`
}
