#include <vector>

using namespace std;

class Solution
{
public:
  vector<int> plusOne(vector<int> &digits)
  {
    auto res = digits;
    auto i = res.end() - 1;
    auto r = 1;
    while (true)
    {
      if (r == 0)
        break;
      auto cur_value = *i;
      cur_value = cur_value + r;
      r = cur_value > 9 ? 1 : 0;
      *i = (cur_value % 10);
      if (!r)
      {
        break;
      }
      if (i == res.begin() && r)
      {
        res.insert(res.begin(), 1);
        return res;
      }
      i--;
    }
    return res;
  }
};