package main

import (
	"bytes"
	"github.com/atotto/clipboard"
)

/*
百度体生成器是怎么回事呢？百度体生成器相信大家都很熟悉， 但是百度体生成器是怎么回事呢？下面就让小编带大家一起了解吧。
百度体生成器，其实就是百度体生成器了。 大家可能会感到很惊讶，百度体生成器怎么会呢？... 但事实就是这样，小编也感到非常惊讶。
那么这就是关于百度体生成器的事情了，大家有什么想法呢？欢迎在评论区告诉小编一起讨论哦
*/
func main() {
	str, err := clipboard.ReadAll()
	if err != nil {
		println(err)
	}
	var buffer bytes.Buffer
	buffer.WriteString(str)
	buffer.WriteString("是怎么回事呢？")
	buffer.WriteString(str)
	buffer.WriteString("相信大家都很熟悉， 但是")
	buffer.WriteString(str)
	buffer.WriteString("是怎么回事呢？下面就让小编带大家一起了解吧。")
	buffer.WriteString(str)
	buffer.WriteString("，其实就是")
	buffer.WriteString(str)
	buffer.WriteString("了。大家可能会感到很惊讶，")
	buffer.WriteString(str)
	buffer.WriteString("怎么会呢？... 但事实就是这样，小编也感到非常惊讶。那么这就是关于")
	buffer.WriteString(str)
	buffer.WriteString("的事情了，大家有什么想法呢？欢迎在评论区告诉小编一起讨论哦~")
	err = clipboard.WriteAll(buffer.String())
	if err != nil {
		println(err)
	}

}
