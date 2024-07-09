package main

import (
	"github.com/dingyuqi/go_numcalc/src/conversion"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	data := make([]float32, 0)
	for i := 10; i < 100000000; i++ {
		data = append(data, float32(i))
	}
	c, err := conversion.NewCalculator()
	if err != nil {
		t.Log(err)
		return
	}
	st := time.Now()

	logFloat32, err := c.LogFloat32(data)
	if err != nil {
		return
	}
	t.Log("计算结果透视: ", logFloat32[:5])
	t.Log("耗时: ", time.Since(st).Seconds(), "秒")
}
