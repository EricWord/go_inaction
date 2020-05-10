package log

import (
	"log"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/5 下午14:24
   @since Go1.13.8
*/

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func TestLog(t *testing.T) {
	//log.Println("message")
	//log.Fatalln("fatal message")
	//log.Panicln("panic message")
}
