// "한조서열정리하고옴ㅋㅋ CPP"

#include <iostream>
#include <algorithm>

using namespace std;

int main()
{
    int n, temp;
    cin >> n;

    int high = 0, count = 0, ans = 0;
    for (int i = 0; i < n; i ++) {
        cin >> temp;
        if (high > temp) {
            count ++;
        }
        else {
            ans = max(ans, count);
            count = 0;
            high = temp;
        }
        ans = max(ans, count);
    }

    cout << ans;

    return 0;
}
