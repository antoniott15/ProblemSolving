#k index
#bill arr
#b pay
def bonAppetit(bill, k, b):
    d = sum(bill)
    totalAmout = int(d/2)
    if totalAmout == b:
        toRef = b - ((d - bill[k])/2)
        return int(toRef)
    else:
        return "Bon Appetit"






k = 1
bill = [3,10,2,9]
b = 12

print(bonAppetit(bill,k,b))