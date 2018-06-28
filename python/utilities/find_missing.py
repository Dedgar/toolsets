""" Find the numbers not included in the provided list, within its range. """

from __future__ import print_function


MY_LIST = [1, 2, 3, 4, 5, 10, 11, 15, 17, 20]

NOT_HERE = []

for i in range(1, 21):
    if not i in MY_LIST:
        NOT_HERE.append(i)

print(MY_LIST)
print(NOT_HERE)
