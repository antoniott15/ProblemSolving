def getMoneySpent(keyboards, drives, b):
    cost = 0
    arr = []
    for elements in keyboards:
        for elementos in drives:
            cost = elementos + elements
            if cost <= b:
                arr.append(cost)
    if not arr:
        return -1
    result = max(arr)
    return result




b = 5
keyboards = [4]
drives = [5]
print(getMoneySpent(keyboards, drives, b))