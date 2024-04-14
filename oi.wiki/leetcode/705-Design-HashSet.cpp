#include <algorithm>
#include <list>
#include <vector>

class MyHashSet {
    std::vector<std::list<int>> data;
    static const int base = 769;
    static int hash_func(int key) { return key % base; }

   public:
    MyHashSet() { data.assign(base, std::list<int>()); }

    void add(int key) {
        int hash = hash_func(key);
        auto it = std::find(data[hash].begin(), data[hash].end(), key);
        if (it != data[hash].end()) {
            return;
        }
        data[hash].push_back(key);
    }

    void remove(int key) {
        int hash = hash_func(key);
        auto it = std::find(data[hash].begin(), data[hash].end(), key);
        if (it == data[hash].end()) {
            return;
        }
        data[hash].erase(it);
    }

    bool contains(int key) {
        int hash = hash_func(key);
        auto it = std::find(data[hash].begin(), data[hash].end(), key);
        return it != data[hash].end();
    }
};
