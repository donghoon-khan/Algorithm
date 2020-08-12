#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(void) {

	int n;
	scanf("%d", &n);

	vector<int> array;
	for (int i = 0; i < n; i++) {
		int num;
		scanf("%d",&num);
		array.push_back(num);
	}

	vector<int> dp(n, -1);

	dp[0] = 0;
	for (int i = 0; i < n; i ++) {
		if (dp[i] == -1) continue;
		for (int j = i + array[i]; j > i; j --) {
			if (j >= n) continue;
			if (dp[j] == -1) {
				dp[j] = dp[i] + 1;
			}
			else {
				dp[j] = min(dp[j], dp[i] + 1);
			}
		}
	}

	printf("%d", dp[n - 1]);
}
