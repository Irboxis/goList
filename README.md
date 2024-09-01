## GoList

GoList 是一个用于处理切片的 Go 库，提供类似 JavaScript 数组方法的操作功能 。

- 特性 : 
    - Add: 在指定索引处插入元素。
    - AddAll: 从指定索引开始插入多个元素。
    - Push: 在切片末尾追加一个或多个元素。
    - Shift: 在切片开头插入一个或多个元素。
    - Update: 根据回调函数更新切片中的每个元素。
    - ToUpdate: 返回一个新的切片，并根据回调函数进行更新。
- 安装 : 
    - 使用 go get 安装 ：
    ```bash
        go get github.com/Irboxis/goList
    ```
- 使用方法 : 
    - 以下是一些基本用法示例 ：

    ```go
    package main
    
    import (
        "fmt"
        "github.com/Irboxis/goList"
    )
    
    func main() {
    s := golist.New(1, 2, 3, 4, 5)
    
        // 添加元素
        err := s.Add(2, 10)
        if err != nil {
            fmt.Println("Add error:", err)
        }
    
        // 添加多个元素
        err = s.AddAll(3, 20, 30, 40)
        if err != nil {
            fmt.Println("AddAll error:", err)
        }
    
        // 更新切片元素
        err = s.Update(func(elem int, index int, slice *[]int) bool {
            (*slice)[index] = elem * 2
            return true
        })
        if err != nil {
            fmt.Println("Update error:", err)
        }
    
        fmt.Println(s.Slice)  // 输出: [1 2 10 40 60 80 8 10]
    }
    ```

- 测试
    - 你可以通过以下命令运行库中的测试：

    ```bash
    复制代码
    go test ./...
    ```