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
    sides = []
    diffs = ((-1, 0, 0), (1, 0, 0), (0, -1, 0), (0, 1, 0), (0, 0, -1), (0, 0, 1))
    for p in sorted(points):
        for d in diffs:
            sides.append(tuple(sum(pair) for pair in zip(p, d)))
    for p in points:
        while p in sides:
            sides.pop(sides.index(p))

    x_range = [float("inf"), -float("inf")]
    y_range = [float("inf"), -float("inf")]
    z_range = [float("inf"), -float("inf")]
    for p in points:
        x_range[0] = min(p[0], x_range[0])
        x_range[1] = max(p[0], x_range[1])
        y_range[0] = min(p[1], y_range[0])
        y_range[1] = max(p[1], y_range[1])
        z_range[0] = min(p[2], z_range[0])
        z_range[1] = max(p[2], z_range[1])

    def is_surrounded(x, y, z):
        if (x, y, z) in points:
            return False
        xx = x
        while True:
            if xx < x_range[0]:
                return False
            if (xx, y, z) in points:
                break
            xx -= 1
        xx = x
        while True:
            if xx > x_range[1]:
                return False
            if (xx, y, z) in points:
                break
            xx += 1
        yy = y
        while True:
            if yy < y_range[0]:
                return False
            if (x, yy, z) in points:
                break
            yy -= 1
        yy = y
        while True:
            if yy > y_range[1]:
                return False
            if (x, yy, z) in points:
                break
            yy += 1
        zz = z
        while True:
            if zz < z_range[0]:
                return False
            if (x, y, zz) in points:
                break
            zz -= 1
        zz = z
        while True:
            if zz > z_range[1]:
                return False
            if (x, y, zz) in points:
                break
            zz += 1
        return True

    surrounded = set()
    for x in range(x_range[0], x_range[1] + 1):
        for y in range(y_range[0], y_range[1] + 1):
            for z in range(z_range[0], z_range[1] + 1):
                if is_surrounded(x, y, z):
                    surrounded.add((x, y, z))

    sub = 0
    diffs = ((-1, 0, 0), (1, 0, 0), (0, -1, 0), (0, 1, 0), (0, 0, -1), (0, 0, 1))
    for s in surrounded:
        for d in diffs:
            if tuple(sum(pair) for pair in zip(s, d)) in points:
                sub += 1

    return len(sides) - sub


def test():
    # assert first(TEST_PUZZLE) == 64
    # assert first(PUZZLE) == 3470
    assert second(TEST_PUZZLE) == 58


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """2,2,2
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
