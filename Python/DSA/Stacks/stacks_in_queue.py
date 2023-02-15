import queue


from queue import LifoQueue

# initializing the stack
stacks = LifoQueue(maxsize=4)

#qsize() show the number of elements in the stack
print("Number of elements at start: ", stacks.qsize())

#put() function to push the elements in the stack
stacks.put(10)
stacks.put(20)
stacks.put(30)
stacks.put(40)

print("stack is full: ", stacks.full())
print("size after adding the elements: ", stacks.qsize())

#get() function to pop the elements in stacks in LIFO order
stacks.get()
stacks.get()
stacks.get()
stacks.get()

print("stack is empty: ", stacks.empty())


