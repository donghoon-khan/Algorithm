#include <iostream>
#include <vector>
#include <algorithm>
#include <stdio.h>

using namespace std;

bool cmp(const pair<int, int> &p1, const pair<int,int> &p2) {
    return p1.second < p2.second;
}

int main(void) {

    int N;
    scanf("%d", &N);

    vector<pair<int, int>> list;
    for (int i = 0; i < N; i ++) {
        int t, s;
        scanf("%d %d", &t, &s);
        list.push_back(make_pair(t, s));
    }

    sort(list.begin(), list.end(), cmp);

    int result = list[list.size() - 1].second;
    for (int i = list.size() - 1; i >= 0; i --) {
        result = min(result, list[i].second);
        result = result - list[i].first;
    }
    if (result < 0) result = -1;
    printf("%d", result);

    exit(0);
}