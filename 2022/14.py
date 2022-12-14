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


def drop_sand(sand, rock_lines, max_rock_y):
    sx, sy = 500, 0
    while True:
        if sy >= max_rock_y:
            return sand, True

        can_go_down = not any([intersects_line((sx, sy + 1), line) for line in rock_lines]) and (sx, sy + 1) not in sand
        if can_go_down:
            sy += 1
            continue

        can_go_left_down = not any([intersects_line((sx - 1, sy + 1), line) for line in rock_lines]) and (
        sx - 1, sy + 1) not in sand
        if can_go_left_down:
            sx, sy = sx - 1, sy + 1
            continue

        can_go_right_down = not any([intersects_line((sx + 1, sy + 1), line) for line in rock_lines]) and (
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
    sand = set()
    while True:
        sand, fell_forever = drop_sand(sand, rock_lines, max_rock_y)
        if fell_forever:
            break
    return len(sand)


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 24
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
