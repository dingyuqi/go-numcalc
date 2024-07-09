package outlier

import (
	"github.com/dingyuqi/go_numcalc/test"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Log(NewCalculator(2.5))
}

func TestCalculator_DetectOutliersByStandardDeviationInt32(t *testing.T) {
	c, _ := NewCalculator(2.5)
	t.Log(c.DetectOutliersByStandardDeviationInt32(test.DataPositiveInt32()))
	t.Log(c.DetectOutliersByStandardDeviationInt32(test.DataNegativeInt32()))
	t.Log(c.DetectOutliersByStandardDeviationInt32(test.DataSpecialInt32()))
	t.Log(c.DetectOutliersByStandardDeviationInt32([]int32{1}))
}

func TestCalculator_DetectOutliersByStandardDeviationFloat32(t *testing.T) {
	c, _ := NewCalculator(2.5)
	t.Log(c.DetectOutliersByStandardDeviationFloat32(test.DataPositiveFloat32()))
	t.Log(c.DetectOutliersByStandardDeviationFloat32(test.DataNegativeFloat32()))
	t.Log(c.DetectOutliersByStandardDeviationFloat32(test.DataSpecialFloat32()))
	t.Log(c.DetectOutliersByStandardDeviationFloat32([]float32{1}))
}

func TestCalculator_DetectOutliersByQuartilesInt32(t *testing.T) {
	c, _ := NewCalculator(2.5)
	t.Log(c.DetectOutliersByQuartilesInt32(test.DataPositiveInt32()))
	t.Log(c.DetectOutliersByQuartilesInt32(test.DataNegativeInt32()))
	t.Log(c.DetectOutliersByQuartilesInt32(test.DataSpecialInt32()))
	t.Log(c.DetectOutliersByQuartilesInt32([]int32{1}))
}

func TestCalculator_DetectOutliersByQuartilesFloat32(t *testing.T) {
	c, _ := NewCalculator(2.5)
	t.Log(c.DetectOutliersByQuartilesFloat32(test.DataPositiveFloat32()))
	t.Log(c.DetectOutliersByQuartilesFloat32(test.DataNegativeFloat32()))
	t.Log(c.DetectOutliersByQuartilesFloat32(test.DataSpecialFloat32()))
	t.Log(c.DetectOutliersByQuartilesFloat32([]float32{1}))
}
