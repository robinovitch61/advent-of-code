import re

import common

PUZZLE = common.string(14)

TEST_PUZZLE = """498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
"""


def get_rock_lines(puzzle):
    rock_lines = []
    min_x, max_x, max_y = float("inf"), -float("inf"), -float("inf")
    for line in puzzle.split("\n")[:-1]:
        points = [tuple(map(int, pair.split(","))) for pair in re.findall(r"\d+,\d+", line)]
        for point in points:
            min_x, max_x, max_y = min(min_x, point[0]), max(max_x, point[0]), max(max_y, point[1])
        for i in range(len(points) - 1):
            rock_lines.append((points[i], points[i + 1]))
    return rock_lines, min_x, max_x, max_y


def get_rocks(rock_lines):
    rocks = set()
    for rock_line in rock_lines:
        rocks |= get_rocks_in_line(rock_line)
    return rocks


def get_rocks_in_line(line):
    rocks = set()
    if line[0][0] == line[1][0]:  # vertical
        for y in range(min(line[0][1], line[1][1]), max(line[0][1], line[1][1]) + 1):
            rocks.add((line[0][0], y))
    if line[0][1] == line[1][1]:  # horizontal
        for x in range(min(line[0][0], line[1][0]), max(line[0][0], line[1][0]) + 1):
            rocks.add((x, line[0][1]))
    return rocks


def drop_sand(sand, rocks, max_rock_y):
    sx, sy = 500, 0
    while True:
        if sy >= max_rock_y:
            return sand, True

        if (sx, sy + 1) not in sand and (sx, sy + 1) not in rocks:
            sy += 1
            continue

        if (sx - 1, sy + 1) not in sand and (sx - 1, sy + 1) not in rocks:
            sx, sy = sx - 1, sy + 1
            continue

        if (sx + 1, sy + 1) not in sand and (sx + 1, sy + 1) not in rocks:
            sx, sy = sx + 1, sy + 1
            continue

        break
    sand.add((sx, sy))
    return sand, False


def first(puzzle):
    rock_lines, _, _, max_y = get_rock_lines(puzzle)
    rocks = get_rocks(rock_lines)
    sand = set()
    while True:
        sand, fell_forever = drop_sand(sand, rocks, max_y)
        if fell_forever:
            break
    return len(sand)


def second(puzzle):
    rock_lines, min_x, max_x, max_y = get_rock_lines(puzzle)
    max_y += 2
    rock_lines.append(((min_x - 500, max_y), (max_x + 500, max_y)))
    rocks = get_rocks(rock_lines)
    sand = set()
    while True:
        sand, _ = drop_sand(sand, rocks, max_y)
        if (500, 0) in sand:
            break
    return len(sand)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 24
    assert first(PUZZLE) == 774
    assert second(TEST_PUZZLE) == 93
    assert second(PUZZLE) == 22499


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
