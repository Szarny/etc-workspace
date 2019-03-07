#include <typeinfo>
#include <iostream>

using namespace std;

int main() {
    cout << typeid(char).name() << endl;
    cout << typeid(int).name() << endl;
    cout << typeid(long int).name() << endl;
}