package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
   @author eric
   @email 2635863525@qq.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午20:42
   @since Go1.13.8
*/

type X int

func TestReflect(t *testing.T) {

	var a X = 10
	ty := reflect.TypeOf(a)
	fmt.Println(ty.Name(), ty.Kind())

}

func TestArray(t *testing.T) {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))

	fmt.Println(a, m)

}

func TestKind(t *testing.T) {
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(tx, tp, tx == tp)

	fmt.Println("------------------")

	fmt.Println(tx.Kind(),tp.Kind())
	fmt.Println("------------------")

	fmt.Println(tx==tp.Elem())

}
