package ds

type Kino struct {
	ID      uint `gorm:"primarykey"`
	Code    uint
	Name    string
	Release int
	Grade   int
}
