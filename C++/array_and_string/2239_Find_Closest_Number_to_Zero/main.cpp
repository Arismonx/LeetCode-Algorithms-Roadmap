#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
    int findClosestNumber(vector<int> &nums)
    {
        int closeZero = nums[0];
        int minAbs = abs(closeZero);

        for (int i = 1; i < nums.size(); i++)
        {
            int currNum = nums[i];
            int currAbs = abs(currNum);

            if (currAbs < minAbs)
            {
                minAbs = currAbs;
                closeZero = currNum;
            }
            else if (currAbs == minAbs && currNum > closeZero)
            {
                closeZero = currNum;
            }
        }

        return closeZero;
    }

    int abs(int x)
    {
        if (x < 0)
        {
            return -x;
        }
        else
        {
            return x;
        }
    }
};

int main()
{
    Solution obj;
    vector<int> v = {-4, -2, 1, 4, 8};
    vector<int> v1 = {2, -1, 1};
    cout << "value close zero: " << obj.findClosestNumber(v) << endl;
    cout << "value close zero: " << obj.findClosestNumber(v1) << endl;
    return 0;
}