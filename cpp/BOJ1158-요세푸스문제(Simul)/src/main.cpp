#include <iostream>
#include <vector>

using namespace std;

int main(void) {
	int N, K;
	scanf("%d %d", &N, &K);

	vector<int> jo;
	for (int i = 1; i <= N; i ++) {
		jo.push_back(i);
	}

	printf("<");
	int pos = 0;
	while(!jo.empty()) {
		pos = (pos + K - 1) % jo.size();

		if (jo.size() > 1)
			printf("%d, ", jo[pos]);
		else
			printf("%d>", jo[pos]);
		jo.erase(jo.begin() + pos);
	}

	return 0;
}
