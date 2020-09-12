#include <iostream>
#include <vector>

using namespace std;

int result = 1000;

void back_tracking(vector<int> stats, vector<int> team, int depth) {
    if (depth == stats.size()) {
        if (result > team.size()) {
            result = team.size();
        }
        return;
    }

    for (int i = depth; i < stats.size(); i ++) {
        for (int j = 0; j < team.size(); j ++) {
            if (team[j] < stats[i]) {
                team[j] = stats[i];
                back_tracking(stats, team, depth + 1);
                stats[i] = team[j];
            }
        }
        team.push_back(stats[i]);
        back_tracking(stats, team, depth + 1);
    }
}

int solution(vector<int> stats) {
    vector<int> team;
    back_tracking(stats, team, 0);
    return result;
}

int main()
{
    vector<int> stats{5, 3, 4, 6, 2, 1};
    std::cout << solution(stats) << std::endl;
    return 0;
}