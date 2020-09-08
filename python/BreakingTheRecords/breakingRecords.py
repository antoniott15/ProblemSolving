
def breakingRecords(scores):
    countMin = 0
    countMax = 0
    minNum = scores[0]
    maxNum = scores[0]
    for index in range (0,len(scores)):
        if index != 0:
            print(scores[index], maxNum,minNum)
            if scores[index] > maxNum:
                maxNum = scores[index]
                countMax += 1
            elif scores[index] < minNum:
                minNum = scores[index]
                countMin += 1
    print(countMax,countMin)


a = [2,3,4,4,5,8]
breakingRecords(a)