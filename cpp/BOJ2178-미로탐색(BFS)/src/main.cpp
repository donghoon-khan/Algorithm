#include <iostream>
#include <queue>

using namespace std;

int main(void) {

	bool map[101][101];
	int N, M;
	scanf("%d %d", &N, &M);

	for (int i = 0; i < N; i ++) {
		for (int j = 0; j < M; j ++) {
			int input;
			scanf("%1d", &input);

			map[i][j] = (input == 1) ? true : false;
		}
	}

	queue<pair<int, pair<int, int>>> q;
	map[0][0] = false;
	q.push(make_pair(1, make_pair(0, 0)));

	int dx[] = {0, 0, 1, -1};
	int dy[] = {1, -1, 0, 0};
	while(!q.empty()) {
		int nowCnt = q.front().first;
		int nowY = q.front().second.first;
		int nowX = q.front().second.second;
		q.pop();

		if (nowY == N - 1 && nowX == M - 1) {
			printf("%d\n", nowCnt);
			break;
		}

		for (int i = 0; i < 4; i ++) {
			int nextY = nowY + dy[i];
			int nextX = nowX + dx[i];

			if (nextY >= 0 && nextX >= 0 && nextY < N && nextX < M) {
				if (map[nextY][nextX]) {
					map[nextY][nextX] = false;
					q.push(make_pair(nowCnt + 1, make_pair(nextY, nextX)));
				}
			}
		}
	}
}
