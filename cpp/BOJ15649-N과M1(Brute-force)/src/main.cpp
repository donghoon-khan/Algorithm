#include <iostream>
#include <vector>
using namespace std;

vector<int> v;
bool visit[9];

void solve(int M, int N)
{
	if (v.size() == M){
		for (int i = 0; i < v.size(); i ++) {
			printf("%d ", v[i]);
		}
		printf("\n");
		return;
	}

	for (int i = 1; i <= N; i ++) {
		if (!visit[i]) {
			visit[i] = true;
			v.push_back(i);
			solve(M, N);
			v.pop_back();
			visit[i] = false;
		}
	}
}

int main() {
	int N, M;

	scanf("%d %d", &N, &M);

	solve(M, N);

	return 0;
}
