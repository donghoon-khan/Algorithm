#include <iostream>
#include <vector>
#include <stdio.h>

using namespace std;

vector<string> map;
bool visit[5][5];

int R, C, K;
int dx[] = {1, -1, 0, 0};
int dy[] = {0, 0, 1, -1};

int solve(int y, int x, int depth) {

    if (depth == K) {
        return (y == 0 && x == C-1) ? 1 : 0;
    }
    
    int ret = 0;
    for (int i = 0; i < 4; i ++) {
        int nextX = x + dx[i];
        int nextY = y + dy[i];
        if (nextX >= 0 && nextY >= 0 && nextY < R && nextX < C) {
            if (map[nextY][nextX] == '.' && !visit[nextY][nextX]) {
                visit[nextY][nextX] = true;
                ret += solve(nextY, nextX, depth + 1);
                visit[nextY][nextX] = false;
            }
        }
    }
    return ret;
}

int main(void){
    scanf("%d %d %d", &R, &C, &K);

    for (int i = 0; i < R; i ++) {
        string str;
        cin >> str;
        map.push_back(str);
    }

    visit[R-1][0] = true;
    printf("%d", solve(R-1, 0, 1));
    exit(0);
}