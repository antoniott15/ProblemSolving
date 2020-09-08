from collections import Counter
def migratoryBirds(arr):
    sorted(arr)
    d = Counter(arr).most_common(1)
    return d[0][0]





arr = [1,2,3,4,5,4,3,2,1,3,4]
print(migratoryBirds(arr))