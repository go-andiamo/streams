package streams

import "strings"

var StringComparator = NewComparator[string](func(v1, v2 string) int {
	return strings.Compare(v1, v2)
})

var StringInsensitiveComparator = NewComparator[string](func(v1, v2 string) int {
	return strings.Compare(strings.ToUpper(v1), strings.ToUpper(v2))
})

var IntComparator = NewComparator[int](func(v1, v2 int) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Int8Comparator = NewComparator[int8](func(v1, v2 int8) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Int16Comparator = NewComparator[int16](func(v1, v2 int16) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Int32Comparator = NewComparator[int32](func(v1, v2 int32) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Int64Comparator = NewComparator[int64](func(v1, v2 int64) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var UintComparator = NewComparator[uint](func(v1, v2 uint) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Uint8Comparator = NewComparator[uint8](func(v1, v2 uint8) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Uint16Comparator = NewComparator[uint16](func(v1, v2 uint16) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Uint32Comparator = NewComparator[uint32](func(v1, v2 uint32) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Uint64Comparator = NewComparator[uint64](func(v1, v2 uint64) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Float32Comparator = NewComparator[float32](func(v1, v2 float32) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})

var Float64Comparator = NewComparator[float64](func(v1, v2 float64) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
})
