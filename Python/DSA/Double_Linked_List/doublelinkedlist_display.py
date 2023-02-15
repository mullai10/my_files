class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
        self.prev = None

class DLLD:
    def __init__(self):
        self.head = None
    
    def display(self):
        if self.head is None:
            print("Double Linked List is Empty!")
        else:
            temp = self.head
            while temp:
                print(temp.data,"-->",end=" ") 
                temp = temp.next

L = DLLD()

n1 = Node(10)
n2 = Node(20)
n3 = Node(30)
n4 = Node(40)
n5 = Node(50)

L.head = n1
n1.next = n2
n2.prev = n1
n2.next = n3
n3.prev = n2
n3.next = n4
n3.prev = n2
n4.next = n5
n4.prev = n3

L.display()
