def catAndMouse(x, y, z):
    countA = 0
    countB = 0
    if x >= z:
        countA = x-z
    if x <= z:
        countA = z-x
    if y >= z:
        countB = y-z
    if y <= z:
        countB = z-y 
    if countA == countB:
        return "Mouse C"
    if countA > countB:
        return "Cat B"
    if countA < countB:
        return "Cat A"






x=1 #cata
y=3 #catb
z=2 #mouse
print(catAndMouse(x, y, z))
