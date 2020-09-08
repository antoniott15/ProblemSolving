def biggerIsGreater(arr):
    arr = list(arr)
    i = len(arr) - 1
    while i > 0 and arr[i - 1] >= arr[i]:
    	i -= 1
    if i <= 0:
    	return "not answer"
	
    j = len(arr) - 1
    while arr[j] <= arr[i - 1]:
    	j -= 1
    arr[i - 1], arr[j] = arr[j], arr[i - 1]
	
    arr[i : ] = arr[len(arr) - 1 : i - 1 : -1]
    result = ""
    for elements in arr:
        result+=elements
    return result

        


w ="hcdk" 

print(biggerIsGreater(w))