#include <string>
#include <vector>

using namespace std;

int solution(string arrangement) {
    int answer = 0;

    int open = 0;
    for (int i = 0; i < arrangement.length(); i++) {
        if (arrangement[i] == '(') {
            open++;
        }
        else {
            open--;
            if (arrangement[i - 1] == '(') {
                answer += open;
            }
            else {
                answer++;
            }
        }
    }


    return answer;
}

int main(void) {
    vector<int> priorities{ 2, 1, 3, 2 };

    printf("%d", solution("()(((()())(())()))(())"));

}