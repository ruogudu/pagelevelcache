#include <iostream>
#include <vector>
#include <fstream>
#include <algorithm>

using namespace std;

int main(int argc, char **argv) {

    ifstream fin;
    fin.open(argv[1]);
    vector<string> lines;

    string line;
    while (getline(fin, line)) {
        lines.push_back(line);
    }

    sort(lines.begin(), lines.end());

    for (auto itr = lines.begin(); itr != lines.end(); itr++) {
        cout << *itr << endl;
    }

    return 0;
}
