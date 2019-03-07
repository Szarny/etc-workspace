#include <iostream>
#include <iomanip>

using namespace std;

int main() {
    int x, y;

    for(x = 1; x <= 9; x++){
        for(y = 1; y <= 9; y++){
            cout << setw(3) << x * y;
        }
        cout << endl;
    }
}