package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // 这里开始不能使用 v
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )

    // 27 >= 20
    // 9 20
}

/*
    if 和 else

    在 if 的便捷语句定义的的变量同样可以在任何对应的 else 块中使用。
    （提示：两个 pow 调用都在 main 调用 fmt.Println 前执行完毕了。）
*/
