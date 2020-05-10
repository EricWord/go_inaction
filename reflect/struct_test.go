package reflect

import (
	"reflect"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午21:45
   @since Go1.13.8
*/

type user struct {

	name string
	age int
}


type manager struct {
	user
	title string
}
func TestStruct(t *testing.T){

	//var m manager
	//ty:=reflect.TypeOf(&m)


	//if ty.Kind()


}