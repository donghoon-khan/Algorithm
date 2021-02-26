#include <iostream>
#include <queue>
#include <string>
#include <algorithm>
#include <vector>
#include <stdio.h>

using namespace std;

#define MAX_FLOOR 1000001

int main(void) {
    
    int F, S, G, U, D;
    scanf("%d %d %d %d %d", &F, &S, &G, &U, &D);

    int visit[MAX_FLOOR];
    for (int i = 0; i < MAX_FLOOR; i ++) {
        visit[i] = MAX_FLOOR;
    }
    queue<int> q;
    q.push(S);
    visit[S] = 0;

    while(!q.empty()) {
        int nowPos = q.front();
        q.pop() ;

        if(nowPos == G) {
            break;
        }

        for (int i = 0; i < 2; i ++) {
            int nextPos = 0;
            if (i == 0) nextPos = nowPos + U;
            else nextPos = nowPos - D;

            if (nextPos >= 1 && nextPos <= F && visit[nextPos] > visit[nowPos] + 1) {
                q.push(nextPos);
                visit[nextPos] = visit[nowPos] + 1;
            }
        }
    }

    string result = (visit[G] == MAX_FLOOR) ? "use the stairs" : to_string(visit[G]);
    cout << result << endl;

    exit(0);
}