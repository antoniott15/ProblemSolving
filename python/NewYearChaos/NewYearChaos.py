def minimumBribes(q):
    newArr = q
    countingBribes = 0
    diference = 0
    newArr = sorted(newArr)
    for elements in range(0,len(q)):
        if q[elements] != newArr[elements]:
            diference = q[elements] - newArr[elements]
            if diference <= 1:
                if q[elements] > newArr[elements]:
                    countingBribes += 1
            elif diference == 2:
                    countingBribes += 2
            else:
                print("Too chaotic")
                return
    print(countingBribes)      


a = [7,1,3,4,1,7]

minimumBribes(a)