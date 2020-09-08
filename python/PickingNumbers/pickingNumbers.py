from collections import Counter
def pickingNumbers(a):
    newArr = []
    moreRepeated = Counter(a).most_common(1)
    for elements in a:
        if elements == moreRepeated[0][0] +1 or elements == moreRepeated[0][0]-1:
            newArr.append(elements)
    nextOne = Counter(newArr).most_common(1)
    return(nextOne[0][1]+moreRepeated[0][1])




    




arr = [4,2,3,4,4,9,98,98,3,3,3,4,2,98,1,98,98,1,1,4,98,2,98,3,9,9,3,1,4,1,98,9,9,2,9,4,2,2,9,98,4,98,1,3,4,9,1,98,98,4,2,3,98,98,1,99,9,98,98,3,98,98,4,98,2,98,4,2,1,1,9,2,4]
print(pickingNumbers(arr))