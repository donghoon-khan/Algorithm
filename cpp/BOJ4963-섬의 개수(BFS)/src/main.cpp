#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
#include <algorithm>

using namespace std;

int main(void) {

	//freopen("input.txt", "r", stdin);

	while(1) {
		int w, h;
		scanf("%d %d", &w, &h);
		if (w == 0 && h == 0) break;

		vector<vector<int>> map;

		for (int i = 0; i < h; i ++) {
			vector<int> v;
			for (int j = 0; j < w; j ++) {
				int num;
				scanf("%d", &num);
				v.push_back(num);
			}
			map.push_back(v);
		}

		int result = 0;
		int dx[] = {1, -1, 0, 0, -1, -1, 1, 1};
		int dy[] = {0, 0, 1, -1, -1, 1, -1, 1};
		queue<pair<int, int>> q;
		for (int i = 0; i < h; i ++) {
			for (int j = 0; j < w; j ++) {
				if (map[i][j] == 1) {
					result ++;
					q.push(make_pair(i, j));
					map[i][j] = 0;
					while(!q.empty()) {
						int nowY = q.front().first;
						int nowX = q.front().second;
						q.pop();

						for (int pos = 0; pos < 8; pos ++) {
							int nextY = nowY + dy[pos];
							int nextX = nowX + dx[pos];
							if (nextX >= 0 && nextY >= 0 && nextY < h && nextX < w) {
								if (map[nextY][nextX] == 1) {
									q.push(make_pair(nextY, nextX));
									map[nextY][nextX] = 0;
								}
							}
						}
					}
				}
			}
		}

		printf("%d\n", result);
	}
}
