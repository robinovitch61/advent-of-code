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
UP, DOWN, LEFT, RIGHT = "^", "v", "<", ">"

# These are manually defined based on some paper cubes I made
# (current_face, direction): (target_face, new_direction, flip)
TEST_TELEPORT_MAP = {
    (1, UP): (2, DOWN, True),
    (1, LEFT): (3, DOWN, False),
    (1, RIGHT): (6, LEFT, True),
    (2, UP): (1, DOWN, True),
    (2, DOWN): (5, UP, True),
    (2, LEFT): (6, UP, True),
    (3, UP): (1, RIGHT, False),
    (3, DOWN): (5, RIGHT, True),
    (4, RIGHT): (6, DOWN, True),
    (5, DOWN): (2, UP, True),
    (5, LEFT): (3, UP, True),
    (6, UP): (4, LEFT, True),
    (6, DOWN): (2, RIGHT, True),
    (6, RIGHT): (1, LEFT, True),
}
TELEPORT_MAP = {
    (1, UP): (6, RIGHT, False),
    (1, LEFT): (4, RIGHT, True),
    (2, UP): (6, UP, False),
    (2, DOWN): (3, LEFT, False),
    (2, RIGHT): (5, LEFT, True),
    (3, LEFT): (4, DOWN, False),
    (3, RIGHT): (2, UP, False),
    (4, UP): (3, RIGHT, False),
    (4, LEFT): (1, RIGHT, True),
    (5, DOWN): (6, LEFT, False),
    (5, RIGHT): (2, LEFT, True),
    (6, DOWN): (2, DOWN, False),
    (6, LEFT): (1, DOWN, False),
    (6, RIGHT): (5, UP, False),
}


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
    if d == RIGHT:
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
    elif d == LEFT:
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
    elif d == UP:
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
    elif d == DOWN:
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
    d = RIGHT
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


def get_face_info(board):
    rs, cs = zip(*board.keys())
    height, width = max(rs) - min(rs), max(cs) - min(cs)
    side = (min(height, width) + 1) // 3
    face = 0
    pos_to_face = {}
    face_to_top_left_pos = {}
    down = 0
    across = 0
    while True:
        while (down, across) not in board:
            across += 1
            if across > side * 4:
                down += side
                if down > side * 4:
                    return side, pos_to_face, face_to_top_left_pos
                across = 0
        face += 1
        face_to_top_left_pos[face] = (down, across)
        for i in range(side):
            for j in range(side):
                assert (down + i, across + j) in board
                pos_to_face[(down + i, across + j)] = face
        across += side


def move_cube(board, teleport_func, r, c, d):
    if d == RIGHT:
        if board.get((r, c + 1)) == OPEN:
            return r, c + 1, d
        elif board.get((r, c + 1)) == WALL:
            return r, c, d
        else:
            return teleport_func(r, c, d)
    elif d == LEFT:
        if board.get((r, c - 1)) == OPEN:
            return r, c - 1, d
        elif board.get((r, c - 1)) == WALL:
            return r, c, d
        else:
            return teleport_func(r, c, d)
    elif d == UP:
        if board.get((r - 1, c)) == OPEN:
            return r - 1, c, d
        if board.get((r - 1, c)) == WALL:
            return r, c, d
        else:
            return teleport_func(r, c, d)
    elif d == DOWN:
        if board.get((r + 1, c)) == OPEN:
            return r + 1, c, d
        if board.get((r + 1, c)) == WALL:
            return r, c, d
        else:
            return teleport_func(r, c, d)
    else:
        raise Exception(d)


def make_teleport_func(board, side, pos_to_face, face_to_top_left_pos, teleport_map):
    """Returns a function that maps an edge position (and direction?) to the next position and direction."""

    def func(r, c, d):
        curr_face = pos_to_face[(r, c)]
        next_face, next_direction, flip = teleport_map[(curr_face, d)]
        curr_down, curr_across = face_to_top_left_pos[curr_face]
        next_down, next_across = face_to_top_left_pos[next_face]
        if d in (LEFT, RIGHT):
            offset = r - curr_down
        elif d in (UP, DOWN):
            offset = c - curr_across
        else:
            raise Exception(d)
        if flip:
            offset = side - 1 - offset
        # assert offset > 0

        if next_direction == UP:
            rr = next_down + side - 1
            cc = next_across + offset
        elif next_direction == DOWN:
            rr = next_down
            cc = next_across + offset
        elif next_direction == LEFT:
            rr = next_down + offset
            cc = next_across + side - 1
        elif next_direction == RIGHT:
            rr = next_down + offset
            cc = next_across
        else:
            raise Exception(next_direction)

        if board[(rr, cc)] == WALL:
            return r, c, d

        return rr, cc, next_direction

    return func


def second(puzzle, teleport_map):
    board, commands = parse(puzzle)
    side, pos_to_face, face_to_top_left_pos = get_face_info(board)
    teleport_func = make_teleport_func(board, side, pos_to_face, face_to_top_left_pos, teleport_map)
    r, c = min(board.keys())  # leftmost open tile
    d = RIGHT
    for command in commands:
        if command == "L":
            d = COUNTER[d]
        elif command == "R":
            d = CLOCKWISE[d]
        else:
            for _ in range(command):
                new_r, new_c, new_d = move_cube(board, teleport_func, r, c, d)
                if new_r == r and new_c == c:
                    break
                r, c, d = new_r, new_c, new_d
    return 1000 * (r + 1) + 4 * (c + 1) + list(CLOCKWISE.keys()).index(d)


def test():
    assert first(TEST_PUZZLE) == 6032
    assert first(PUZZLE) == 126350
    assert second(TEST_PUZZLE, TEST_TELEPORT_MAP) == 5031
    # assert second(TEST_PUZZLE, TELEPORT_MAP) == 5031


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE, TELEPORT_MAP))
