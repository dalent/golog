# golog
golang log modules
输出格式类似于

[2015-03-11 22:23:05][log_test.go:46][ERROR] me


1. file log
文件log的时候，可以设置文件大小，当超过大小的时候会重新换一个文件，还有可以设置最大保留天数
```
package main
import(
  "github.com/dalent/golog"
)
func main(){ 
     writer := golog.NewFileWriter()
     writer.Prefix("appname")//默认是undefined，表示文件前缀
    log := golog.New(writer)
    log.SetLevel(golog.LDEBUG)  //最小纪录日志级别
    log.SetCallDepth(2)    //纪录调用的该函数的文件与行数
    log.Debug("%s", "me")
    }
```

2. console 有颜色的打印
```
package main
import(
  "github.com/dalent/golog"
)
func main(){ 
    writer := golog.NewConsoleWriter()
    log := golog.New(writer)
    log.Error("%s", "me")
    }
```



