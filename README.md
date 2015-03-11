# golog
golang log modules
文件log的时候，可以设置文件大小，当超过大小的时候会重新换一个文件，还有可以设置最大保留天数
1.file log
```
package main
import(
  "github.com/dalent/golog"
)
func main(){ 
     writer := NewFileWriter()
     writer.Prefix("appname")//默认是undefined，表示文件前缀
    log := New(writer)
    log.SetLevel(golog.WARN)  //最小纪录日志级别
    log.SetCallDepth(2)    //纪录调用的该函数的文件与行数
    log.Debug("%s", "me")
    }
```
2. console
```
package main
import(
  "github.com/dalent/golog"
)
func main(){ 
    writer := NewConsoleWriter()
    log := New(writer)
    log.Debug("%s", "me")
    }
```
