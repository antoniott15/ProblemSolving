


def getTotalX(a, b):
    minA = max(a)
    minB = max(b)
    arr = [minA,minB]
    total = max(arr)
    count = 0
    newArr = []
    countTotal = 0
    for elements in a:
        newArr.append(elements)
    for elements in b:
        newArr.append(elements)
    for elements in range (1,total):
        for elementos in newArr:
            if elementos % elements == 0:
                count+=1
        if count == len(newArr):
            countTotal+=1
        count=0
    return countTotal

        


a = [3,9,6]
b = [36,72]
print(getTotalX(a, b))