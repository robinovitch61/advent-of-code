import re

import common

PUZZLE = common.string(5)

TEST_PUZZLE = """    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
"""


def parse(puzzle):
    stack, moves = puzzle.split("\n\n")

    # parse stacks into lists
    stack_lines = stack.split("\n")
    stacks = []
    for i in range(0, len(stack_lines[-1]), 4):
        new_stack = list(reversed([l[i:i + 4].replace("[", "").replace("]", "").strip() for l in stack_lines[:-1]]))
        new_stack = [s for s in new_stack if s != ""]
        stacks.append(new_stack)

    # parse moves
    parsed_moves = []
    for m in moves.split("\n"):
        parsed = tuple(int(n) for n in re.findall(r"\d+", m))
        if len(parsed) == 3:
            parsed_moves.append(parsed)

    return stacks, parsed_moves


def solve(puzzle, move):
    stacks, moves = parse(puzzle)
    for m in moves:
        stacks = move(stacks, *m)
    return "".join(s[-1] for s in stacks)


def move_first(stacks, cnt, src, dest):
    for _ in range(cnt):
        val = stacks[src - 1].pop()
        stacks[dest - 1].append(val)
    return stacks


def first(puzzle):
    return solve(puzzle, move_first)


def move_second(stacks, cnt, src, dest):
    val = stacks[src - 1][-cnt:]
    stacks[src - 1] = stacks[src - 1][:-cnt]
    stacks[dest - 1] += val
    return stacks


def second(puzzle):
    return solve(puzzle, move_second)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == "CMZ"
    assert second(TEST_PUZZLE) == "MCD"


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
