
def getBinary(x):
    a = '{0:b}'.format(x)
    a = a.strip()
    newArr = []
    for elements in range(0,len(a)):
        newArr.append(a[elements])
    return newArr

def reversing(n):
    n =  n[::-1]
    return  n

def binaries(n,m):
    a = getBinary(n)
    b = getBinary(m)
    counting = 0
    if len(a) == len(b):
        for x in range(0,len(a)):
            if a[x] != b[x]:
                counting+=1
    elif len(a)>len(b):
        counting = len(a) - len(b)
        b = reversing(b)
        a = reversing(a)
        for x in range(0,len(b)):
            if a[x] != b[x]:  
                counting+=1
    elif len(a)<len(b):
        counting = len(b)-len(a)
        a = reversing(a)
        b = reversing(b)
        for x in range(0,len(a)):
            if a[x] != b[x]:
                counting+=1
    return counting




print(binaries(24,35))