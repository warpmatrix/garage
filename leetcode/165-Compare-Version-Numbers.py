class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        arr1, arr2 = version1.split("."), version2.split(".")
        for i in range(max(len(arr1), len(arr2))):
            num1 = int(arr1[i]) if i < len(arr1) else 0
            num2 = int(arr2[i]) if i < len(arr2) else 0
            if num1 != num2:
                return 1 if num1 > num2 else -1
        return 0
