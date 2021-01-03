// 对于堆的问题，还是尽量使用C++现成的库比较好，go这点太不好了
// 大顶堆，小顶堆的写法学习一下
class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        priority_queue<int, vector<int>, less<int>> q;
        for (int s: stones) {
            q.push(s);
        }
        while (q.size() > 1) {
            int a = q.top();
            q.pop();
            int b = q.top();
            q.pop();
            if (a > b) {
                q.push(a-b);
            }
            
        }
        return q.empty()? 0 : q.top();
    }
};