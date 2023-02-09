package model

type PostDetail struct {
	Id        int64  `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Category  string `gorm:"category" json:"category"`
	Image     string `gorm:"image" json:"image"`
	Paragraph string `gorm:"paragraph" json:"paragraph"`
	// Date      string `gorm:"date" json:"date"`
}

func (m *PostDetail) TableName() string {
	return "post_details"
}
