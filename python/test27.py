def reverse(x: int) -> int: 
    if  x < -2**31:
        numberToReverse = str(x)
    if numberToReverse[0] == '-':
        temp = numberToReverse[1:]
        newNumberToReverse =  numberToReverse[0] + temp[::-1]
    else:
        return int(numberToReverse[::-1])
    return int(newNumberToReverse)


print(reverse(15469))