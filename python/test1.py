

from collections import defaultdict

def countingClasses(q):
    price = {}
    finalAnswer = {}
    d = defaultdict(list)
    for elements in range(0,len(q)):
        name, prices = q[elements]
        price[name] = price.get(name,0) + prices
        d[name].append(prices)
    for elements in price:
        d[elements].append(price[elements])

    i = 0
    for elements in d:
        for elementos in d[elements]:
            i+=1
            if i == 1:
                print("Item "+elements +":", end=" ")
            print(elementos, end=" ")
        i=0
        print()


        






a=[("banana",100),("orange",200),("banana",120),("orange",100),("banana",1200)]

countingClasses(a)