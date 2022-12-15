import re

import common

PUZZLE = common.string(15)


def man(x, y):
    return abs(x[0] - y[0]) + abs(x[1] - y[1])


def first(puzzle, n):
    sensors = {}
    for line in puzzle.split("\n")[:-1]:
        sx, sy, bx, by = map(int, re.findall(r"-?\d+", line))
        sensors[(sx, sy)] = (bx, by)
    print(sensors)
    min_x, max_x = float("inf"), -float("inf")
    for s, b in sensors.items():
        min_x = min(min_x, s[0], b[0])
        max_x = max(max_x, s[0], b[0])
    # for every col x_i in row y, not a beacon if (is not already beacon) and
    # for any sensor, man(B, S_i) <= man((x_i, y), S_i)
    beacons = set(sensors.values())
    res = 0
    for x in range(min_x-6000000, max_x+6000000):
        if (x, n) in beacons:
            continue
        for s, b in sensors.items():
            if man((x, n), s) <= man(b, s):
                res += 1
                break

    return res


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE, 10) == 26
    assert first(PUZZLE, 2000000) == 5508234
    assert second(TEST_PUZZLE) == -1


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
