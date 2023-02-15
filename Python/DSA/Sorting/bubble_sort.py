def bubble_sort(list):
    for j in range(0,len(list)):
        for i in range(0,len(list)-j-1):
            # ascending order
            if (list[i]>list[i+1]):
            # decending order
            # if (list[i]<list[i+1]):
                temp = list[i]
                list[i]= list[i+1]
                list[i+1]= temp
                # list[i],list[i+1] = list[i+1],list[i]

if __name__=="__main__":
    list = [2,56,44,31,67,10,4,8,100,32,42]    
    bubble_sort(list)
    print(list)
