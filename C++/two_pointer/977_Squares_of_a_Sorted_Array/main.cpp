#include <iostream>
#include <vector>
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
    vector<int> sortedSquares(vector<int> &nums)
    {
        vector<int> result;
        // int left = 0;
        // int right = nums.size() - 1;

        for (int i = 0; i < nums.size(); i++)
        {
            int currNum = ads_and_multiply(nums[i]);
            result.push_back(currNum);
        }

        sort(result.begin(), result.end());

        return result;
    }

    int ads_and_multiply(int x)
    {
        if (x < 0)
        {
            return -x * -x;
        }
        return x * x;
    }
};

int main()
{
    Solution sol;
    vector<int> nums = {-4, -1, 0, 3, 10};
    sol.sortedSquares(nums);
    return 0;
}