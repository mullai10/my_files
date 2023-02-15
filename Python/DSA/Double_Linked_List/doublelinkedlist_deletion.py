class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
        self.prev = None
class DLLD:
    def __init__(self):
        self.head = None

    def deletion_biginning(self):
        temp = self.head
        self.head = temp.next
        temp.next = None
        self.head.prev = None    

    def deletion_end(self):
        temp = self.head.next
        before = self.head
        while temp.next is not None:
            temp = temp.next
            before = before.next
        before.next = None
        temp.prev = None    

    def deletion_pos(self,pos):
        temp = self.head.next
        before = self.head
        for i in range(1,pos-1):
            temp = temp.next
            before = before.next
        before.next = temp.next
        temp.next.prev = before
        temp.next = None
        temp.prev = None
        
    def display(self):
        if self.head is None:
            print("Double Linked List is Empty!")
        else:
            temp = self.head
            while temp is not None:
                print(temp.data,"-->",end= " ")
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
n4.prev = n3
n4.next = n5
n5.prev = n4

# L.deletion_biginning()

# L.deletion_end()

L.deletion_pos(2)

L.display()