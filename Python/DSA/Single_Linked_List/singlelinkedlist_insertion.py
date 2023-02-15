class Node:
    def __init__(self,data):
        self.data = data
        self.next = None

class SLL_Insertion:
    def __init__(self):
        self.head = None

    def insertion_beginning(self,data):
        nb = Node(data)
        nb.next = self.head
        self.head = nb

    def insertion_end(self,data):
        ne = Node(data)
        temp = self.head
        while temp.next:
            temp = temp.next
        temp.next = ne

    def insertion_pos(self,pos,data):
        np = Node(data)
        temp = self.head
        for i in range(pos-1):
            temp = temp.next
        np.data = data
        np.next = temp.next
        temp.next = np

    def display(self):
        if self.head is None:
            print("Single Linked List is Empty")
        else:
            temp = self.head
            while temp:
                print(temp.data,"-->",end=" ")
                temp = temp.next

L = SLL_Insertion()
n1 = Node(10)
L.head = n1
n2 = Node(20)
n1.next = n2
n3 = Node(30)
n2.next = n3
n4 = Node(40)
n3.next = n4
n5 = Node(50)
n4.next = n5

L.insertion_beginning(5)

# L.insertion_end(60)

# L.insertion_pos(2,25)

L.display()