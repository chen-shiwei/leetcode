package _042

import "fmt"

func Trap(height []int) int {

	l := len(height)
	var level int

	var ok bool
	var rain int
	var count int
	for {
		isChanged := false
		ok = false

		var tempRain int // 临时 不确定
		for i := 0; i < l; i++ {

			if height[i] <= level /*无阻隔*/ {
				if ok /*之前阻隔*/ {
					tempRain++ /*不缺确定后面是否有阻隔*/
				}
			} else {
				if ok { /*确定有阻隔 加入正式雨*/
					rain += tempRain
					tempRain = 0
				}
				ok = true
				isChanged = true
			}

			count++

		}

		if !isChanged {
			break
		}
		level++

	}

	fmt.Println(count)
	return rain

}
