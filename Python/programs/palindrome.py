# my_str = 'MaDaM'
# my_str = my_str.casefold()
# my_rev = reversed(my_str)
# if (list(my_str) == list(my_rev)) :
#     print("This is Palindrome!")
# else:
#     print("This is not Palindrome!")

# print("-------------------------")    

# string = input("Enter the word: ")
# if (string == string[::-1]):
#     print("Given word is palindrome!")
# else:
#     print("Not palindrome!")    

# print("-------------------------")  

# num=int(input("Enter a number:"))
# temp=num
# rev=0
# while(num>0):
#     dig=num%10
#     rev=rev*10+dig
#     num=num//10
# if(temp==rev):
#     print("The number is palindrome!")
# else:
#     print("Not a palindrome!")    

# print("-------------------------")  

# def palindrome(s):
#     return s == s[::-1]
# s = "mam"
# ans = palindrome(s)
# if ans:
#     print("Yes")
# else:
#     print("No")   
# 
# print("-------------------------")  

list = []
rev = 0
for i in range(10,1000):
    while(i>0):
        dig = i%10
        rev = (rev*10) + dig
        i = i // 10
        list += [i]
        
        if (i == list[i]):
            print(list)


    

    
    