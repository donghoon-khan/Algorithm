#include <string>
#include <vector>
#include <queue>

using namespace std;

vector<int> solution(vector<int> progresses, vector<int> speeds) {
    vector<int> answer;

    queue<int> q;
    int deploy_day = (100 - progresses[0]) / speeds[0];
    q.push(0);
    for (int i = 1; i < progresses.size(); i++) {
        int day = (100 - progresses[i]) / speeds[i];
        if (day > deploy_day) {
            deploy_day = day;
            answer.push_back(q.size());
            while (!q.empty()) {
                q.pop();
            }
        }
        q.push(i);
    }

    if (q.size() > 0) {
        answer.push_back(q.size());
    }

    return answer;
}

int main(void) {
    vector<int> progresses{ 93,30,55 };
    vector<int> speeds{ 1,30,5 };
    vector<int> ret;

    ret = solution(progresses, speeds);

    for (int i = 0; i < ret.size(); i++) {
        printf("%d ", ret[i]);
    }

}