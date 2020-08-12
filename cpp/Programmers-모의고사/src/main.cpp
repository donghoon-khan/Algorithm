#include <string>
#include <vector>
#include <iostream>
using namespace std;

vector<int> solution(vector<int> answers) {
    vector<int> answer;

    const int num[3][10] = {
    		{1, 2, 3, 4, 5},
			{2, 1, 2, 3, 2, 4, 2, 5},
			{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
    };

    vector<int> user(3, 0);
    for (int i = 0; i < answers.size(); i ++) {
    	if (answers[i] == num[0][i % 5]) {
    		user[0] ++;
    	}
    	if (answers[i] == num[1][i % 8]) {
    		user[1] ++;
    	}
    	if (answers[i] == num[2][i % 10]) {
    		user[2] ++;
    	}
    }

    int maxValue = max(user[0], max(user[1], user[2]));
    for (int i = 0; i < 3; i ++) {
    	if (user[i] == maxValue) {
    		answer.push_back(i + 1);
    	}
    }
    return answer;
}

int main() {

	solution({1,2,3,4,5, 6, 7, 8, 9, 10});
    return 0;

}
