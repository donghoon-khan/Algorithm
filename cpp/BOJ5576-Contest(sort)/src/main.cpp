#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

int main(void) {

	//freopen("input.txt", "r", stdin);

	vector<int> W;
	for (int i = 0; i < 10; i ++) {
		int num;
		scanf("%d", &num);
		W.push_back(num);
	}

	vector<int> K;
	for (int i = 0; i < 10; i ++) {
		int num;
		scanf("%d", &num);
		K.push_back(num);
	}

	sort(W.begin(), W.end());
	sort(K.begin(), K.end());

	printf("%d ", W[W.size() - 1] + W[W.size() - 2] + W[W.size() - 3]);
	printf("%d ", K[K.size() - 1] + K[K.size() - 2] + K[K.size() - 3]);

}
