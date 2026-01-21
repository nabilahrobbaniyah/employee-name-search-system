#include <iostream>
#include <vector>
#include <chrono>

using namespace std;
using namespace chrono;

struct Product {
    int id;
};

// Generate data berurutan agar skenario terkontrol
vector<Product> generate(int n) {
    vector<Product> v;
    v.reserve(n);
    for (int i = 0; i < n; i++) {
        v.push_back({i});
    }
    return v;
}

// Linear search iteratif
int linearIter(const vector<Product>& v, int target) {
    for (int i = 0; i < v.size(); i++) {
        if (v[i].id == target)
            return i;
    }
    return -1;
}

// Linear search rekursif
int linearRec(const vector<Product>& v, int target, int idx) {
    if (idx >= v.size())
        return -1;
    if (v[idx].id == target)
        return idx;
    return linearRec(v, target, idx + 1);
}

// Pengukuran waktu iteratif
double timeIter(const vector<Product>& v, int target) {
    volatile int result; // mencegah optimasi compiler
    auto start = high_resolution_clock::now();
    for (int i = 0; i < 10000; i++) { // pengulangan untuk stabilitas
        result = linearIter(v, target);
    }
    auto end = high_resolution_clock::now();
    return duration<double, milli>(end - start).count();
}

// Pengukuran waktu rekursif
double timeRec(const vector<Product>& v, int target) {
    volatile int result;
    auto start = high_resolution_clock::now();
    for (int i = 0; i < 10000; i++) {
        result = linearRec(v, target, 0);
    }
    auto end = high_resolution_clock::now();
    return duration<double, milli>(end - start).count();
}

int main(int argc, char* argv[]) {
    int n = stoi(argv[1]);
    string scenario = argv[2];

    auto data = generate(n);

    // Kontrol best, average, worst case
    int target;
    if (scenario == "best") {
        target = data[0].id;
    } else if (scenario == "average") {
        target = data[n / 2].id;
    } else { // worst case
        target = -1; // tidak ditemukan
    }

    double itTime = timeIter(data, target);
    double recTime = timeRec(data, target);

    string complexity =
        (scenario == "best") ? "O(1)" : "O(n)";

    cout << "{"
         << "\"n\":" << n << ","
         << "\"scenario\":\"" << scenario << "\","
         << "\"target\":" << target << ","
         << "\"iterative\":" << itTime << ","
         << "\"recursive\":" << recTime << ","
         << "\"complexity\":\"" << complexity << "\""
         << "}";
}
