""" make a random list, prove you can count all unique occurrences of each item """

from __future__ import print_function
from collections import Counter
from Crypto.Random.random import randint

def random_list():
    """ Make random-length list of random numbers

    Returns:
        a list of ints
    """

    empty_list = []

    for _ in range(randint(10, 100)):
        empty_list.append(int(randint(1, 10)))

    return empty_list


def main():
    """ Main function, print how many times the most common number appears """

    my_list = random_list()
    print(my_list)
    print(Counter(my_list).most_common(1))


if __name__ == '__main__':
    main()
