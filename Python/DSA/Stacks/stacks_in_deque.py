from collections import deque
from inspect import stack

stacks = deque()

stacks.append("one")
stacks.append("two")
stacks.append("three")
stacks.append("four")

print("initial stack: ", stacks)

stacks.pop()
stacks.pop()
stacks.pop()
stacks.pop()

print(stacks)

