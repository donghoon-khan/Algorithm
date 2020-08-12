#include <iostream>
#include <vector>
#include <queue>
#include <fstream>
#include <algorithm>
using namespace std;

int main() {
	int M, N;

	//freopen("input.txt", "r", stdin);
	scanf("%d %d", &M, &N);

	vector<vector<int>> tomatos;
	queue<pair<int, pair<int, int>>> q;
	int cnt = 0;

	for (int i = 0; i < N; i ++) {
		vector<int> v;
		for (int j = 0; j < M; j ++) {
			int tomato;
			scanf("%d", &tomato);

			if (tomato == -1) {
				cnt ++;
			}
			else if (tomato == 1) {
				q.push(make_pair(0, make_pair(i, j)));
			}
			v.push_back(tomato);
		}
		tomatos.push_back(v);
	}

	int result = 0;
	int dx[] = {1, -1, 0, 0};
	int dy[] = {0, 0, 1, -1};
	while(!q.empty()) {
		int nowDay = q.front().first;
		int nowY = q.front().second.first;
		int nowX = q.front().second.second;
		q.pop();
		result = max(result, nowDay);
		cnt ++;

		for (int i = 0; i < 4; i ++) {
			int nextY = dy[i] + nowY;
			int nextX = dx[i] + nowX;

			if (nextY >= 0 && nextX >= 0 && nextY < N && nextX < M) {
				if (tomatos[nextY][nextX] == 0) {
					q.push(make_pair(nowDay + 1, make_pair(nextY, nextX)));
					tomatos[nextY][nextX] = 1;
				}
			}
		}
	}

	if (cnt < M * N) {
		result = -1;
	}

	printf("%d\n", result);

	return 0;
}
