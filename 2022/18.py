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
    points = []
    for l in puzzle.split("\n")[:-1]:
        points.append(tuple(map(int, re.findall("\d+", l))))
    xy = [["." for _ in range(10)] for _ in range(10)]
    xz = [["." for _ in range(10)] for _ in range(10)]
    yz = [["." for _ in range(10)] for _ in range(10)]
    xypoints = [(x, y) for x, y, z in points]
    xzpoints = [(x, z) for x, y, z in points]
    yzpoints = [(y, z) for x, y, z in points]
    for x in range(10):
        for y in range(10):
            for z in range(10):
                if (x, y) in xypoints:
                    xy[x][y] = "#"
                if (x, z) in xzpoints:
                    xz[x][z] = "#"
                if (y, z) in yzpoints:
                    yz[y][z] = "#"
    print("\n".join(["".join(l) for l in xy]), "\n")
    print("\n".join(["".join(l) for l in xz]), "\n")
    print("\n".join(["".join(l) for l in yz]), "\n")
    return -1


def test():
    # assert first(TEST_PUZZLE) == 64
    # assert first(PUZZLE) == 3470
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
