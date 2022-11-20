package streams

import "strings"

var (
	StringComparator            = _StringComparator            // StringComparator is a pre-made comparator for comparing strings
	StringInsensitiveComparator = _StringInsensitiveComparator // StringComparator is a pre-made comparator for comparing strings (case insensitive)
	IntComparator               = _IntComparator               // IntComparator is a pre-made comparator for comparing int
	Int8Comparator              = _Int8Comparator              // Int8Comparator is a pre-made comparator for comparing int8
	Int16Comparator             = _Int16Comparator             // Int16Comparator is a pre-made comparator for comparing int16
	Int32Comparator             = _Int32Comparator             // Int32Comparator is a pre-made comparator for comparing int32
	Int64Comparator             = _Int64Comparator             // Int64Comparator is a pre-made comparator for comparing int64
	UintComparator              = _UintComparator              // UintComparator is a pre-made comparator for comparing uint
	Uint8Comparator             = _Uint8Comparator             // Uint8Comparator is a pre-made comparator for comparing uint8
	Uint16Comparator            = _Uint16Comparator            // Uint16Comparator is a pre-made comparator for comparing uint16
	Uint32Comparator            = _Uint32Comparator            // Uint32Comparator is a pre-made comparator for comparing uint32
	Uint64Comparator            = _Uint64Comparator            // Uint64Comparator is a pre-made comparator for comparing uint64
	Float32Comparator           = _Float32Comparator           // Float32Comparator is a pre-made comparator for comparing float32
	Float64Comparator           = _Float64Comparator           // Float64Comparator is a pre-made comparator for comparing float64
)

var (
	_StringComparator = NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(v1, v2)
	})
	_StringInsensitiveComparator = NewComparator[string](func(v1, v2 string) int {
		return strings.Compare(strings.ToUpper(v1), strings.ToUpper(v2))
	})
	_IntComparator = NewComparator[int](func(v1, v2 int) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Int8Comparator = NewComparator[int8](func(v1, v2 int8) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Int16Comparator = NewComparator[int16](func(v1, v2 int16) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Int32Comparator = NewComparator[int32](func(v1, v2 int32) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Int64Comparator = NewComparator[int64](func(v1, v2 int64) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_UintComparator = NewComparator[uint](func(v1, v2 uint) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Uint8Comparator = NewComparator[uint8](func(v1, v2 uint8) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Uint16Comparator = NewComparator[uint16](func(v1, v2 uint16) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Uint32Comparator = NewComparator[uint32](func(v1, v2 uint32) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Uint64Comparator = NewComparator[uint64](func(v1, v2 uint64) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Float32Comparator = NewComparator[float32](func(v1, v2 float32) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
	_Float64Comparator = NewComparator[float64](func(v1, v2 float64) int {
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
		return 0
	})
)
