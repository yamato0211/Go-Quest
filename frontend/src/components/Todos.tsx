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