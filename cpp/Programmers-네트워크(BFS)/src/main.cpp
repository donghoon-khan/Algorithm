#include <string>
#include <vector>
#include <queue>

using namespace std;

int solution(int n, vector<vector<int>> computers) {
    int answer = 0;

    queue<int> q;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (computers[i][j] == 1) {
                computers[i][j] = computers[j][i] = 0;
                q.push(j);
            }
        }

        if (!q.empty()) {
            answer++;
            while (!q.empty()) {
                int now = q.front();
                computers[now][now] = 0;
                q.pop();
                for (int j = 0; j < n; j++) {
                    if (now == j) continue;
                    else if (computers[now][j] == 1) {
                        computers[now][j] = computers[j][now] = 0;
                        q.push(j);
                    }
                }
            }
        }

    }
    return answer;
}

int main(void) {
    vector<vector<int>> map{ {1, 1, 0},{1, 1, 0},{0, 0, 1} };
    printf("%d\n", solution(3, map));
    //printf("%d", solution(numbers, 3));
}