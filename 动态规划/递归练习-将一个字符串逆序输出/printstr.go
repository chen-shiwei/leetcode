package 递归练习_将一个字符串逆序输出

func reversePrintStr(s string) {
	l := len(s)
	println(string(s[l-1]))
	if l == 1 {
		return
	}
	reversePrintStr(s[:l-1])
	return
}
