#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

int main() {
	int N, M;
	scanf("%d %d", &N, &M);

	string str;
	for (int i = 0; i < N; i ++) {
		cin >> str;
		reverse(str.begin(), str.end());
		cout << str << endl;
	}

	return 0;
}
