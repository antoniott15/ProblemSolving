def beautifulDays(i, j, k):
    count = 0
    for elements in range(i,j+1):
        if (elements-int(str(elements)[::-1]))%k == 0:
            count+=1
    return count



i = 13
j = 45
k = 3
print(beautifulDays(i, j, k))