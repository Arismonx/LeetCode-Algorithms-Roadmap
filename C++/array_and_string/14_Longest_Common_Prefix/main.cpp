#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution
{
public:
    string longestCommonPrefix(vector<string> &strs)
    {
        string pref = strs[0];
        int prefLen = pref.length();
        for (int i = 1; i < strs.size(); i++)
        {
            string currPref = strs[i];
            while (prefLen > currPref.length() || pref != currPref.substr(0, prefLen))
            {
                prefLen--;
                if (prefLen == 0)
                {
                    return "";
                }
                pref = pref.substr(0, prefLen);
            }
        }
        return pref;
    }
};

int main()
{
    Solution obj;
    vector<string> str1 = {"flower", "flow", "flight"};
    cout << obj.longestCommonPrefix(str1) << endl;
    cout << "test\n";
    return 0;
}