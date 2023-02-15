def selection_sort(list):
    for j in range(0,len(list)):
        smallest = list[j]
        smallest_index = j
        for i in range(j+1,len(list)):
            if list[i]<smallest:
                smallest = list[i]
                smallest_index = i
        list[j],list[smallest_index] = list[smallest_index],list[j]

if __name__=="__main__":
    list = [2,56,44,31,67,10,4,8,100,32,42]    
    selection_sort(list)
    print(list)
