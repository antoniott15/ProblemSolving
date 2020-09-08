def sumOfTwoArrays(list1,list2,value):
    result = {}
    for i in range(0,len(list1)):   
        result[value - list1[i]] = i
    for value in list2:
        if result.__contains__(value):
            return True
    return False





list1 = [2,3,4,1,5,-2,3,10]
list2 = [2,34,12,45,2,3,5,7,9]
value = 100

print(sumOfTwoArrays(list1,list2,value))