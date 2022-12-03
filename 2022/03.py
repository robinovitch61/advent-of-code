import common
import string

PUZZLE = common.string(3)

PRIORITIES = {v: i + 1 for i, v in enumerate(string.ascii_lowercase + string.ascii_uppercase)}


def first(puzzle):
    res = 0
    for l in puzzle.split():
        split_idx = len(l) // 2
        first_half = set(l[:split_idx])
        second_half = set(l[split_idx:])
        shared = list(first_half.intersection(second_half))[0]
        res += PRIORITIES[shared]
    return res


def second(puzzle):
    res = 0
    lines = puzzle.split()
    for i in range(0, len(lines), 3):
        group = lines[i: i + 3]
        common = list(set(group[0]).intersection(set(group[1])).intersection(set(group[2])))[0]
        res += PRIORITIES[common]
    return res


TEST_PUZZLE = """vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
"""


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 157
    assert second(TEST_PUZZLE) == 70


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
