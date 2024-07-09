package test

import "math"

func DataPositiveInt32() []int32 {
	var data []int32
	for i := 1; i <= 10; i++ {
		data = append(data, int32(i))
	}
	return data
}

func DataPositiveFloat32() []float32 {
	var data []float32
	for i := 1; i <= 10; i++ {
		data = append(data, float32(i))
	}
	return data
}

func DataNegativeInt32() []int32 {
	var data []int32
	for i := -10; i < 0; i++ {
		data = append(data, int32(i))
	}
	return data
}

func DataNegativeFloat32() []float32 {
	var data []float32
	for i := -10; i < 0; i++ {
		data = append(data, float32(i))
	}
	return data
}

func DataSameInt32() []int32 {
	return []int32{1, 1, 1, 1, 1}
}

func DataSameFloat32() []float32 {
	return []float32{1, 1, 1, 1, 1}
}

func DataSpecialInt32() []int32 {
	return []int32{int32(math.NaN()), int32(math.Inf(1)), int32(math.Inf(-1)), int32(math.Inf(0)), 0}
}
func DataSpecialFloat32() []float32 {
	return []float32{float32(math.NaN()), float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.Inf(0)), 0}
}
