#!/bin/env python
# vim: expandtab:tabstop=4:shiftwidth=4

""" test. """


def answer(some_number):
    """ Find the lowest number of operations
        needed for the provided int.

    Returns:
        An integer.
    """

    steps = 0
    while some_number > 1:

        # print(some_number)
        if some_number % 2 == 0:
            some_number /= 2
            steps += 1
        else:
            if (some_number - 1) / 2 % 2 == 0 or \
               (some_number - 1) / 2 == 1:
                some_number -= 1
                steps += 1
            else:
                some_number += 1
                steps += 1

    return steps


def main():
    """ Main function. """

    print(answer(99))


if __name__ == '__main__':
    main()
