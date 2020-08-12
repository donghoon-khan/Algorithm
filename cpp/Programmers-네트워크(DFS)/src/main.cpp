#include <iostream>
#include <vector>
#include <string>
using namespace std;

bool dfs(vector<vector<int>>& computers, int n) {
    if (!computers[n][n])    return false;
    computers[n][n] = 0;
    for (int i = 0; i < computers.size(); i++) {
        if (computers[n][i])    dfs(computers, i);
    }
    return true;
}

int solution(int n, vector<vector<int>> computers) {
    int answer = 0;
    for (int i = 0; i < n; i++) {
        if (computers[i][i] && dfs(computers, i))
            answer++;
    }
    return answer;
}

void print(int n, vector<vector<int>> computers, int answer){
    int t = solution(n, computers);
    if (t == answer)    cout << "정답" << endl;
    else    cout << "틀림" << endl;
}

int main() {
    print(3, { {1,1,0}, {1,1,0}, {0,0,1} }, 2);
    print(3, { {1,1,0}, {1,1,1}, {0,1,1} }, 1);
    return 0;
}
