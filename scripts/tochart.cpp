#include <iostream>
#include <vector>
#include <map>
#include <fstream>
#include <algorithm>

using namespace std;

int main(int argc, char **argv) {

    string curType = string(argv[2]);
    string curAlgo = string(argv[3]);

    ifstream fin;
    fin.open(argv[1]);
    
    string type;
    string algo;
    long size;
    int samples;
    double ratio;

    map<long, map<int, double> > m;

    while (fin >> type >> algo >> size >> samples >> ratio) {
        if (type != curType || algo != curAlgo) {
            continue;
        }

        if (m.find(size) == m.end()) {
            m[size] = map<int, double>();
        }

        m[size][samples] = ratio;
    }

    bool titlePrinted = false;

    for (auto const& t1 : m) {

        if (!titlePrinted) {
            cout << curType << "_" << curAlgo << '\t';
            for (auto const& t2 : t1.second) {
                cout << t2.first << '\t';
            }
            cout << endl;
            titlePrinted = true;
        }
        cout << t1.first << '\t';   // string (key)
              
        for (auto const& t2 : t1.second) {
            cout << t2.second << '\t';
        }
        cout  << std::endl ;
    }


    return 0;
}
