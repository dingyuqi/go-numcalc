package src

func Int32ToInt(data []int32) []int {
	result := make([]int, len(data))
	for i, v := range data {
		result[i] = int(v)
	}
	return result
}

func Float32ToFloat64(data []float32) []float64 {
	result := make([]float64, len(data))
	for i, v := range data {
		result[i] = float64(v)
	}
	return result
}

func Int32ToFloat64(data []int32) []float64 {
	result := make([]float64, len(data))
	for i, v := range data {
		result[i] = float64(v)
	}
	return result
}
