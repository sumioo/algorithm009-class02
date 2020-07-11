package week8

/*
& 位运算and
| or
^ xor
&^ 位清空 z = x&^y 若y的某位是1，则z的对应的为等于0，否则等于x的对应位
>> 右移
<< 左移
*/

// BitSet xx
type BitSet uint8

// NewBitSet xx
func NewBitSet(args ...uint8) BitSet {
	var x BitSet
	for _, i := range args {
		x = x | 1<<i
	}
	return x
}

func getSet(set BitSet) []uint {
	nums := []uint{}
	for i := uint(0); i < 8; i++ {
		if set&(1<<i) != 0 {
			nums = append(nums, i)
		}
	}
	return nums
}

func and(x BitSet, y BitSet) BitSet {
	return x & y
}

func union(x BitSet, y BitSet) BitSet {
	return x | y
}

func symmetricDiff(x BitSet, y BitSet) BitSet {
	return x ^ y
}

func diff(x BitSet, y BitSet) BitSet {
	return x &^ y
}

func bin(x interface{}) string {
	var ux uint64
	var n int
	switch x.(type) {
	case int:
		ux = uint64(x.(int))
		n = 64
	case int8:
		ux = uint64(x.(int8))
		n = 8
	case int16:
		ux = uint64(x.(int16))
		n = 16
	case int32:
		ux = uint64(x.(int32))
		n = 32
	case int64:
		ux = uint64(x.(int64))
		n = 64
	case uint:
		ux = uint64(x.(uint))
		n = 64
	case uint8:
		ux = uint64(x.(uint8))
		n = 8
	case uint16:
		ux = uint64(x.(uint8))
		n = 16
	case uint32:
		ux = uint64(x.(uint32))
		n = 32
	case uint64:
		ux = x.(uint64)
		n = 64
	default:
		return ""
	}

	s := []byte{}
	for i := 0; ux != 0 && i < n; i++ {
		if ux&1 == 1 {
			s = append(s, '1')
		} else {
			s = append(s, '0')
		}
		ux = ux >> 1
	}

	if len(s) == 0 {
		return "0000"
	}

	for i := 4 - len(s)%4; i < 3 && i > 0; i-- {
		s = append(s, '0')
	}

	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}
