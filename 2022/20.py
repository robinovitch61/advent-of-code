import common

PUZZLE = common.lines(20)


def first(puzzle):
    # start
    # idxs: 0  1   2  3   4  5  6
    # id:   0  1   2  3   4  5  6
    # nums: 1, 2, -3, 3, -2, 0, 4

    # 1
    # idxs: 0  1   2  3   4  5  6
    # id:   1  0   2  3   4  5  6
    # nums: 2, 1, -3, 3, -2, 0, 4

    # 2
    # idxs: 0  1   2  3   4  5  6
    # id:   0  1,  2  3   4  5  6
    # nums: 1, 2, -3, 3, -2, 0, 4

    # idxs: 0   1  2  3   4  5  6
    # id:   0   2  1  3   4  5  6
    # nums: 1, -3, 2, 3, -2, 0, 4

    # -3
    # idxs:  0  1  2  3   4  5  6
    # id:    2  0  1  3   4  5  6  ##
    # nums: -3, 1, 2, 3, -2, 0, 4

    # idxs: 0  1  2   3  4   5  6
    # id:   0  1  3   4  5   2  6  ##
    # nums: 1, 2, 3, -2, 0, -3, 4

    # idxs: 0  1  2   3   4  5  6
    # id:   0  1  3   4   2  5  6
    # nums: 1, 2, 3, -2, -3, 0, 4

    nums = list(map(int, puzzle))
    id_to_num = {i: n for i, n in enumerate(nums)}  # static
    id_to_idx = {i: i for i, _ in enumerate(nums)}
    idx_to_id = {i: i for i, _ in enumerate(nums)}
    for id, num in list(id_to_num.items())[:4]:
        print(id, num)
        left = num < 0
        for move in range(abs(num)):
            print(move+1)
            if left:
                if id_to_idx[id] == 0:
                    for left_idx in range(len(nums) - 2):
                        right_idx = left_idx + 1
                        idx_to_id[left_idx], id_to_idx[right_idx] = idx_to_id[right_idx], id_to_idx[left_idx]
                    place_idx = len(nums) - 2
                    id_to_idx[id] = place_idx
                    idx_to_id[place_idx] = id
                else:
                    right_idx = id_to_idx[id]
                    left_idx = right_idx - 1
                    right_id = id
                    left_id = idx_to_id[left_idx]
                    id_to_idx[right_id], id_to_idx[left_id] = id_to_idx[left_id], id_to_idx[right_id]
                    idx_to_id[right_idx], idx_to_id[left_idx] = idx_to_id[left_idx], idx_to_id[right_idx]
            else:
                if id_to_idx[id] == len(nums) - 1:
                    # TODO: handle wrap
                    pass
                else:
                    print(id_to_idx)
                    print([f"{idx:2}" for idx in range(len(nums))])
                    print([f"{idx_to_id[idx]:2}" for idx in range(len(nums))])
                    print([f"{id_to_num[idx_to_id[idx]]:2}" for idx in range(len(nums))])
                    print("")
                    left_idx = id_to_idx[id]
                    right_idx = left_idx + 1
                    left_id = id
                    right_id = idx_to_id[right_idx]
                    print(left_idx, right_idx, left_id, right_id)
                    id_to_idx[right_id], id_to_idx[left_id] = id_to_idx[left_id], id_to_idx[right_id]
                    idx_to_id[right_idx], idx_to_id[left_idx] = idx_to_id[left_idx], idx_to_id[right_idx]

    return -1


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == -1
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
