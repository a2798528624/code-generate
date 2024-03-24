package models

type Authors struct {
	Id      int    `gorm:"id"`
	NAME    string `gorm:"NAME"`
	Country string `gorm:"country"`
	Age     int    `gorm:"age"`
}

func (*Authors) TableName() string {
	return "authors"
}
