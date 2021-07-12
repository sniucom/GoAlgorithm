package hashmap

import "testing"

func TestBitmap(t *testing.T) {
	bitmap := NewBitmap(100000000)
	bitmap.Add("采蘑菇的小姑娘")
	if bitmap.Exists("采蘑菇的小姑娘") {
		t.Log("采蘑菇的小姑娘 存在")
	}

	if !bitmap.Exists("爱吃鱼的大脸猫") {
		t.Log("爱吃鱼的大脸猫 不存在")
	}
}
