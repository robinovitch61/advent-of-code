import common

PUZZLE = common.lines(20)


def first(puzzle):
    nums = list(map(int, puzzle))
    unique_nums = [i for i, _ in enumerate(nums)]
    unique_num_to_num = {u: n for u, n in zip(unique_nums, nums)}
    for un in range(len(unique_nums)):
        curr_idx = unique_nums.index(un)
        amount_to_move = unique_num_to_num[un]
        if amount_to_move > 0:
            if curr_idx + amount_to_move < len(unique_nums):
                num = unique_nums.pop(curr_idx)
                unique_nums.insert((curr_idx + amount_to_move) % len(unique_nums), num)
            else:
                wrap = (curr_idx + amount_to_move) - len(unique_nums) + 1
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(wrap % len(unique_nums), num)
        else:
            if curr_idx + amount_to_move > 0:
                num = unique_nums.pop(curr_idx)
                unique_nums.insert((curr_idx + amount_to_move) % len(unique_nums), num)
            else:
                wrap = len(unique_nums) + (curr_idx + amount_to_move) - 1
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(wrap % len(unique_nums), num)
    shuffled = [unique_num_to_num[un] for un in unique_nums]
    return sum(shuffled[(shuffled.index(0) + idx) % len(shuffled)] for idx in (1000, 2000, 3000))


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == 3
    assert first(PUZZLE) != -2910
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = [
    "1",
    "2",
    "-3",
    "3",
    "-2",
    "0",
    "4",
]
