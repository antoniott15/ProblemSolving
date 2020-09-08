def countingValleys(n, s):
    count = 0
    upOrDown = False
    for elements in s:    
        if elements == 'U':
            if  not upOrDown:
                count+=1
                upOrDown = True
        if elements == 'D':
            if upOrDown:
                count-=1
                upOrDown = False
    return count









n = 8
s = "DDUUDDUDUUUD"
print(countingValleys(n, s))
