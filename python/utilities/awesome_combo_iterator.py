"""
This will print every possible combination of items in a list
"""

from __future__ import print_function
from itertools import chain, combinations


def all_subsets(some_list):
    """ Chains through every iteration of the provided list

    Returns:
        an itertools chain of the provided list
    """
    #    return chain(*map(lambda x: combinations(some_list, x), range(0, len(some_list)+1)))
    return chain(*[combinations(some_list, x) for x in range(0, len(some_list)+1)])


def main():
    """ Main """
    range_of_stuff = list(range(0, 4))
    for subset in all_subsets(range_of_stuff):
        print(subset)

if __name__ == '__main__':
    main()
