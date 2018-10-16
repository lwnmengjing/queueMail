package form

type MessageMail struct {
	Host        string   `json:"host" binding:"required"`
	Port        int      `json:"port" binding:"required"`
	Username    string   `json:"username" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	Attachments []string `json:"attachments" binding:"required"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body" binding:"required"`
	From        From     `json:"from" binding:"required"`
	To          string   `json:"to" binding:"required"`
	Auth        bool     `json:"auth" binding:"required"`
}

type From struct {
	Email string `json:"email" binding:"email"`
	Name  string `json:"name" binding:"required"`
}
