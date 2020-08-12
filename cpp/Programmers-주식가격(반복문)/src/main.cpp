#include <string>
#include <vector>
#include <queue>

using namespace std;

vector<int> solution(vector<int> prices) {
    vector<int> answer(prices.size(), 0);

    for (int i = 0; i < prices.size() - 1; i++) {
        for (int j = i + 1; j < prices.size(); j++) {
            if (prices[i] > prices[j]) {
                answer[i] = j - i;
                break;
            }
        }
        if (answer[i] == 0) {
            answer[i] = prices.size() - 1 - i;
        }
    }

    return answer;
}

int main(void) {
    vector<int> prices{ 2, 3, 4, 5, 1 };
    vector<int> ret;
    ret = solution(prices);

    for (int i = 0; i < ret.size(); i++) {
        printf("%d ", ret[i]);
    }


}