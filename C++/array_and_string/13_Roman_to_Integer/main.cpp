#include <iostream>
#include <map>
using namespace std;

class Solution
{
public:
    int romanToInt(string s)
    {
        map<char, int> m{
            {'I', 1},
            {'V', 5},
            {'X', 10},
            {'L', 50},
            {'C', 100},
            {'D', 500},
            {'M', 1000},
        };
        int result = 0;
        for (int i = 0; i < s.length(); i++)
        {
            if (i + 1 < s.length() && m[s[i]] < m[s[i + 1]])
            {
                result -= m[s[i]];
            }
            else
            {
                result += m[s[i]];
            }
        }
        return result;
    }
};

int main()
{
    Solution obj;
    string str1 = {"III"};
    string str2 = {"LVIII"};
    string str3 = {"MCMXCIV"};
    cout << obj.romanToInt(str1) << endl;
    cout << obj.romanToInt(str2) << endl;
    cout << obj.romanToInt(str3) << endl;
    return 0;
}