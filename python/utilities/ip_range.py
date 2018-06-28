""" make a random list, prove you can count all uniqe occurrences of each item """

from __future__ import print_function

def rule_list():
    """ Make list of iptables rules for the 10.0.0.0/16 subnet

    Returns:
        a list of strings
    """

    ip_list = []

    for i in range(0, 256):
        for j in range(0, 256):
#            print('10.0.' + str(i) + '.' + str(j))
            ip_list.append('-A oo-udp-limit -s 10.0.' + str(i) + '.' + str(j) +\
            ' -m limit --limit 30/s -m comment --comment "oo_udp_limit_id_%s_%s" -j RETURN'\
            % (i, j))

    return ip_list


def main():
    """ Main function, print total number of rules generated """

    my_list = rule_list()
    print('total number of rules:', len(my_list))

    for i in my_list:
        print(i)


if __name__ == '__main__':
    main()
