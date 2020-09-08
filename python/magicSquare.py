def formingMagicSquare(s):
    horizontal = []
    vertical = []
    for index,elements in enumerate(s):
        horizontal.append(sum(elements)) 
        if s[index] == [elements  for elements in range(0,3)]:
            vertical.append(s[index])
    print(vertical)



c = [[4, 9, 2], [3, 5, 7], [8, 1, 5]]
formingMagicSquare(c)