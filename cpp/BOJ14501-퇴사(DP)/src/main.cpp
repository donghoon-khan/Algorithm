#include <iostream>
#include <vector>

using namespace std;

int main() {
	int n;
	int tempDay, tempCost;
	int temp;
	cin >> n;
	vector<int> cost(n + 1, 0);
	vector<int> day(n + 1, 0);
	vector<int> dp(n + 1, 0);
	for (int i = 0; i < n; i++) {
		cin >> tempDay >> tempCost;
		cost[i] = tempCost;
		day[i] = tempDay;
	}
	for (int i = 0; i < n; i++) {
		if (i + day[i] > n)
			continue;
		if (dp[i + day[i]] < dp[i] + cost[i]) {
			dp[i + day[i]] = dp[i] + cost[i];
			for (int j = i + day[i]; j <= n; j++) {
				if (dp[j] < dp[i + day[i]]) {
					dp[j] = dp[i + day[i]];
				}
			}
		}
	}
	cout << dp[n];
	return 0;
}
