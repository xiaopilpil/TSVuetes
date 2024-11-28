package models

type Teacher struct {
	Teacher_id string `gorm:"primaryKey;not null" json:"teacher_id"` // 主键或唯一键
	Password   string `json:"password"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Faculty    string `json:"faculty"`
	Position   int    `json:"position"` // 0是普通用户，1是管理员
}
