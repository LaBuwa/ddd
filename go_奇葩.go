func learnFunctionFactory() {
    // 空行分割的两个写法是相同的，不过第二个写法比较实用
    fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

    d := sentenceFactory("原谅")
    fmt.Println(d("当然选择", "她！"))
    fmt.Println(d("你怎么可以", "她？"))
}

// Decorator在一些语言中很常见，在go语言中，
// 接受参数作为其定义的一部分的函数是修饰符的替代品
func sentenceFactory(mystring string) func(before, after string) string {
    return func(before, after string) string {
        return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
    }
}

func learnDefer() (ok bool) {
    // defer表达式在函数返回的前一刻执行
    defer fmt.Println("defer表达式执行顺序为后进先出（LIFO）")
    defer fmt.Println("\n这句话比上句话先输出，因为")
    // 关于defer的用法，例如用defer关闭一个文件，
    // 就可以让关闭操作与打开操作的代码更近一些
    return true
}
