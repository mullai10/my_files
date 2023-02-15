from tkinter import N


class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
        self.prev = None
class DLLI:
    def __init__(self):
        self.head = None

    def insertion_biginning(self,data):
        nb = Node(data)
        temp = self.head
        temp.prev = nb
        nb.next = temp
        self.head = nb

    def insertion_end(self,data):
        ne = Node(data)
        temp = self.head
        while temp.next is not None:
            temp = temp.next
        temp.next = ne
        ne.prev = temp

    def insertion_pos(self,pos,data):
        np = Node(data)
        temp = self.head
        for i in range(1,pos-1):
            temp = temp.next
        np.prev = temp
        np.next = temp.next
        temp.next.prev = np
        temp.next = np

    def display(self):
        if self.head is None:
            print("Double Linked List is Empty!")
        else:
            temp = self.head
            while temp is not None:
                print(temp.data,"-->",end= " ")
                temp = temp.next 

L = DLLI()

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
n4.prev = n3
n4.next = n5
n5.prev = n4

# L.insertion_biginning(5)
# L.insertion_end(60)
# L.insertion_pos(3,25)

L.display()


