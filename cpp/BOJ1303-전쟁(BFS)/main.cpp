#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>
#include <stdio.h>

using namespace std;

int N, M;
bool visit[101][101];
vector<string> map;

int BFS(int n, int m) {
    queue<pair<int, int>> q;
    q.push( make_pair(n, m));
    visit[n][m] = true;
    int cnt = 0;
    while(!q.empty()) {
        cnt ++;
        int nowY = q.front().first;
        int nowX = q.front().second;
        q.pop();

        int dx[] = {0, 0, 1, -1};
        int dy[] = {1, -1, 0, 0};
        for (int i = 0; i < 4; i ++) {
            int nextY = nowY + dy[i];
            int nextX = nowX + dx[i];
            if (nextY >= 0 && nextY < N && nextX >= 0 && nextX < M) {
                if (!visit[nextY][nextX] && (map[nowY][nowX] == map[nextY][nextX])) {
                    q.push(make_pair(nextY, nextX));
                    visit[nextY][nextX] = true;
                }
            }
        }
    }
    return cnt * cnt;
}

int main(void) {
    cin >> M >> N;
    
    for (int i = 0; i < N;i ++) {
        string str;
        cin >> str;
        map.push_back(str);
    }

    int w = 0, b = 0;
    for (int i = 0; i < map.size(); i ++) {
        for (int j = 0; j < map[i].length(); j ++) {
            if (!visit[i][j]) {
                if (map[i][j] == 'W') {
                    w += BFS(i, j);
                } else {
                    b += BFS(i, j);
                }
            }
        }
    }

    cout << w << " " << b << endl;
    exit(0);
}