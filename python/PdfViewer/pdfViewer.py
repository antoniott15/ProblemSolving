def designerPdfViewer(h, word):
    newArr = []
    for index,elementos in enumerate(h):
        for elements in word:
            if index == ord(elements)-97:
                newArr.append(elementos)
    maxHei = max(newArr)
    result = maxHei* len(newArr)
    return result








h = [1,3,1,3,1,4,1,3,2,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,7]
word ="zaba"
print(designerPdfViewer(h, word))