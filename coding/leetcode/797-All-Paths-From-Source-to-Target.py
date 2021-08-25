class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        ret = [[0]]
        def dfs(src: int):
            for adjNode in graph[src]:
                ret[-1].append(adjNode)
                if adjNode == len(graph) - 1:
                    ret.append(ret[-1][:-1])
                    continue
                dfs(adjNode)
                del ret[-1][-1]
        dfs(0)
        return ret[:-1]
