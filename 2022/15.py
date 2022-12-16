import re

import common

PUZZLE = common.string(15)


def parse_sensors(puzzle):
    sensors = {}
    for line in puzzle.split("\n")[:-1]:
        sx, sy, bx, by = map(int, re.findall(r"-?\d+", line))
        sensors[(sx, sy)] = (bx, by)
    return sensors


def man(x, y):
    return abs(x[0] - y[0]) + abs(x[1] - y[1])


def get_union(sensors, n):
    spans = []
    for s, b in sensors.items():
        man_sb = man(s, b)
        if man((s[0], n), s) <= man_sb:
            dx = man_sb - abs(s[1] - n)
            spans.append((s[0] - dx, s[0] + dx))

    union = []
    for begin, end in sorted(spans):
        if union and union[-1][1] >= begin - 1:
            union[-1][1] = max(union[-1][1], end)
        else:
            union.append([begin, end])
    return union


def first(puzzle, n):
    sensors = parse_sensors(puzzle)
    union = get_union(sensors, n)
    res = 0
    for u in union:
        res += u[1] - u[0]
    return res


def second(puzzle):
    sensors = parse_sensors(puzzle)
    for n in range(4000001):
        union = get_union(sensors, n)
        if len(union) > 1:
            return (union[0][1] + 1) * 4000000 + n


# `pytest *`
def test():
    assert first(TEST_PUZZLE, 10) == 26
    assert first(PUZZLE, 2000000) == 5508234
    assert second(TEST_PUZZLE) == 56000011
    assert second(PUZZLE) == 10457634860779


if __name__ == "__main__":
    print(first(PUZZLE, 2000000))
    print(second(PUZZLE))

TEST_PUZZLE = """Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
"""
