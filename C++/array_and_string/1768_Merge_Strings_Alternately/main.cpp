#include <iostream>
using namespace std;

class Solution
{
public:
    string mergeAlternately(string word1, string word2)
    {
        int i = 0, j = 0;
        string result;

        while (i < word1.length() && j < word2.length())
        {
            result += word1[i];
            result += word2[j];
            i++;
            j++;
        }
        while (i < word1.length())
        {
            result += word1[i];
            i++;
        }
        while (j < word2.length())
        {
            result += word2[j];
            j++;
        }

        return result;
    }
};

int main()
{
    Solution so;
    string x, y;
    // x = "abc";
    // y = "pqr";
    x = "ab";
    y = "pqrs";
    cout << so.mergeAlternately(x, y) << endl; // apbqrs
    return 0;
}