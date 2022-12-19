import re

import common

PUZZLE = common.string(18)


def first(puzzle):
    points = []
    for l in puzzle.split("\n")[:-1]:
        points.append(tuple(map(int, re.findall("\d+", l))))
    sides = []
    diffs = ((-1, 0, 0), (1, 0, 0), (0, -1, 0), (0, 1, 0), (0, 0, -1), (0, 0, 1))
    for p in sorted(points):
        for d in diffs:
            sides.append(tuple(sum(pair) for pair in zip(p, d)))
    for p in points:
        while p in sides:
            sides.pop(sides.index(p))
    return len(sides)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == 64
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))


TEST_PUZZLE="""2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
"""
