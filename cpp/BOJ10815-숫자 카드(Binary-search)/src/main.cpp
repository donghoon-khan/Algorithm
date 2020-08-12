#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(void) {

	int N;
	scanf("%d", &N);

	vector<int> card;
	for (int i = 0; i < N; i ++) {
		int num;
		scanf("%d", &num);
		card.push_back(num);
	}

	sort(card.begin(), card.end());

	int M;
	scanf("%d", &M);
	while (M --) {
		int num;
		scanf("%d", &num);

		if (binary_search(card.begin(), card.end(), num)) {
			printf("1 ");
		} else {
			printf("0 ");
		}
	}

}
