package main

//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
//锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
//列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
//字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
//

func openLock(deadends []string, target string) int {
	// 先排除边界情况
	const start = "0000"
	if target == start {
		return 0
	}
	var deadendsMap = map[string]bool{}
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}
	if deadendsMap[start] {
		return -1
	}

	// 都转动一次的情况下所有的可能性
	nextStrList := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	// 构建数据
	type pair struct {
		str  string //密码
		step int    //走了多少步
	}

	queue := []pair{{start, 0}}
	passed := map[string]bool{start: true}

	// 遍历
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, item := range nextStrList(p.str) {
			if item == target {
				return p.step + 1
			}
			if deadendsMap[item] {
				continue
			}
			if passed[item] {
				continue
			}
			passed[item] = true
			queue = append(queue, pair{item, p.step + 1})
		}
	}
	return -1

}
