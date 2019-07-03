# log

----------

**介绍**
这是一个简易的go package,  用来代替go原有的log包，优点是能够很方便地
创建许多个日志对象，用来将不同类型的输出写到不同的文件，还有使用方法简单，容易扩展出多种不同的日志记录格式。
不断更新中。。。。

----------

**使用案例**

    import (
    "github.com/BlackCarDriver/log" 
    ”time"
    )

    func test_log(){
    	//设置存放日志的目录
    	log.SetLogPath("./logpath")
    	//创建日志对象
    	msg := log.NewLogger("message")
    	//设计日志格式
    	msg.SetFlag(1)
    	for i:=0;i<20;i++ {
    		//记录日志
    		msg.Write("It is %d log ... \n", i)
    	}
    	time.Sleep(5*time.Second)
    	//清空日志
    	msg.Clear()
    	msg.Write("test complete!")
    }

----------

lastChange:2019/7/3 18:08:00 