import common

PUZZLE = common.lines(20)


def mix(unique_nums, unique_num_to_num):
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
    mixed = unique_nums
    for _ in range(10):
        mixed = mix(mixed, unique_num_to_num)
    mixed_nums = [unique_num_to_num[n] for n in mixed]
    return sum(mixed_nums[(mixed_nums.index(0) + idx) % len(mixed_nums)] for idx in (1000, 2000, 3000))


def test():
    assert first(TEST_PUZZLE) == 3
    assert first(PUZZLE) != -2910
    assert second(TEST_PUZZLE) == 1623178306


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
