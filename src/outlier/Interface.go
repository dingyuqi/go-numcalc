package outlier

// DetectOutlier 离群值判断接口
type DetectOutlier interface {
	DetectOutliersByStandardDeviationInt32([]int32) ([]bool, error)
	DetectOutliersByStandardDeviationFloat32([]float32) ([]bool, error)
	DetectOutliersByQuartilesInt32([]int32) ([]bool, error)
	DetectOutliersByQuartilesFloat32([]float32) ([]bool, error)
}
