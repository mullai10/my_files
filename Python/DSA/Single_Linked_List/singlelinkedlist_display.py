from platform import node


class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
class SLL_DISPLAY:
    def __init__(self):
        self.head = None
        self.tail = None
        self.count = 0

    def append_items(self,data):
         node = Node(data)
         if self.tail:
            self.tail.next = node
            self.tail = node
         else:
            self.head = node
            self.tail = node
         self.count += 1    


    def display(self):
        if self.head is None:
            print("Single Linked List is Empty")
        else:
            temp = self.head
            while temp is not None:
                val = temp.data  # print(temp.data)
                temp = temp.next
                yield val

#to search a particular item, return true if it is present, otherwise return false
    def search_item(self,val):
        for node in self.display():
            if val == node:
                return True
            return False    

L = SLL_DISPLAY()

L.append_items(10)
L.append_items(20)
L.append_items(30)
L.append_items(40)

for val in L.display():
    print(val,end = " ")


L.display()

# if L.search_item(20)


# L.head = Node(10)
# n2 = Node(20)
# n3 = Node(30)
# n4 = Node(40)
# n5 = Node(50)
# L.head.next = n2
# n2.next = n3
# n3.next = n4
# n4.next = n5

# n1 = Node(10)
# L.head = n1
# n2 = Node(20)
# n1.next = n2
# n3 = Node(30)
# n2.next = n3
# n4 = Node(40)
# n3.next = n4
# n5 = Node(50)
# n4.next = n5


# print("size: ", L.count)

