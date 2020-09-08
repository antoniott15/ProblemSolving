import itertools 
def birthday(s, d, m):
    count = 0
    for x in range(0,len(s)):
        if sum(s[x:x+m]) == d and len(s[x:x+m]) == m:
            count+=1
    return count


s = [1,2,1,3,2]
d = 3
m = 2

print(birthday(s,d,m))