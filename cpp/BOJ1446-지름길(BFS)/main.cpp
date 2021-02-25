#include <iostream>
#include <queue>
#include <vector>
#include <stdio.h>

using namespace std;

int main(void) {
    
    int N, D;
    scanf("%d %d", &N, &D);

    vector<pair<int,pair<int,int>>> map;
    for (int i = 0; i < N; i ++) {
        int src, dst, len;
        scanf("%d %d %d", &src, &dst, &len);
        if (dst > D) continue;
        map.push_back(make_pair(src, make_pair(dst, len)));
    }

    int result = D;
    queue<pair<int, int>> q;
    q.push(make_pair(0, 0));

    while(!q.empty()) {
        int now = q.front().first;
        int len = q.front().second;
        q.pop();

        if(now == D) {
            result = min(result, len);
            continue;
        }

        for (int i = 0; i < map.size(); i ++) {
            if (map[i].first < now) continue;
            int nextLen = len + (map[i].first - now) + map[i].second.second;
            q.push(make_pair(map[i].second.first, nextLen));
        }
        q.push(make_pair(D, len + (D - now)));
    }

    printf("%d", result);
    exit(0);
}