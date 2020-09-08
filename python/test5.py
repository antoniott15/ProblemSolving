def minimumBribes(q):
    newArr = q
    countingBribes = 0
    diference = 0
    newArr = sorted(newArr)
    for elements in range(0,len(q)):
        print(q[elements], newArr[elements])
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


a = [1,2,5,3,7,8,6,4]

minimumBribes(a)