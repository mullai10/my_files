number = int(input("Enter the digits you want: "))
n1 = int(input("Enter the n1: "))
n2 = int(input("Enter the n2: "))

if number <= 0:
    print("please enter the positive integer!")
elif number == 1:
    print("the fibonacci series is",n1)
else:
    print("the fibonacci series is ")
    # i = 0
    # while i < number:
    #     print(n1)
    #     n3 = n1 + n2
    #     n1 = n2
    #     n2 = n3
    #     i = i + 1     
    i = n1
    for i in range(n1,number+1):
        n3 = n1 + n2
        n1 = n2
        n2 = n3
        i = i+1
        print(n1)
