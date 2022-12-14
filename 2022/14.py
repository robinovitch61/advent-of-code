import re

import common

PUZZLE = common.string(14)

TEST_PUZZLE = """498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
"""


def intersects_line(point, line):
    """only works for vertical or horizontal lines"""
    if line[0][0] == line[1][0]:  # vertical
        return point[0] == line[0][0] and min(line[0][1], line[1][1]) <= point[1] <= max(line[0][1], line[1][1])
    if line[0][1] == line[1][1]:  # horizontal
        return point[1] == line[0][1] and min(line[0][0], line[1][0]) <= point[0] <= max(line[0][0], line[1][0])
    assert False


def drop_sand(sand, rocks, max_rock_y):
    sx, sy = 500, 0
    while True:
        if sy >= max_rock_y:
            return sand, True

        can_go_down = (sx, sy + 1) not in rocks and (sx, sy + 1) not in sand
        if can_go_down:
            sy += 1
            continue

        can_go_left_down = (sx - 1, sy + 1) not in rocks and (
            sx - 1, sy + 1) not in sand
        if can_go_left_down:
            sx, sy = sx - 1, sy + 1
            continue

        can_go_right_down = (sx + 1, sy + 1) not in rocks and (
            sx + 1, sy + 1) not in sand
        if can_go_right_down:
            sx, sy = sx + 1, sy + 1
            continue

        break
    sand.add((sx, sy))
    return sand, False


def first(puzzle):
    rock_lines = []
    max_rock_y = -float("inf")
    for line in puzzle.split("\n")[:-1]:
        points = [tuple(map(int, pair.split(","))) for pair in re.findall(r"\d+,\d+", line)]
        for p in points:
            max_rock_y = max(max_rock_y, p[1])
        for i in range(len(points) - 1):
            rock_lines.append((points[i], points[i + 1]))
    rocks = set()
    for rl in rock_lines:
        if rl[0][0] == rl[1][0]:  # vertical
            for y in range(min(rl[0][1], rl[1][1]), max(rl[0][1], rl[1][1]) + 1):
                rocks.add((rl[0][0], y))
        if rl[0][1] == rl[1][1]:  # horizontal
            for x in range(min(rl[0][0], rl[1][0]), max(rl[0][0], rl[1][0]) + 1):
                rocks.add((x, rl[0][1]))
    sand = set()
    while True:
        sand, fell_forever = drop_sand(sand, rocks, max_rock_y)
        if fell_forever:
            break
    return len(sand)


def drop_sand_two(sand, rocks):
    sx, sy = 500, 0
    while True:
        can_go_down = (sx, sy + 1) not in rocks and (sx, sy + 1) not in sand
        if can_go_down:
            sy += 1
            continue

        can_go_left_down = (sx - 1, sy + 1) not in rocks and (
            sx - 1, sy + 1) not in sand
        if can_go_left_down:
            sx, sy = sx - 1, sy + 1
            continue

        can_go_right_down = (sx + 1, sy + 1) not in rocks and (
            sx + 1, sy + 1) not in sand
        if can_go_right_down:
            sx, sy = sx + 1, sy + 1
            continue

        break
    sand.add((sx, sy))
    return sand


def second(puzzle):
    rock_lines = []
    max_rock_y = -float("inf")
    min_rock_x = float("inf")
    max_rock_x = -float("inf")
    for line in puzzle.split("\n")[:-1]:
        points = [tuple(map(int, pair.split(","))) for pair in re.findall(r"\d+,\d+", line)]
        for p in points:
            max_rock_y = max(max_rock_y, p[1])
            min_rock_x = min(min_rock_x, p[0])
            max_rock_x = max(max_rock_x, p[0])
        for i in range(len(points) - 1):
            rock_lines.append((points[i], points[i + 1]))
    rock_lines.append(((int(min_rock_x - 1e3), max_rock_y + 2), (int(max_rock_x + 1e3), max_rock_y + 2)))
    rocks = set()
    for rl in rock_lines:
        if rl[0][0] == rl[1][0]:  # vertical
            for y in range(min(rl[0][1], rl[1][1]), max(rl[0][1], rl[1][1]) + 1):
                rocks.add((rl[0][0], y))
        if rl[0][1] == rl[1][1]:  # horizontal
            for x in range(min(rl[0][0], rl[1][0]), max(rl[0][0], rl[1][0]) + 1):
                rocks.add((x, rl[0][1]))
    sand = set()
    while True:
        sand = drop_sand_two(sand, rocks)
        if (500, 0) in sand:
            break
    return len(sand)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 24
    assert second(TEST_PUZZLE) == 93


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
