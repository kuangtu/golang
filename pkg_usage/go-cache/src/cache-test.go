package main


import (
    "fmt"
    "github.com/patrickmn/go-cache"
    "time"
)

func main() {
    //创建一个缓存
    c := cache.New(5*time.Minute, 10*time.Minute)
    //设置
    c.Set("foo", "bar", cache.DefaultExpiration)
    
    foo, found := c.Get("foo")
    if found {
        fmt.Println(foo)
    }

}
