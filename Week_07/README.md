# 学习笔记

## 分析单词搜索2时间复杂度

设二维网格大小为m*n, 单词数量为k, 平均长度为i
### 从单词出发，使用DFS
遍历每个单词，然后从网格位置0开始进行dfs搜索，刚开始我们有四个方向可以选择，到下一个位置时只有3个（不包括来时的方向），时间复杂度为m\*n\*4\*3^(i-1)\*k, 
优化一下，我们可以先扫描一遍二维网格，记录每个字母出现的位置，然后从单词出发时就可以直接定位到网格开始的位置，假设每个字母出现的次数平均为p,
那么有p\*4\*3^(i-1)\*k


### 从网格出发，使用DFS + trie树
遍历网格时间复杂为m\*n, 对每个字母进行深度优先搜索时间复杂度为4\*3^(i-1)，总的时间复杂度为m\*n\*4\*3^(i-1)。少了单词数量因子的影响。

假如单词数量较少的话，直接从从单词出发，使用DFS会比较好。

## 双向BFS代码模版

```go
unc bbfs(start int, end int) int {
    level := 0
    visited := map[int]bool{}
    beginSet := map[int]bool{}
    endSet := map[int]bool{}
    beginSet[start] = true
    endSet[end] = true
    visited[start] = true
    visited[end] = true
    
    for len(beginSet) > 0 {
        if len(beginSet) > len(endSet) {
            beginSet, endSet = endSet, beginSet
        }
        
        newBeginSet := map[int]bool{}
        for v := range beginSet {
            // if endSet[v] {
            //         return level+1
            //     } 不能放在这里， 因为v此时已经是visited过的了，而已visited过的v不会再被加入
            for _, nextV := range getEdges(v) {
                if endSet[nextV] {
                    return level + 1 //表示走过的边数
                }
                if !visited[nextV] {
                    newBeginSet[nextV] = true
                    visited[nextV] = true
                }
            }
        }
        beginSet = newBeginSet
        level++
    }
    return 0
}
```