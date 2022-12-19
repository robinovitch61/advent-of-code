import re
from collections import Counter

import common

PUZZLE = common.string(18)

DIFFS = ((-1, 0, 0), (1, 0, 0), (0, -1, 0), (0, 1, 0), (0, 0, -1), (0, 0, 1))


def add_points(a, b):
    return tuple(sum(pair) for pair in zip(a, b))


def get_points(puzzle):
    points = set()
    for l in puzzle.split("\n")[:-1]:
        points.add(tuple(map(int, re.findall("\d+", l))))
    return points


def area_all(points):
    sides = Counter()
    for p in points:
        for d in DIFFS:
            sides[add_points(p, d)] += 1
    for p in points:
        del sides[p]
    return sum(sides.values())


def first(puzzle):
    points = get_points(puzzle)
    return area_all(points)


def get_ranges(points):
    xs, ys, zs = zip(*points)
    return [min(xs), max(xs)], [min(ys), max(ys)], [min(zs), max(zs)]


def second(puzzle):
    points = get_points(puzzle)
    x_range, y_range, z_range = get_ranges(points)

    def is_surrounded(x, y, z):
        def closed(v, r, f):
            vv = v
            while True:
                if vv < r[0]:
                    return False
                if f(vv) in points:
                    break
                vv -= 1
            vv = v + 1
            while True:
                if vv > r[1]:
                    return False
                if f(vv) in points:
                    break
                vv += 1
            return True

        x_closed = closed(x, x_range, lambda xx: (xx, y, z))
        y_closed = closed(y, y_range, lambda yy: (x, yy, z))
        z_closed = closed(z, z_range, lambda zz: (x, y, zz))
        return x_closed and y_closed and z_closed

    surrounded = set()
    for x in range(x_range[0], x_range[1] + 1):
        for y in range(y_range[0], y_range[1] + 1):
            for z in range(z_range[0], z_range[1] + 1):
                if (x, y, z) not in points and is_surrounded(x, y, z):
                    surrounded.add((x, y, z))

    sub = 0
    for s in surrounded:
        for d in DIFFS:
            if add_points(s, d) in points:
                sub += 1

    return area_all(points) - sub


def test():
    assert first(TEST_PUZZLE) == 64
    assert first(PUZZLE) == 3470
    assert second(TEST_PUZZLE) == 58
    assert second(PUZZLE) == 1986


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
