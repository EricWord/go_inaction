package goroutine

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午15:50
   @since Go1.13.8
*/

//鉴于通道本身就是一个并发安全的序列，所以可以用作ID generator 、 Pool等用途
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	//返回
	case v = <-p:
	default:
		v = make([]byte, 10)

	}
	return v
}

func (p pool) put(b []byte) {
	select {
	//放回
	case p <- b:
	default:
		//放回失败，放弃
	}
}
