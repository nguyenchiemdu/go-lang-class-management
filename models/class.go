package models

type Class struct {
	ID         string     `json:"_key,omitempty"`
	Name       string     `json:"name" binding:"required"`
	TeacherID  string     `json:"teacher_id,omitempty"`
	StudentIDs []string   `json:"student_ids,omitempty"`
	Teacher    *Teacher   `json:"teacher,omitempty"`
	Students   []*Student `json:"students,omitempty"`
}
