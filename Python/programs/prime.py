# n1 = int(input("Enter the value: "))
# n2 = int(input("Enter the value: "))

# print("the prime numbers in the ranges are: ")

# for number in range(n1, n2+1):
#     if number > 1:
#         for i in range(2, number):
#             if (number % i) == 0:
#                 break
#         else:
#             print(number,end=" ")


# def prime_number(min, max):

#     prime_numbers_list = []
#     for num in range(min,max+1):
#         if num>1:
#             for i in range(2,num):
#                 if (num % i) == 0:
#                     break
#             else:
#                 prime_numbers_list.append(num)
#     return prime_numbers_list


# if __name__== "__main__":

#     min = int(input("Enter the minimum number: "))
#     max = int(input("Enter the maximum number: "))

#     print("The Prime Numbers from {} to {} are {}".format(min,max,prime_number(min, max)))
    
for i in range(1,101):
    if i>1:
        for j in range(2, i):
            if i%j == 0:
                break
        else:
            print(i,end=" ")