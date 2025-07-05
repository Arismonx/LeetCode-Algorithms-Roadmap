#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution
{
public:
    vector<string> summaryRanges(vector<int> &nums)
    {
        int first = 0;
        vector<string> result;
        for (int i = 1; i < nums.size(); i++)
        {
            if (nums[i] != nums[i - 1] + 1)
            {
                if (first == i - 1)
                {
                    result.push_back(to_string(nums[first]));
                }
                else
                {
                    result.push_back(to_string(nums[first]) + "->" + to_string(nums[i - 1]));
                }
                first = i;
            }
        }

        if (first == nums.size() - 1)
        {
            result.push_back(to_string(nums[first]));
        }
        else
        {
            result.push_back(to_string(nums[first]) + "->" + to_string(nums.back()));
        }
        return result;
    }
};

int main()
{
    Solution obj;
    vector<int> nums = {0, 1, 2, 4, 5, 7};
    obj.summaryRanges(nums);

    return 0;
}