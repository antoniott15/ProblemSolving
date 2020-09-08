import collections

def cutTheSticks(arr):
    result = []
    l = len(arr)
    print(collections.Counter(arr).items())
    for k,v in sorted(collections.Counter(arr).items()):
        print(k,v)
        result.append(l)
        l-=v
    return result
        


a = [2,3,4,4,5,8]
print(cutTheSticks(a))