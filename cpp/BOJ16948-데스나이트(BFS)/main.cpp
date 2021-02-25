#include <iostream>
#include <vector>
#include <queue>
#include <stdio.h>

using namespace std;

int main(void) {

    int result = -1;
    bool map[200][200];

    int N;
    scanf("%d", &N);
    
    int r1, r2, c1, c2;
    scanf("%d %d %d %d", &r1, &c1, &r2, &c2);

    queue<pair<int, pair<int,int>>> q;
    q.push(make_pair(0,make_pair(r1, c1)));

    while(!q.empty()) {
        int nowDepth = q.front().first;
        int nowR = q.front().second.first;
        int nowC = q.front().second.second;

        if (nowR == r2 && nowC == c2) {
            result = nowDepth;
            break;
        }

        q.pop();
        
        int dc[] = {-1, 1, -2, 2, -1, 1};
        int dr[] = {-2, -2, 0, 0, 2, 2};
        for (int i = 0; i < 6; i ++) {
            int nextR = nowR + dr[i];
            int nextC = nowC + dc[i];
            if (nextR >= 0 && nextC >= 0 && nextR <= N && nextC <= N) {
                if (!map[nextR][nextC]) {
                    q.push(make_pair(nowDepth + 1, make_pair(nextR, nextC)));
                    map[nextR][nextC] = true;
                }
            }
        }
    }
    printf("%d", result);
    exit(0);
}