import common

PUZZLE = common.string(6)


def solve(puzzle, n):
    l, r = 0, n
    while r <= len(puzzle):
        if len(set(puzzle[l:r])) == n:
            return r
        l, r = l + 1, r + 1


def first(puzzle):
    return solve(puzzle, 4)


def second(puzzle):
    return solve(puzzle, 14)


# `pytest *`
def test():
    assert first("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 7
    assert first("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5
    assert first("nppdvjthqldpwncqszvftbrmjlhg") == 6
    assert first("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10
    assert first("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11
    assert first(PUZZLE) == 1804

    assert second("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19
    assert second("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23
    assert second("nppdvjthqldpwncqszvftbrmjlhg") == 23
    assert second("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29
    assert second("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26
    assert second(PUZZLE) == 2508


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
