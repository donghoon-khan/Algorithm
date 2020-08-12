#include <iostream>
#include <vector>
#include <queue>
#include <stack>
using namespace std;

int solution(vector<vector<int>> board, vector<int> moves) {
    int answer = 0;

    vector<queue<int>> map(30);
    stack<int> st;

    for (unsigned int i = 0; i < board.size(); i ++) {
    	for (unsigned int j = 0; j < board[i].size(); j ++) {
    		if (board[i][j] != 0) {
    			map[j].push(board[i][j]);
    		}
    	}
    }
    for (unsigned int i = 0; i < moves.size(); i ++) {
    	if (map[moves[i] - 1].empty()) {
    		continue;
    	} else if (st.empty()) {
    		st.push(map[moves[i] - 1].front());
    	}
    	else if (map[moves[i] - 1].front() == st.top()) {
    		answer += 2;
    		st.pop();
    	} else {
    		st.push(map[moves[i] - 1].front());
    	}
    	map[moves[i] - 1].pop();
    }

    return answer;
}

int main() {
	cout << solution({{0,0,0,0,0},{0,0,1,0,3},{0,2,5,0,1},{4,2,4,4,2},{3,5,1,3,1}},{1,5,3,5,1,2,1,4});
    return 0;
}
