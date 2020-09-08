import math

def viralAdvertising(n):
    begining = 5
    count = 0
    for elements in range(1,n+1):
        begining = math.floor(begining/2)
        count += begining
        begining *= 3
    return count

n=3
print(viralAdvertising(n))