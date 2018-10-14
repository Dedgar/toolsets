""" Find if the string representation of a number is greater or less than another. """

from __future__ import print_function


MY_LIST = ['1', '2', '3', '4', '5', '10', '11', '15', '17', '20', '700']
MY_NUM = '7'

print(MY_LIST)
for i in MY_LIST:
    if len(i) == len(MY_NUM):# and MY_NUM > i:
        for each_char in i:
            if MY_NUM[i.index(each_char)] > each_char:
                print('%s is greater than %s' % (MY_NUM, i))
    else:
        print('%s is less than %s' % (MY_NUM, i))
