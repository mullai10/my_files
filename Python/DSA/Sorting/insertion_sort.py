def insertion_sort(list):
    for i in range(0,len(list)):
        cur = list[i]
        for j in range(i-1,-1,-1):
            if list[j]>cur:
                list[j+1]=list[j]
            else:
                list[j+1]=cur
                break
    
if __name__=="__main__":
    list = [2,56,44,31,67,10,4,8,100,32,42]    
    insertion_sort(list)
    print(list)
    