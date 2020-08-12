#include <string>
#include <vector>

using namespace std;

vector<int> solution(vector<int> heights) {
    vector<int> answer(heights.size(), 0);

    for (int i = heights.size() - 1; i > 0; i--) {
        for (int j = i - 1; j >= 0; j--) {
            if (heights[i] < heights[j]) {
                answer[i] = j + 1;
                break;
            }
        }
    }

    return answer;
}

int main(void) {
    vector<int> heights{ 6,9,5,7,4 };

    vector<int> answer = solution(heights);

    for (int i = 0; i < answer.size(); i++) {
        printf("%d ", answer[i]);
    }
}