package xmath

// 来自 go 1.17
const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	MaxInt  = 1<<(intSize-1) - 1  // MaxInt32 or MaxInt64 depending on intSize.
	MinInt  = -1 << (intSize - 1) // MinInt32 or MinInt64 depending on intSize.
	MaxUint = 1<<intSize - 1      // MaxUint32 or MaxUint64 depending on intSize.
)
