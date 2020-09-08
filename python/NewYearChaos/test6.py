def minimumBribes(Q):
    moves = 0 
    Q = [P-1 for P in Q]
    print(Q)
    for i,P in enumerate(Q):
        if P - i > 2:
            print("Too chaotic")
            return
        print(max(P-1,0))
        for j in range(max(P-1,0),i):
            if Q[j] > P:
                moves += 1
    print(moves)
                


arr = [1,2,5,3,7,8,6,4]

minimumBribes(arr)