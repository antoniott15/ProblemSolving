def pageCount(n,p):
    items = []
    for elements in range(0,n+1,2):
        items.insert(elements,{elements,elements+1})
    count = 0
    ascending = 0
    descending = 0
    for elements,values in items:
        count+=1
        if elements == p or values == p:
            ascending = count
            break
    count = 0
    for elements,values in reversed(items):
        count+=1
        if elements == p or values ==p:
            descending = count
            break;
    resultArr = [ascending,descending]
    print(ascending,descending)
    result = min(resultArr) -1
    return result




n=5
p=4
print(pageCount(n,p))