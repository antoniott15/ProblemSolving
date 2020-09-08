def maximumToys(prices, k):
    prices.sort()
    elementos = 0
    countingItems = 0
    for elements in prices:
        if elementos <= k:
            elementos += elements
            countingItems += 1

    if elementos > k:
        countingItems -= 1
    return countingItems
