#include <iostream>
#include <queue>
#include <vector>
#include <memory.h>
#include <algorithm>
#include <stdio.h>

using namespace std;

typedef struct position{
    int wall;
    int y;
    int x;
}position;

int main(void) {
    
    int N, M;
    scanf("%d %d", &N, &M);

    int map[101][101];
    for (int i = 1; i <= M; i ++) {
        for (int j = 1; j <= N; j ++) {
            scanf("%1d", &map[i][j]);
        }
    }

    queue<position> q;
    q.push(position{0, 1, 1});
    int visit[101][101];
    int dx[] = {0, 0, 1, -1};
    int dy[] = {1, -1, 0, 0};
    int result = N * M;

    for (int i = 0; i < 101; i ++) {
        for (int j = 0; j < 101; j ++) {
            visit[i][j] = N*M;
        }
    }

    while(!q.empty()) {
        int x = q.front().x;
        int y = q.front().y;
        int wall = q.front().wall;
        if (x == N && y == M) {
            result = min(result, wall);
        }
        q.pop();

        for (int i = 0; i < 4; i ++) {
            int nextX = x + dx[i];
            int nextY = y + dy[i];
            if (nextX >= 1 && nextY >= 1 && nextX <= N && nextY <= M) {
                
                if (map[nextY][nextX] == 1 && visit[nextY][nextX] > wall + 1) {
                    visit[nextY][nextX] = wall + 1;
                    q.push(position{wall + 1, nextY, nextX});
                } else if (map[nextY][nextX] == 0 && visit[nextY][nextX] > wall) {
                    visit[nextY][nextX] = wall;
                    q.push(position{wall, nextY, nextX});
                }
            }
        }
    }
    printf("%d\n", result);
    exit(0);
}