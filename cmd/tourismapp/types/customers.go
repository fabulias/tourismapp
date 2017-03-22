package 

`form:"user" json:"user" binding:"required"`
type Customer struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	S_surname string `json:"s_surname"`
	Rut       string `json:"rut"`
	Mail      string `json:"mail"`
	Password  string `json:"pass"`
	Status    bool   `json:"status"`
}
