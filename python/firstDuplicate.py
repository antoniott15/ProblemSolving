def firstDuplicate(list1):
    for x in range(0,len(list1)):
        if list1[abs(list1[x])-1]<0:
            return abs(list1[x])
        else:
            list1[abs(list1[x])-1] = - list1[abs(list1[x])-1]
    return -1
n


list1 = [2,3,5,6,3,10,4,1,3]
print(firstDuplicate(list1))