#include <armadillo>
#include "logTransform.hpp"

extern "C" void logTransform(const double* data, double* result, int size) {
     // 将 C 的数组转换为 Armadillo 的向量
     arma::vec dataVec(const_cast<double*>(data), size, false, true);

     // 调用你的 C++ 函数
     arma::vec resultVec = arma::log(dataVec);

     // 将结果复制回 C 的数组
     std::memcpy(result, resultVec.memptr(), size * sizeof(double));
}