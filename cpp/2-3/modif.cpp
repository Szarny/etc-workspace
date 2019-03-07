#include <iostream>
#include <cmath>

using namespace std;

int main() {
    double x;

    cout << "Real Number: ";
    cin >> x;

    if(double mod = fmod(x, 10)) {
        cout << "Mod is " << mod << endl;
    } else {
        cout << "No mod" << endl;
    }
}