class Solution:
    def unhappyFriends(self, n: int, preferences: List[List[int]], pairs: List[List[int]]) -> int:
        preferIdxs = [ [0] * n for _ in range(n) ]
        for i, friends in enumerate(preferences):
            for idx, friend in enumerate(friends):
                preferIdxs[i][friend] = idx
                
        matchs = [0] * n
        for _, pair in enumerate(pairs):
            matchs[pair[0]] = pair[1]
            matchs[pair[1]] = pair[0]

        ret = 0
        for i, match in enumerate(matchs):
            for idx, friend in enumerate(preferences[i]):
                if match == friend:
                    break
                if preferIdxs[friend][i] < preferIdxs[friend][matchs[friend]]:
                    ret += 1
                    break
        return ret
