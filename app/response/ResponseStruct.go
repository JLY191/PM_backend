package response

type User struct {
	Name     string `json:"name"   form:"name"    query:"name" `
	Email    string `json:"email"  form:"email"   query:"email" `
	Password string `json:"password"    form:"password"     query:"password" `
}
