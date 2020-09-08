import math

def squares(a, b):
    counter = 0
    for elements in range (a,b+1):
        if (math.sqrt(elements)) % 1 == 0:
            counter+=1
    return counter


def squares2(a,b):
    return (int(math.floor(math.sqrt(b))))-(int(math.ceil(math.sqrt(a)))) + 1


a=2
b=8
print(squares(a,b))
print(squares2(a,b))