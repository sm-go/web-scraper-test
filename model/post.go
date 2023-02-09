package model

type Post struct {
	Id       int64  `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"title" json:"title"`
	SubTitle string `gorm:"sub_title" json:"subtitle"`
	Category string `gorm:"category" json:"category"`
	Image    string `gorm:"image" json:"image"`
}

func (m *Post) TableName() string {
	return "posts"
}
