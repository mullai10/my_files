from math import factorial


# n = int(input("Enter the value of n:"))
# factorial = 1
# i = 1
# while(i<=n):
#     factorial = factorial * i
#     print("factorial is ", factorial)
#     i= i+1

def factorial(n):
    f = 1
    i = 1
    while (i<= n):
        f = f * i
        print("factorial is ", f)
        i= i+1

factorial(6)
