#include <string>
#include <vector>
#include <map>
#include <iostream>

using namespace std;

int solution(string skill, vector<string> skill_trees) {
    int answer = 0;

    map<char, int> m;
    for (unsigned int i = 0; i < skill.length(); i ++) {
    	m.insert(make_pair(skill[i], i + 1));
    }

	for (unsigned int i = 0; i < skill_trees.size(); i ++) {
		int cnt = 1;
		bool isValid = false;
		for (unsigned int j = 0; j < skill_trees[i].length(); j ++) {
			if (m[skill_trees[i][j]] > cnt) {
				isValid = true;
				break;
			} else if (m[skill_trees[i][j]] == cnt) {
				cnt ++;
			}
		}

		if (!isValid) {
			answer ++;
		}
	}
    return answer;
}
int main() {

	cout << solution("CBD", {"BACDE", "CBADF", "AECB", "BDA"});

    return 0;

}
