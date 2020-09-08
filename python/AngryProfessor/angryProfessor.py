def angryProfessor(k, a):
    count = 0
    for elements in a:
        if elements<=0:
            count+=1
    if count>=k :
        return "NO"
    else: 
        return "YES"




k = 2
a = [0,-1,2,1]
print(angryProfessor(k,a))