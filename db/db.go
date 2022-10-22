package db

//gormとdriverのインポート
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Ssql *gorm.DB
)

// データベースを初期化する関数の定義
func InitDB() (err error) {
	//ここでsqliteを開く
	Ssql, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//errが返ってきた場合はpanicを起こすようにする
	if err != nil {
		return
	}
	//さっき作ったTodoモデルでテーブルを作成する
	if err = Ssql.AutoMigrate(&Todo{}); err != nil {
		return
	}

	return
}
