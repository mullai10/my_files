class Node:
    def __init__(self,data):
        self.data = data
        self.next = None

class SLL:
    def __init__(self):
        self.head = None

    def display(self):
        if self.head is None:
            print("Single Linked List is Created!")
       
L = SLL()

L.display()
