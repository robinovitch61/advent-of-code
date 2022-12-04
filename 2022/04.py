import re

import common

PUZZLE = common.string(4)


# first solution was a lot less clean, involved ranges and any/all
# REMEMBER TO CONVERT STRINGS TO INTS YA SILLY LEO

def sum_condition(puzzle, f):
    return sum(f(*(int(n) for n in re.findall(r"\d+", line))) for line in puzzle.split())


def first(puzzle):
    return sum_condition(puzzle, lambda a, b, c, d: a <= c and b >= d or a >= c and b <= d)


def second(puzzle):
    return sum_condition(puzzle, lambda a, b, c, d: max(a, c) <= min(b, d))


TEST_PUZZLE = """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
"""


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 2
    assert second(TEST_PUZZLE) == 4
    assert first(PUZZLE) == 305
    assert second(PUZZLE) == 811


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
