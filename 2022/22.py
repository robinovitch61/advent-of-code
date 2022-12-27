import re

import common

PUZZLE = common.string(22)

TEST_PUZZLE = """        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
"""

WALL = "#"
OPEN = "."
VOID = " "
CLOCKWISE = {
    ">": "v",
    "v": "<",
    "<": "^",
    "^": ">",
}
COUNTER = {v: k for k, v in CLOCKWISE.items()}


def parse(puzzle):
    board, commands = puzzle.split("\n\n")
    commands = re.findall(r"(\d+|R|L)", commands)
    commands = [int(c) if c not in ("L", "R") else c for c in commands]
    board_lines = board.split("\n")
    board = {}
    for j, line in enumerate(board_lines):
        for i, c in enumerate(line):
            if c == OPEN:
                board[(j, i)] = OPEN
            elif c == WALL:
                board[(j, i)] = WALL
    return board, commands


def move(board, r, c, d):
    if d == ">":
        if board.get((r, c + 1)) == OPEN:
            return r, c + 1
        elif board.get((r, c + 1)) == WALL:
            return r, c
        else:
            cc = min(col for row, col in board.keys() if row == r)
            if board[(r, cc)] == WALL:
                return r, c
            else:
                return r, cc
    elif d == "<":
        if board.get((r, c - 1)) == OPEN:
            return r, c - 1
        elif board.get((r, c - 1)) == WALL:
            return r, c
        else:
            cc = max(col for row, col in board.keys() if row == r)
            if board[(r, cc)] == WALL:
                return r, c
            else:
                return r, cc
    elif d == "^":
        if board.get((r - 1, c)) == OPEN:
            return r - 1, c
        if board.get((r - 1, c)) == WALL:
            return r, c
        else:
            rr = max(row for row, col in board.keys() if col == c)
            if board[(rr, c)] == WALL:
                return r, c
            else:
                return rr, c
    elif d == "v":
        if board.get((r + 1, c)) == OPEN:
            return r + 1, c
        if board.get((r + 1, c)) == WALL:
            return r, c
        else:
            rr = min(row for row, col in board.keys() if col == c)
            if board[(rr, c)] == WALL:
                return r, c
            else:
                return rr, c
    else:
        raise Exception(d)


def first(puzzle):
    board, commands = parse(puzzle)
    r, c = min(board.keys())  # leftmost open tile
    d = ">"
    for command in commands:
        if command == "L":
            d = COUNTER[d]
        elif command == "R":
            d = CLOCKWISE[d]
        else:
            for _ in range(command):
                new_r, new_c = move(board, r, c, d)
                if new_r == r and new_c == c:
                    break
                r, c = new_r, new_c
    return 1000 * (r + 1) + 4 * (c + 1) + list(CLOCKWISE.keys()).index(d)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == 6032
    assert first(PUZZLE) == 126350
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
