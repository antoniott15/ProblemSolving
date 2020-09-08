def divisibleSumPairs(n, k, ar):
    count = 0
    for value in range(0,len(ar)):
        for elements in range(0,len(ar)):
            if value != elements:
                if elements>value:
                    if (ar[value] + ar[elements]) % k==0:
                        count+=1
    return count




k = 3
ar = [1,3,2,6,1,2]
d = 6
print(divisibleSumPairs(d,k,ar))