#include <iostream>
#include <algorithm>

using namespace std;

int main(void) {
	int N, M;

	scanf("%d %d", &N, &M);

	int candies[1001][1001], dp[1001][1001];
	for (int i = 0; i < N; i ++) {
		for (int j = 0; j < M; j ++) {
			scanf("%d", &candies[i][j]);
			dp[i][j] = candies[i][j];
		}
	}

	for (int i = 0; i < N; i ++) {
		for (int j = 0; j < M; j ++) {
			int dr[] = {1, 0, 1};
			int dc[] = {0, 1, 1};

			for (int pos = 0; pos < 3; pos ++) {
				int nr = i + dr[pos];
				int nc = j + dc[pos];
				if (nr < N && nc < M) {
					dp[nr][nc] = max(candies[nr][nc] + dp[i][j], dp[nr][nc]);
				}
			}
		}
	}

	printf("%d\n", dp[N - 1][M - 1]);
	return 0;
}
