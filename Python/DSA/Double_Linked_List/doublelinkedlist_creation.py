class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
        self.prev = None

class DLL:
    def __init__(self):
        self.head = None
    def display(self):
        if self.head is None:
            print("Double Linked List is Empty!")

L = DLL()
L.display()
                