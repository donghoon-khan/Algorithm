#include <string>
#include <vector>
#include <queue>

using namespace std;

bool cmp(string a, string b) {
    int cnt = 0;
    for (int i = 0; i < a.length(); i++) {
        if (a[i] != b[i]) {
            cnt++;
        }
        if (cnt > 1) return false;
    }
    return true;
}

int solution(string begin, string target, vector<string> words) {
    int answer = 0;

    vector<bool> visit(words.size(), false);
    queue<pair<string, int>> q;
    q.push(make_pair(begin, 0));

    while (!q.empty()) {
        pair<string, int> now = q.front();
        q.pop();

        if (now.first == target) {
            answer = now.second;
            break;
        }

        for (int i = 0; i < words.size(); i++) {
            if (visit[i]) continue;
            if (cmp(now.first, words[i])) {
                visit[i] = true;
                q.push(make_pair(words[i], now.second + 1));
            }
        }
    }
    return answer;
}