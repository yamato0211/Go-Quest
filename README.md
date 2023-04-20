# Go Quest
![image](https://user-images.githubusercontent.com/88587703/197323802-e4bf05c4-2e20-4819-869c-08e430a03f61.png)

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

{"message" : "Hello World!!"}が表示されていたら成功！！

これで今日からginマスター!(やったね)

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

db/model.go
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

db/db.go
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
	//エラーハンドリング
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
package main

//db packageのインポート
import (
	"go-quest/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//さっき作ったInitDBを実行
	db.InitDB()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})
	
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
データべースに接続できたので、データを投入できるようにしていこう。
プロジェクトディレクトリ直下に**crudsディレクトリ**を作成しよう。

その中に**todo.go**を作成しよう。

ここからは自分で考えて書いてもらいたい思います。

Gormは初めてで難しいかもしれませんが、自分で考えて書くという力はとても大事です。

それに、ドキュメントを見て書くという力もつけてほしいので、一応答えは用意していますが、

できる限り自分の力でやってみましょう。質問もどんどんしてください！

[Gormドキュメント](https://gorm.io/ja_JP/docs/)

cruds/todo.go
```todo.go
package cruds

import (
	"go-quest/db"
)

//Todoをデータベースに保存する関数
func CreateTodo(content string) (res_todo db.Todo, err error) {
	//res_todoのContentフィールドに引数のcontentを入れよう
	
	//Gormガイドを見てデータをインサートする処理を書いてみよう
	
	return
}

//すべてのTodoの配列を返す関数
func GetAllTodos() (res_todos []db.Todo, err error) {
	//データベースからすべてのデータを取得する処理を調べて書こう
	
	return
}

//引数のidのTodoを削除する関数
func DeleteTodo(id uint) (err error) {
	//まずデータベースから引数のidと一致するデータを探してこよう
	//そのデータをどうやって削除するかを調べて書こう
	
	return
}
```



<details><summary>回答例</summary><div>


```todo.go
package cruds

import (
	"go-quest/db"
)

func CreateTodo(content string) (res_todo db.Todo, err error) {
	res_todo = db.Todo{Content: content}
	err = db.Ssql.Create(&res_todo).Error
	return
}

func GetAllTodos() (res_todos []db.Todo, err error) {
	err = db.Ssql.Find(&res_todos).Error
	return
}

func DeleteTodo(id uint) (err error) {
	err = db.Ssql.Where("id = ?", id).Delete(&db.Todo{}).Error
	return
}
```

</div></details>


## Lv.8 schemaを定義しよう
このschemaは主にフロントエンドとバックエンド間でのデータやり取りを行うための型です。(modelはデータベース関連)

今回は一つしか必要ないですが、書いてみましょう。

プロジェクトディレクトリ直下に**schema**ディレクトリを作成して、

中に**todo.go**を作成しよう。

schema/todo.go
```todo.go
package schema

//CreateTodoという構造体を定義
//`json:"content"`については後で説明します。
type CreateTodo struct {
	Content string `json:"content"`
}
```

これで、スキーマ作成は完了です。

## Lv.9 ルーティングを追加しよう
APIのurlのことをよく**エンドポイント**と言ったりしますが、

どのエンドポイントで、どういった処理をするかをまだ書いていません。

今回のルーティングは以下のように行います。
#### localhost:8000/todo  Method: GET すべてのtodoを返す
#### localhost:8000/todo  Method: POST todoの新規登録
#### localhost:8000/todo/:id Method: DELETE  パスパラメータのidのTodoを削除する

これを頭に入れておいてください

## Lv.10 CORS対策
今回は、バックエンドが今作っているgoのAPIサーバ

フロントエンドはReactなどで想定をしています。

つまりフロントとバックでデータのやり取りを行うと思うのですが、

なにも対処をせずに通信を行おうとすると、エラーが出てしまします。

これは**CORS**(オリジン間リソース共有)というものでブロックされているからなんです.

詳しくはこれ

[なんとなく CORS がわかる...はもう終わりにする。](https://qiita.com/att55/items/2154a8aad8bf1409db2b)

つまり、このCORSを許可するための記述をしなければならないです。めんどくさい！

プロジェクトディレクトリ直下に、**routersディレクトリ**を作成し、中に**main.go**を作成しよう。

routers/main.go
```main.go
package routers

import (
	//これのインポートでエラーが出るときは
	//go get github.com/gin-contrib/corsをコマンドで実行してください
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Routerを初期化する関数の定義
func InitRouter(r *gin.Engine) {
	//corsミドルウェアを使用する
	//難しいと思うので最初はおまじないで全然OK!
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
	}))
}
```

これでCORSの対策もばっちり

## Lv.11 Todoルーター
cors対策もできたところで、Lv.9で示したルーティングを書いていきましょう。
routersディレクトリ内に**todo.go**を作成しよう。

routers/todo.go
```todo.go
package routers

import (
	"go-quest/cruds"
	"go-quest/db"
	"go-quest/schema"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TodoRouterを初期化する関数
func InitTodoRouters(tr *gin.RouterGroup) {
	//""というパスに対してGETメソッドで、getTodos関数を設定している感じ
	tr.GET("", getTodos)
	//これも同様
	tr.POST("", postTodo)
	// "/:id"とすることで任意の文字列をidが拾ってこれるようになる
	tr.DELETE("/:id", deleteTodo)
}

//データベースからTodoの配列を受け取りそれをフロントに返す関数
func getTodos(c *gin.Context) {
	var (
		todos []db.Todo
		err   error
	)
	
	//crudsで定義したGetAllTodosを使用
	if todos, err = cruds.GetAllTodos(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	//errorがない場合にtodosを返す
	c.JSON(http.StatusOK, &todos)
}

//フロントから送られてきた情報をデータベースに保存する関数
func postTodo(c *gin.Context) {
	var (
		err     error
		payload schema.CreateTodo
		todo    db.Todo
	)
	
	//BindJSONは構造体とjsonを結び付けてくれる。
	//schema.CreateTodoのContentに`json:"content"`
	//と書いていたのは、フロントから送られてくる{"content": "テキストテキスト"}
	//をpayloadのContent要素に勝手に代入してくれる。便利！！
	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	//crudsで定義したCreateTodoを実行
	if todo, err = cruds.CreateTodo(payload.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	//errorがなければ登録したtodoを返す。
	c.JSON(http.StatusOK, &todo)
}

func deleteTodo(c *gin.Context) {
	// c.Param("id")として、パスパラメータを所得できる。
	//でもこれは文字列なのでuint型に変換したい.
	//strconvパッケージを使用して10進数で、64itのuintに変換
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	//crudsで定義したDeleteTodoを実行
	//引数にはuintでキャストしたidを渡しています
	err = cruds.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	//errorがない場合に、{"message":"OK!!"}を返す。
	c.JSON(http.StatusOK, gin.H{
		"message": "OK!!",
	})
}
```
これでTodoRouterの定義が完了！！

## Lv.12 TodoRouterをRouterに登録

routers/main.goにTodoRouterを登録しよう。

```routers/main.go
package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"*",
		},
	}))
	
	//TodoRouterのすべてのパスを"/todo"で始まるようにグループ化
	todo_router := r.Group("/todo")
	//InitTodoRouterを実行
	InitTodoRouters(todo_router)
}
```
これで一通り完成！！


## Lv. 13 APIをたたいてみる

完成したAPIたたいてみよう。

まず**go run main.go**でサーバを起動しよう。

そしてコマンドで以下を実行しよう
```
$ curl -X POST -H "Content-Type: application/json" -d '{"content":"Goの勉強をする"}' localhost:8000/todo
```

これを実行して{"id":1,"content":"Goの勉強をする"}みたいにかえってきたら成功！！
他にも"Goの勉強をする"の部分を変えて実行してみよう。

そして
```
$ curl localhost:8000/todo
```
を実行すると、さっき登録したTodoが配列で帰ってくるはず！

さらに
```
$ curl -X DELETE localhost:8000/todo/1
```
1の部分に登録されているTodoのidを入れると、{"message": "OK!!"}が返ってきて、
Todoが削除されているはずです。

## Lv. 14 勝利
ここまでくればもう大丈夫、これからはバックエンド要員として働いてもらいます。
ソースコードはGo-Questリポジトリのdevelopブランチを覗いてみてください。

## おまけ フロントを作ってみる。
クッソ適当にフロントのほうも実装しました。
React × TypeScriptで書きました。
主要なファイルだけ載せときます。

App.tsx
```App.tsx
import './App.css';
import { Todos } from './components/Todos';



function App() { 
  return (
    <div className="App">
      <h1>Todo App with Golang</h1>
      <Todos/>
    </div>
  );
}

export default App;
```

components/Todos.tsx
```components/Todos.tsx
import { useState, useEffect } from "react"
import axios from "axios"

interface Todo {
    id: number,
    content: string,
}

export const Todos = () => {
    const [inputValue, setInputValue] = useState<string>("")
    const [todos, setTodos] = useState<Todo[]>([])

    useEffect(() => {
        const FetchData = async() => {
            await axios.get('http://localhost:8000/todo')
            .then(res => res.data)
            .then(data => {setTodos(data)})
            .catch(error => {
                console.error(error)
            })
        }
        FetchData()
    },[])

    const handleSubmit = async() => {
        console.log(inputValue)
        await axios.post('http://localhost:8000/todo',{
            "content": inputValue
        })
        .then(res => res.data)
        .then(data => {
            console.log('Success:', data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });

        setInputValue("")

        await axios.get('http://localhost:8000/todo')
        .then(res => res.data)
        .then(data => {setTodos(data)})
        .catch(error => {
            console.error("Error:",error)
        })
    }

    const handleDelete = async(id:number) => {
        await axios.delete(`http://localhost:8000/todo/${id}`)
        .then(res => res.data)
        .then(data => {
            console.log(data)
        })
        .catch(error => {
            console.error(error)
        })

        await axios.get('http://localhost:8000/todo')
        .then(res => res.data)
        .then(data => {setTodos(data)})
        .catch(error => {
            console.error("Error:",error)
        })
    }
    return (
        <div className="form-wrapper">
            <div>
                <input value={inputValue} type="text" id="content" onChange={(e) => {setInputValue(e.target.value)}}/>
                <button onClick={handleSubmit}>追加</button>
            </div>
            <ul className="todos">
                {
                    todos.map((todo) => {
                        return(
                            <li key={todo.id}>
                                {todo.content}
                                <button onClick={() => {handleDelete(todo.id)}}>削除</button>
                            </li>
                        )
                    })
                }
            </ul>
        </div>
    )
}
```

CSSはだいきらいなのでほぼ書いていません。自分好みにカスタマイズしてみてください。
フロントのソースコードはdevelopブランチのfrontendのディレクトリにあります。


## Comming Soon... ?　Docker
