def getNumber(arr, left,right,x):
    if x > max(arr) or x < min(arr):
        return -1
    if right < left: 
        return -1
    mid = left + (right - left) // 2
    if arr[mid] == x: 
        return mid 
    elif arr[mid] > x:
        return getNumber(arr, left, mid-1, x) 
    else: 
        return getNumber(arr, mid + 1, right, x) 



arr = [ 2, 3, 4, 10, 40 ] 
x = 50


print(getNumber(arr,0,len(arr),x))