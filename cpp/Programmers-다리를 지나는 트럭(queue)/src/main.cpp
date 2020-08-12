#include <string>
#include <vector>
#include <queue>

using namespace std;

int solution(int bridge_length, int weight, vector<int> truck_weights) {
    int answer = 0;

    int sum = 0;
    queue<int> q;

    int idx = 0;
    while (idx < truck_weights.size()) {

        if (q.size() < bridge_length) {
            answer++;
            if (sum + truck_weights[idx] <= weight) {
                sum += truck_weights[idx];
                q.push(truck_weights[idx++]);
            }
            else {
                q.push(0);
            }
        }
        else {
            sum -= q.front();
            q.pop();
        }
    }

    return answer + bridge_length;
}

int main(void) {
    vector<int> param{ 10,10,10,10,10,10,10,10,10,10 };
    printf("%d ", solution(100, 100, param));
}