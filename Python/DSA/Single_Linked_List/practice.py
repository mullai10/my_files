# from pickletools import read_uint1


# class Node:
#     def __init__(self,data):
#         self.data = data
#         self.next = None
# class SLL:
#     def __init__(self):
#         self.head = None
#         self.tail = None
#         self.count = 0

#     def display(self):
#         if self.head is None:
#             print("empty")
#         else:
#             temp = self.head
#             while temp:
#                 val = temp.data
#                 temp = temp.next
#                 yield val

#     def append(self,data):
#         node = Node(data)
#         if self.tail:
#             self.tail.next = node
#             self.tail = node
#         else:
#             self.head = node
#             self.tail = node
#         self.count += 1

#     def search_item(self,val):
#         for node in self.display():
#             if val == node:
#                 return True
#         return False

#     def __getitem__(self,index):
#         if index  > self.count -1:
#             return "Index out of range"
#         temp = self.tail
#         for i in range(index):
#             temp = temp.next
#         return temp.data


# L = SLL()

# L.append(10)
# L.append(20)
# L.append(30)
# L.append(40)

# for val in L.display():
#     print(val)

# print("size of the SLL: ", L.count)    

# print(L[0])


# if L.search_item(80):
#     print("True")
# else:
#     print("False")

# i=1
# while i <= 50:
#     if i%3==0:
#         print("Fizz", end="")
#         if i%5==0:
#             print("Buzz", end="")
#     elif i%5==0:
#         print("Buzz", end="")
#     else:
#         print(i, end="")
#     print()
#     i+=1
def fizzbuzz(n):
    for i in range(1,n+1):
        if  i %3 == 0 and i %5 == 0:
            print("Fizzbuzz")
        elif i %3 == 0 and i %5 != 0:
            print("Fizz")
        elif i %5 == 0 and i %3 != 0:
            print("Buzz")
        else:
            print(i)
if __name__=='__main__':
    n = int(input())     
    fizzbuzz(n)
