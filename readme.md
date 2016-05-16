`Loggo` is a totally new implementation of [log4go](http://github.com/kimiazhu/log4go).

 
### Features 

- log4go风格的配置文件

- 异步日志写入

- 按天/按文件大小/按日志行数滚动日志

- 按包名定义的logger

- 日志级别支持

- Fatal级别日志出现异常时打印堆栈信息

- 重用现有日志文件。如果当前已经存在日志文件，并且不需要滚动，则在启动loggo的时候会使用现有的日志文件。

- 专门的Access级别日志，记录每次请求的详细信息

- 自动重新加载配置文件
