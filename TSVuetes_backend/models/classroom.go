package models

type Classroom struct {
	Room_id   string `gorm:"primaryKey;not null" json:"room_id"` // 主键或唯一键
	Device_id string `json:"device_id"`
}
