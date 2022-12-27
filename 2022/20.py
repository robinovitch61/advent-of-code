import common

PUZZLE = common.lines(20)


def mix(unique_nums, unique_num_to_num):
    len_nums = len(unique_nums)
    for un in range(len_nums):
        curr_idx = unique_nums.index(un)
        amount_to_move = unique_num_to_num[un]
        if amount_to_move > 0:
            if curr_idx + amount_to_move < len_nums:
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(curr_idx + amount_to_move, num)
            else:
                wrap = (curr_idx + amount_to_move) - len_nums + 1
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(wrap % (len_nums - 1), num)
        else:
            if curr_idx + amount_to_move > 0:
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(curr_idx + amount_to_move, num)
            else:
                wrap = len_nums + (curr_idx + amount_to_move) - 1
                num = unique_nums.pop(curr_idx)
                unique_nums.insert(wrap % (len_nums - 1), num)
    return unique_nums


def first(puzzle):
    nums = list(map(int, puzzle))
    unique_nums = [i for i, _ in enumerate(nums)]
    unique_num_to_num = {u: n for u, n in zip(unique_nums, nums)}
    mixed_unique_nums = mix(unique_nums, unique_num_to_num)
    mixed_nums = [unique_num_to_num[n] for n in mixed_unique_nums]
    return sum(mixed_nums[(mixed_nums.index(0) + idx) % len(mixed_nums)] for idx in (1000, 2000, 3000))


def second(puzzle):
    nums = list(map(int, puzzle))
    nums = [n * 811589153 for n in nums]
    unique_nums = [i for i, _ in enumerate(nums)]
    unique_num_to_num = {u: n for u, n in zip(unique_nums, nums)}
    for _ in range(10):
        unique_nums = mix(unique_nums, unique_num_to_num)
    mixed_nums = [unique_num_to_num[n] for n in unique_nums]
    return sum(mixed_nums[(mixed_nums.index(0) + idx) % len(mixed_nums)] for idx in (1000, 2000, 3000))


def test():
    assert first(TEST_PUZZLE) == 3
    assert first(PUZZLE) == 5498
    assert second(TEST_PUZZLE) == 1623178306
    assert second(PUZZLE) == 3390007892081


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
