#include <iostream>
#include <string>

using namespace std;

int main() {
    string name;

    cout << "Name: ";
    getline(cin, name);

    cout << "Name = " << name << endl;
}