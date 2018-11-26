#include <dirent.h>
#include <sys/types.h>
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <iterator>

std::vector<std::string> open(std::string path = ".") {

    DIR*    dir;
    dirent* pdir;
    std::vector<std::string> files;

    dir = opendir(path.c_str());

    while (pdir = readdir(dir)) {
        files.push_back(pdir->d_name);
    }

    return files;
}

void process(std::string fn) {

    std::ifstream fin;
    fin.open(fn.c_str());

    std::string temp;
    std::string algo;
    int granularity;
    int reportThres;
    long cacheSize;
    int buckets;
    int samples;
    int pruningItems;
    int admissionThres;

    int num1, num2, num3, num4;

    fin >> temp >> temp;
    fin >> temp >> algo;
    fin >> temp >> granularity;
    fin >> temp >> temp >> reportThres;
    fin >> temp >> temp >> cacheSize;
    fin >> temp >> buckets;
    fin >> temp >> samples;
    fin >> temp >> temp >> pruningItems;
    fin >> temp >> temp >> admissionThres;

    for (int i = 0; i < 249; i++) {
        fin >> temp >> temp >> temp >> temp;
    }

    fin >> temp >> num1 >> num2 >> temp;

    for (int i = 0; i < 49; i++) {
        fin >> temp >> temp >> temp >> temp;
    }

    int reqNum;
    std::string type = "p";

    fin >> reqNum >> num3 >> num4 >> temp;

    if (reqNum != num3) {
        type = "o";
    }

    if (admissionThres != 0) {
        algo = "h2-ad";
    }

    double ratio = (1.0 * (num4 - num2)) / (1.0 * (num3 - num1));

    std::cout << type << " " << algo << " " << cacheSize << " " << samples << " ";
    std::cout << ratio << std::endl;
}

int main(int arc, char **argv) {
    
    std::vector<std::string> f;
    //std::string path = "/home/ruogu/Desktop/capstone/experiments/runs/";
    std::string path = std::string(argv[1]);
    f = open(path); // or pass which dir to open

    for (auto itr = f.begin() + 2; itr != f.end(); itr++) {
        process(path+*itr);
    }

    return 0;
}
