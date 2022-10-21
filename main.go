// packageというものをmainという名前で定義している。
package main

//インストールしたginとdbパッケージを使えるようにしている。
import (
	"go-quest/db"
	"go-quest/routers"

	"github.com/gin-gonic/gin"
)

// main関数を定義
func main() {
	//ginを初期化してrという変数に代入
	//goでは、:=とすることで暗黙的に型宣言を行ってくれます。
	r := gin.Default()

	//さっき作ったInitDBを実行
	db.InitDB()

	routers.InitRouter(r)

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
