#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int minPre = prices[0];
        int result = 0;
        for (int i = 1; i < prices.size(); i++)
        {
            int curr = prices[i];
            minPre = min(curr, minPre);
            result = max(result, curr - minPre);
        }
        return result;
    }

    int min(int x, int y)
    {
        if (x < y)
        {
            return x;
        }
        return y;
    }
    int max(int x, int y)
    {
        if (x < y)
        {
            return y;
        }
        return x;
    }
};

int main()
{
    Solution obj;
    vector<int> v1 = {7, 1, 5, 3, 6, 4};
    vector<int> v2 = {7, 6, 4, 3, 1};
    cout << obj.maxProfit(v1) << endl;
    cout << obj.maxProfit(v2) << endl;
    return 0;
}