#include <iostream>

using namespace std;

int main() {
    int a = 10;
    int b = 5;

    double res = (a + b) / 2;
    double res_casted = static_cast<double>(a+b) / 2;

    cout << res << endl;
    cout << res_casted << endl;
}