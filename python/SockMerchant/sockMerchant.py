from collections import Counter

def sockMerchant(n, ar):
    count = 0
    for index,values in Counter(ar).items():
        a =  int(values/2) 
        count+=a

    return count


n=9
ar=[10,20,20,10,10,30,50,10,20]
print(sockMerchant(n,ar))
