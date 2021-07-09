package recursion

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

// 猜数字游戏
func TestGuessNum(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)

	t.Log("请想一个0-1000的随机数字，3秒后游戏开始")
	time.Sleep(time.Duration(2) * time.Second)

	start := 0
	end := 1000
	n := 0
	var middle int
	for {
		if start == end {
			t.Logf("猜到了，这个数是：%d", start)
			t.Logf("一共用了 %d 次", n)
			break
		}

		middle = (end - start) / 2
		t.Log("这个数组是不是大于等于500？(输入y代表是 n代表不是)")
		answer, err := reader.ReadString('\n')
		if err == nil {
			answer = strings.Replace(answer, "\n", "", -1)
			if answer != "y" && answer != "n" {
				t.Log("您只能回答y与n，代码是和否")
				continue
			}

			n++
			if answer == "y" {
				start = middle
			} else {
				end = middle - 1
			}
		}
	}
}
