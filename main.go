package main

import (
	"fmt"
	"sync"
	"os"
	"strings"
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

type logger struct {
	name string
	sync.Mutex
	file *os.File 
}

func NewLogger(name string) *logger{
	tl := new(logger)
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
	file, err := os.Open(logPath + name)
	if err!=nil && os.IsNotExist(err){
		file, err = os.Create(logPath + name)
		if err != nil{
			panic(err)
		}
	}
	logNames[name] = true
	tl.name = name
	tl.file = file
	return tl
}

func (l* logger)Write(format string, any ...interface{}){
	str := fmt.Sprintf(format, any...)
	fmt.Println(str)
	l.file.WriteString(str)
}

func main(){
	SetLogPath("./temp")
	mylog := NewLogger("test2.log")
	mylog.Write("hahahahahah")
}