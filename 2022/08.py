from __future__ import annotations

import common

PUZZLE = common.string(8)

TEST_PUZZLE = """30373
25512
65332
33549
35390
"""


def parse(puzzle):
    return [[int(n) for n in line] for line in puzzle.split("\n") if len(line)]


def peek(trees, visible, x, y, dx, dy):
    w, h = len(trees[0]), len(trees)
    max_height = trees[y][x]
    x, y = x + dx, y + dy
    while 0 <= x < w and 0 <= y < h:
        if trees[y][x] > max_height:
            visible[y][x] = True
        max_height = max(trees[y][x], max_height)
        x, y = x + dx, y + dy
    return visible


def first(puzzle):
    trees = parse(puzzle)
    w, h = len(trees[0]), len(trees)
    visible = [
        [
            True if i == 0 or i == w - 1 or j == 0 or j == h - 1 else False
            for i, _ in enumerate(range(w))
        ] for j, _ in enumerate(range(h))
    ]
    for i in range(1, w - 1):
        visible = peek(trees, visible, i, 0, 0, 1)  # top
        visible = peek(trees, visible, i, h - 1, 0, -1)  # bottom
    for j in range(1, h - 1):
        visible = peek(trees, visible, 0, j, 1, 0)  # left
        visible = peek(trees, visible, w - 1, j, -1, 0)  # right
    return sum(sum(l) for l in visible)


def view_dist(trees, x, y, dx, dy):
    tree_height = trees[y][x]
    w, h = len(trees[0]), len(trees)
    d = 0
    x, y = x + dx, y + dy
    while 0 <= x < w and 0 <= y < h:
        d += 1
        if trees[y][x] >= tree_height:
            break
        x, y = x + dx, y + dy
    return d


def scenic_score(trees, x, y):
    return (
        view_dist(trees, x, y, 1, 0)
        * view_dist(trees, x, y, -1, 0)
        * view_dist(trees, x, y, 0, 1)
        * view_dist(trees, x, y, 0, -1)
    )


def second(puzzle):
    trees = parse(puzzle)
    w, h = len(trees[0]), len(trees)
    max_scenic_score = 0
    for x in range(1, w - 1):
        for y in range(1, h - 1):
            max_scenic_score = max(
                max_scenic_score,
                scenic_score(trees, x, y)
            )
    return max_scenic_score


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 21
    assert first(PUZZLE) == 1801
    assert second(TEST_PUZZLE) == 8
    assert second(PUZZLE) == 209880


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
