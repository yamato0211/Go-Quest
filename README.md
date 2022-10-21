# Go Quest
goのフレームワークのginを使って簡単なtodoアプリケーションを構築しよう。
データベースには、ファイル形式で簡単なsqliteを使用します。
分からないことは適宜調べるなり、質問するなりしよう。

## Lv.1 Goインストール
goをインストールしよう。
以下を参考にしてみてください。
[wslにgoをインストールする](https://syoblog.com/golang-ubuntu/)


## Lv.2 Ginのインストール
さっそくginをインストールしていこう！
まず適当にプロジェクトディレクトリを作成しましょう。
そして、ターミナルで
```
$ go mod init go-quest
$ go get -u github.com/gin-gonic/gin
```
上の2コマンドを実行しましょう。
ここでエラーが出る場合はgoのインストールがうまくいっていない可能性があります。
エラーを調べてみて分からなかったら、ぜひ質問をしてください！
ディレクトリの中を見てみると、go.sumとgo.modができていると思います。
これはgoのパッケージを管理するためのファイルです。


## Lv.3 Ginを書いていこう
インストールが確認できたところで、さっそく書いていきましょう。
vscodeなどのエディターを開いて、go.modやgo.sumと同じ階層に
main.goを作成しよう。

main.goを以下のように編集しよう。
```main.go
// packageというものをmainという名前で定義している。
package main

//インストールしたginを使えるようにしている。
import "github.com/gin-gonic/gin"

// main関数を定義
func main() {
	//ginを初期化してrという変数に代入
	//goでは、:=とすることで暗黙的に型宣言を行ってくれます。
	r := gin.Default()

	//GETメソッドを定義
	//"/"にアクセスしたときに、
	//{"message": "Hello World!!"}を返すよ！
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})

	//ginを実行するよ
	//()内に下のような書き方をするとポートの指定ができるよ！
	r.Run(":8000")
}
```

## Lv.4 Ginを実行しよう
プロジェクトディレクトリで以下のコマンドを実行しよう！
```
$ go run main.go 
```
そうしてブラウザのurlの欄に
**localhost:8000**としてアクセスしてみよう。
**{"message" : "Hello World!!"}**が表示されていたら成功！！
これで今日からginマスター(やったね)

## Lv.5 データベースについて
データベースは簡単に言うと、データを保存する場所です。
フロントのみのアプリケーションだとデータが保存されないので、
リロードすると消えてしまいます。なので、データベースに情報を保存して、
必要なときに引っ張ってこれるようにしたいです。

データベースを操作するにはSQLというクエリ言語を書かないといけないのですが、
ORM(Object Relation Mapper)を使用することで、
その言語likeにクエリを投げることができます。すごい！！
今回はGormというormを使っていきます。
[GORMガイド](https://gorm.io/ja_JP/docs/index.html)

データベースはSqliteを使用します。
これはファイル形式のデータベースで管理が簡単なので使用します。
Postgresqlなどを使用したい場合はサーバを立てる必要があります。
Dockerなどを使用して、サーバを立てたりしますが、これはまた今度。

## Lv.6 goでSqliteに接続しよう
まずgormとsqliteDriverをインストールする必要があるので、
```
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/sqlite
```
プロジェクトディレクトリにdbディレクトリを作成しよう。
dbディレクトリに中に**db.go**と**model.go**を作成します。

model.goにはデータベースに使用するモデルを構造体で定義します。
db.goにはデータベースに接続するためのあれこれを書いていきます。

```model.go
// db package
package db

// Todoモデルを定義
// PrimaryKeyはID、Contentはtodoの内容
type Todo struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
}
```

```db.go
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
```

InitDB()ができたらmain.goを修正します。
```main.go
// packageというものをmainという名前で定義している。
package main

//インストールしたginとdbパッケージを使えるようにしている。
import (
	"go-quest/db"

	"github.com/gin-gonic/gin"
)

// main関数を定義
func main() {
	//ginを初期化してrという変数に代入
	//goでは、:= とすることで暗黙的に型宣言を行ってくれます。
	r := gin.Default()

	//さっき作ったInitDBを実行
	db.InitDB()
	//GETメソッドを定義
	//"/"にアクセスしたときに、
	//{"message": "Hello World!!"}を返すよ！
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})

	//ginを実行するよ
	//()内に下のような書き方をするとポートの指定ができるよ！
	r.Run(":8000")
}
```
そして、
```
go run main.go
```
を実行してみよう。プロジェクトディレクトリ直下に
test.dbファイルが作成されていると成功！！
Foo!!これでデータベースの接続に完了したよ！

## Lv.7 データを挿入する処理を書こう
データエースに接続できたので、データを投入できるようにしていこう。
プロジェクトディレクトリ直下に**crudsディレクトリ**を作成しよう。
