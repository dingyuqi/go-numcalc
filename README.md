# go_numcalc

**Read this in other languages: [简体中文](./README.zh-CN.md)**

The go_numcalc package is a package developed in Go language. Its main function is to perform basic numerical processing
such as data conversion, data grouping, and data smoothing for numerical type data.
All functions support only two data types: Int32 and Float32.

> [!TIP]
> This package is written in Go language and contains part of C++, that is, the cgo part

<!-- TOC -->

* [go_numcalc](#go_numcalc)
    * [Installation](#installation)
    * [Dependencies](#dependencies)
    * [Usage](#usage)
    * [Project structure](#project-structure)
    * [Development](#development)
        * [First phase functions](#first-phase-functions)
        * [Phase II functions (TODO)](#phase-ii-functions-todo)
    * [License](#license)

<!-- TOC -->

## Installation

Use the `go get` to install go_numcalc.

```shell
go get github.com/dingyuqi/go_numcalc
```

## Dependencies

There are two main external libraries used in the NumCalc package:

1. Go language: [Gonum](https://www.gonum.org/)
2. C++ language: [Armadillo](https://arma.sourceforge.net/docs.html#top)

Armadillo does not need to be installed and is called in Cgo as a static library.

## Usage

The following is a simple example of the `LogInt32()` method in conversion. The usage of other functions is consistent
with this Demo.

1. Call the `NewCalculator()` method in conversion to initialize a numerical conversion object.
2. Call the `LogInt32()` method. This method will return the logarithm calculation result of the corresponding subscript
   of the slice `data`.

```go
package main

import (
	"github.com/dingyuqi/go_numcalc/src/conversion"
	"log"
)

func main() {
	data := []int32{1, 2, 3, 4, 5}
	c, err := conversion.NewCalculator()
	result, err := c.LogInt32(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("result is: ", result)
}
```

> [!TIP]
> This test code is in `example/example_test.go` and can be run directly.

## Project structure

1. **src**: src code
    - binning: data binning functions
    - conversion: data conversion functions
    - outlier: data outlier functions
    - smoothness: data smoothness functions
2. **pkg**: Used to place static compilation files and related header files of the C++ language Armadillo library.
3. **example**: Used to place Cgo calling samples and Go language calling samples.
4. **test**: Unit test data and test cases.

> [!TIP]
> If you want to test Cgo, you can directly execute the following command in the example folder:
>   ```shell
>    go build ./example.exe
>   ```
> `example_test.go` contains a pure Go language implementation of the same function (logarithmic calculation), used to
> compare the calculation speed of Cgo.

## Development

As of August 2023, only the first phase of functions has been implemented, and the implementation language is all based
on the Go language.

### First phase functions

| Serial number | Type             | Function                            | Detailed description of function                                                                                                                                                                                                                                  | Remarks                                                                                                       |
|---------------|------------------|-------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| 1             | Data conversion  | Minimum and maximum standardization | Perform a linear transformation on the data series so that the processed data all fall within the interval [0, 1]                                                                                                                                                 |                                                                                                               |
| 2             | Data conversion  | Z-score standardization             | Subtract the mean and divide by the variance for each data point in the data series so that the processed data approximately conforms to the standard normal distribution of (0, 1)                                                                               |                                                                                                               |
| 3             | Data conversion  | Logarithmic transformation          | $y = $\log_{base}{x}$                                                                                                                                                                                                                                             | 1. Negative value processing </br> 2. base value                                                              |
| 4             | Data conversion  | Square root transformation          | $y = \sqrt{x}$                                                                                                                                                                                                                                                    | Negative value processing                                                                                     |
| 5             | Data Grouping    | Cluster Grouping                    | Use cluster analysis methods to group data points into clusters with similar characteristics. Cluster grouping can be used to discover clustering patterns and categories in the data, which is useful for data mining and classification tasks.                  | 1. Clustering method (random_subset, static_subset, etc.) 2. Number of clusters                               | |
| 6             | Data Grouping    | Equal Width Grouping                | Divide the value range of the data into intervals of equal width. This method is simple and intuitive, but may not reflect the distribution characteristics of the data well, especially when there are imbalanced data or outliers.                              | Group Width                                                                                                   |
| 7             | Data Grouping    | Equal Frequency Grouping            | Divide the data into groups containing the same number of data points. This method can better consider the distribution characteristics of the data, but for data containing a large number of repeated values, it may cause some groups to have the same values. | Number of Groups                                                                                              |
| 8             | Data Grouping    | Grouping Based on Statistics        | Divide the data into groups based on the quantiles of the data. Common methods include quartile grouping, decile grouping, etc. This method can divide data into groups with the same data density, which is more effective for skewed distribution data.         | Grouping conditions _(Not implemented temporarily due to overlap with the equal frequency grouping function)_ |
| 9             | Outlier judgment | Standard deviation                  | By calculating the difference between the standard deviation of the data point and the mean, values ​​exceeding a certain threshold are considered outliers. Usually, values ​​exceeding 3 times the standard deviation are considered outliers                   | 1. true indicates an outlier</br>2. false indicates a non-outlier</br> 3. Threshold for outlier judgment      |
| 10            | Outlier judgment | Box plot                            | According to the quartiles and outlier range of the data, values ​​beyond the upper and lower boundaries are considered outliers.                                                                                                                                 | 1. true indicates an outlier</br>2. false indicates a non-outlier                                             |

### Phase II functions (TODO)

| Serial number | Type           | Function                 | Function description                                                                                                                                                                                                                                                                                                                              | Remarks                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
|---------------|----------------|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1             | Data smoothing | Wavelet filtering        | Decompose and reconstruct signals by applying wavelet transform to remove noise or mutations and retain important features in the signal. Wavelet filtering provides better analysis and processing capabilities in both time and frequency domains.                                                                                              | Different base functions have a great impact on the results. Different data need to choose different base functions and frequency ranges according to the analysis requirements. </br> 1. Wavelet basis function (Daubechies wavelet, Haar wavelet, Morlet wavelet)</br>2. Scale parameter (determines the scaling factor of each wavelet basis function in the wavelet transform; a smaller scale parameter can capture higher frequency and detailed signal characteristics, while a larger scale parameter can capture lower frequency and overall trend signal characteristics)</br>3. Decomposition level (determines the order of wavelet transform; a higher decomposition level can provide more detailed frequency and scale information)</br>4. Threshold processing method (keep/discard) |
| 2             | Data smoothing | Moving average           | This method smoothes the data by calculating the average value within a certain window size around the data point. The window size determines the degree of smoothing, and a larger window will smooth more fluctuations. Common moving averages include simple moving average and weighted moving average.                                       | If the boundary cannot completely construct a window that meets the window size, the data points at these boundaries are usually removed in the output result, resulting in unequal lengths. 1. Window size </br>2. Weight </br>Boundary processing method                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| 3             | Data smoothing | Exponential smoothing    | Exponential smoothing is a recursive smoothing method that gives a higher weight to recent data. The weight of past observations is controlled by specifying a smoothing coefficient, where the larger the smoothing coefficient, the greater the impact on recent data. Exponential smoothing methods are often used to smooth time series data. | Smoothing factor                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| 4             | Data smoothing | Savitzky-Golay smoothing | This is a smoothing method based on polynomial fitting, which smoothes data by fitting neighboring data around the data point to a polynomial curve. The Savitzky-Golay smoothing method can retain the overall shape and trend of the data and has a good noise suppression effect.                                                              | 1. Smoothing window </br>2. Polynomial order </br> 3. Derivative order (optional)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| 5             | Data smoothing | Loess smoothing          | Similar to Lowess smoothing, Loess smoothing is also a nonparametric local regression method. It smoothes data by fitting a polynomial to the neighboring data around the data point. Unlike Lowess smoothing, Loess smoothing uses adaptive weighted least squares to better handle nonlinear relationships in the data.                         | 1. Smoothing coefficient (controls the weight given to past observations) </br>2. Weighting function (default in the library)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | |
| 6             | Data smoothing | Lowess smoothing         | Lowess smoothing is a nonparametric local regression method that smooths data by fitting a local linear regression model. The method uses weighted least squares to estimate the smoothed value of a data point, with weights assigned based on how far away the data point is.                                                                   | 1. Smoothing coefficient (controls the weight given to past observations) </br>2. Weighting function (default in the library)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |

## License

[MIT](https://choosealicense.com/licenses/mit/) License