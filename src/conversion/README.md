# Conversion
**Read this in other languages: [简体中文](./README.zh-CN.md)**

This package mainly encapsulates interfaces related to numerical conversion.

## Structure

```go
package conversion

type Calculator struct {
logBase float32 //The base of logarithmic calculation, defaults to natural number e
}

```

You can call ```func NewCalculator(optionalParam ...float32) Calculator ``` to get the object. If the dynamic parameter exists, it is set to logBase. If it does not exist, it is set to natural number e by default.

## Interface

### 1. StandardizationInt32([]int32) ([]float32, error)

Maximum and minimum value standardization.

After processing a set of int32 type data, the data all fall within the range of [0, 1]. After processing, the output is no longer an integer.

The calculation formula is:
$$normalized=(maxValue−minValue)/(value−minValue)$$

> [!IMPORTANT]
> It is not possible to normalize an array with zero range (equal maximum and minimum values).

### 2. StandardizationFloat32([]float32) ([]float32, error)

Maximum and minimum value normalization.

After processing a set of float32 data, the data all fall within the interval [0, 1].

The calculation formula is:
$$normalized=(maxValue−minValue)/(value−minValue)$$

> [!IMPORTANT]
> It is not possible to normalize an array with zero range (equal maximum and minimum values).

### 3. ZScoreInt32([]int32) ([]float32, error)

Z-Score standardization.

The calculation formula is:
$$Z=(X−\mu)/\sigma$$

$X$ is the original data, $\mu$ is the mean value of the data, and $\sigma$ is the standard deviation of the data.

> [!IMPORTANT]
> Cannot normalize an array with zero range (equal maximum and minimum values).

### 4. ZScoreFloat32([]float32) ([]float32, error)

Z-Score normalization.

The calculation formula is:
$$Z=(X−\mu)/\sigma$$

$X$ is the original data, $\mu$ is the mean value of the data, and $\sigma$ is the standard deviation of the data.

### 5. LogInt32([]int32) ([]float32, error)
Calculate the logarithm.
The calculation formula is:
$$log(base, x)$$
> [!TIP]
> base defaults to the natural number e

> [!IMPORTANT]
> Make sure that all input parameters are positive, otherwise:
> Log(+Inf) = +Inf
> Log(0) = -Inf
> Log(x < 0) = NaN
> Log(NaN) = NaN

### 6. LogFloat32([]float32) ([]float32, error)
Calculates logarithm.
Calculation formula:
$$log(base, x)$$
> [!TIP]
> base defaults to natural number e

> [!IMPORTANT]
> Input parameters must be positive, otherwise:
> Log(+Inf) = +Inf
> Log(0) = -Inf
> Log(x < 0) = NaN
> Log(NaN) = NaN

### 7. SquareInt32([]int32) ([]float32, error)
Calculates square root.

Calculation formula:
$$\sqrt{x}$$

> [!IMPORTANT]
> Input parameters must be non-negative real numbers, otherwise:
> Sqrt(+Inf) = +Inf
> Sqrt(±0) = ±0
> Sqrt(x < 0) = NaN
> Sqrt(NaN) = NaN

### 8. SquareFloat32([]float32) ([]float32, error)
Calculate the square root.

The calculation formula is:
$$\sqrt{x}$$

> [!IMPORTANT]
> Ensure that all input parameters are non-negative real numbers, otherwise:
> Sqrt(+Inf) = +Inf
> Sqrt(±0) = ±0
> Sqrt(x < 0) = NaN
> Sqrt(NaN) = NaN