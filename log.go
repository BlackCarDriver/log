package log

import (
	"fmt"
	"sync"
	"os"
	"strings"
	"time"
)

var (
	logPath string
	logNames map[string]bool
)

func init(){
	logPath = ""
	logNames = make(map[string]bool, 0)
}

func SetLogPath(path string){
	if logPath != "" {
		panic("LogPath can only set one time!")
	}
	if !strings.HasSuffix(path, "/"){
		path += "/"
	}
	_, err := os.Stat(path)
	if err!=nil && os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModeDir )
		if err!=nil{
			panic(err)
		}
		_, err = os.Stat(path)
	}
	if err!=nil {
		panic(err)
	}
	logPath = path
}

type Logger struct { 
	name string
	flag uint64
	links uint64
	sync.Mutex
	file *os.File 
}

//create an Logger object
func NewLogger(name string) *Logger{
	tl := new(Logger)
	if logPath == ""{
		panic("You must use SetLogPath to set path first !")
	}
	if len(name) == 0 {
		ts := fmt.Sprint("name is nill !\n")
		panic(ts)
	} 
	if logNames[name] {
		ts := fmt.Sprintf("name %s already used !", name)
		panic(ts)
	}
	file, err := os.Create(logPath + name)
	if err!=nil {
		panic(err)	
	}
	logNames[name] = true
	tl.name = name
	tl.file = file
	return tl
}

//reference fmt.Sprintf() and custom by flag
func (l* Logger)Write(format string, any ...interface{}){
	str := ""
	fmt.Print(str)
	switch l.flag {
	case 1:		//add a time stamp 
		prefix := time.Now().Format("060102-15:04:05    ")
		str = fmt.Sprintf(prefix + format, any...)
	case 2:		//add link number 
		prefix := fmt.Sprintf("%-7d", l.links)
		str = fmt.Sprintf(prefix + format, any...)
		l.links ++
	default:
		str = fmt.Sprintf(format, any...)
	}
	l.Mutex.Lock()
	l.file.WriteString(str)
	l.Mutex.Unlock()
}

//custom the format of log style
func (l *Logger)SetFlag(flag uint64){
	l.flag = flag
}

//clear all content in logfile 
func (l *Logger)Clear(){
	file, err := os.Create(logPath + l.name)
	if err!=nil {
		panic(err)	
	}
	l.links = 0
	l.file = file
}


