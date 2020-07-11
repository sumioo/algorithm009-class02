package week8

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	set1 := NewBitSet(uint8(3), uint8(5))
	set2 := NewBitSet(uint8(0), uint8(2))
	t.Log(getSet(set1), ",", getSet(set2))

	set3 := and(set1, set2)
	t.Log(getSet(set3))

	set4 := union(set1, set2)
	t.Log(getSet(set4))

	set5 := diff(set1, set2) //差集 取一个集合中另一集合没有的元素
	t.Log(getSet(set5))
	set5 = diff(set2, set1) //差集 取一个集合中另一集合没有的元素
	t.Log(getSet(set5))

	set6 := symmetricDiff(set1, set2) //对称差，取集合 A 和 B 中不属于 A&B 的元素
	t.Log(getSet(set6))

}

func TestXor(t *testing.T) {
	t.Log(3 ^ 0)
	t.Log(^3) //按位取反
	t.Log(3 ^ 3)
}

func TestBitClearLeftN(t *testing.T) {
	//右边n位清零
	x := 11
	n := 2
	y := uint(x) & (^uint(0) << n)
	t.Log(bin(x), bin(y))
}

func TestBitGetN(t *testing.T) {
	//获取第N位
	x := 11
	n := 1
	y := uint(x) >> n & 1
	t.Log(bin(y))
}

func TestBitGetNPow(t *testing.T) {
	//获取第N位幂集
	x := 11
	n := 1
	y := uint(x) & (uint(1) << n)
	t.Log(y, bin(y))
}

func TestBitSetN(t *testing.T) {
	//设置第N位为1
	x := 11
	n := 2
	y := uint(x) | (uint(1) << n)
	t.Log(bin(x), bin(y))
}

func TestBitClearN(t *testing.T) {
	//设置第N位为0
	x := int8(-1)
	n := 3
	y := uint8(x) & (^(uint8(1) << n))
	t.Log(bin(x), bin(y))
}

func TestBitClearRightToN(t *testing.T) {
	//将x最高位到第N位清零
	x := int8(-1)
	n := 2
	y := uint8(x) & (uint8(1)<<n - 1)
	t.Log(bin(x), bin(y))
}

func TestBitClearLowest1(t *testing.T) {
	//清除最低位1
	x := int8(-1)
	y := uint8(x)
	t.Log(bin(x))
	y = uint8(y) & (uint8(y) - 1)
	t.Log(bin(y))

	y = uint8(y) & (uint8(y) - 1)
	t.Log(bin(y))

	y = uint8(y) & (uint8(y) - 1)
	t.Log(bin(y))

	y = uint8(y) & (uint8(y) - 1)
	t.Log(bin(y))

}

func TestBitGetLowest1(t *testing.T) {
	//获取最低位的1
	x := int8(10)
	y := uint8(x)&uint8(-x) - 1
	t.Logf("%08b %d", x, y)
}
