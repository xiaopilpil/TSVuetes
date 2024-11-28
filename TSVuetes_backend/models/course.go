package models

type Course struct {
	Teacher_id string    `gorm:"primaryKey;not null" json:"teacher_id"`
	Room_id    string    `gorm:"primaryKey;not null" json:"room_id"`
	Teacher    Teacher   `gorm:"foreignKey:Teacher_id;references:Teacher_id"`
	Classroom  Classroom `gorm:"foreignKey:Room_id;references:Room_id"`
	Name       string    `json:"name"`
	Faculty    string    `json:"faculty"`
	Snumber    int       `json:"snumber"`
	State      int       `json:"state"` //0是申请中，1是已通过，2是已拒绝
	Agree_id   string    `json:"agree_id"`
	Ps         string    `json:"ps"`
}
