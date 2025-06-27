#include <iostream>
using namespace std;

class Solution
{
public:
    bool isSubsequence(string s, string t)
    {
        if (s.length() > t.length())
            return false;
        int countS = 0, countT = 0;

        while (countS < s.length() && countT < t.length())
        {
            if (s[countS] == t[countT])
            {
                countS++;
            }
            countT++;
        }

        return s.length() == countS;
    }
};

int main()
{
    Solution obj;
    string s1 = "abc", s2 = "ahbgdc";
    cout << obj.isSubsequence(s1, s2) << endl;
    return 0;
}