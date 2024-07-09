package conversion

// Standardization 最小最大标准化接口
type Standardization interface {
	StandardizationInt32([]int32) ([]float32, error)
	StandardizationFloat32([]float32) ([]float32, error)
}

// ZScore Z-Score标准化接口
type ZScore interface {
	ZScoreInt32([]int32) ([]float32, error)
	ZScoreFloat32([]float32) ([]float32, error)
}

// Log 对数计算接口
type Log interface {
	LogInt32([]int32) ([]float32, error)
	LogFloat32([]float32) ([]float32, error)
}

// Square 平方根计算接口
type Square interface {
	SquareInt32([]int32) ([]float32, error)
	SquareFloat32([]float32) ([]float32, error)
}
