// db package
package db

// Todoモデルを定義
// PrimaryKeyはID、Contentはtodoの内容
type Todo struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
}
