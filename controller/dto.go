package controller

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Class

type CreateClassRequest struct {
	Name     string   `json:"name" binding:"required"`
	Teacher  string   `json:"teacher" binding:"required"`
	Students []string `json:"students"`
}

type UpdateClassRequest struct {
	Name     string   `json:"name,omitempty"`
	Teacher  string   `json:"teacher,omitempty"`
	Students []string `json:"students,omitempty"`
}
