

def nth_number(n, input):
    nums = [int(x) for x in input.split(',')]
    prev_turns = {num: [idx + 1] for idx, num in enumerate(nums[:-1])}  # num -> [prev_turn,  prev_prev_turn]
    last_said = nums[-1]
    # print(prev_turns)
    for turn in range(len(nums) + 1, n + 1):
        # print(f"turn: {turn}")
        if last_said not in prev_turns:
            # print(f"  last_said {last_said} is new. This turn is {nums[0]}")
            prev_turns[last_said] = [turn - 1]
            last_said = 0
        else:
            record = prev_turns[last_said]
            # print(f"  last_said {last_said} previously said. Record {record}")
            updated_record = [turn - 1, record[0]]
            prev_turns[last_said] = updated_record
            last_said = updated_record[0] - updated_record[1]
    # print(f"final: {last_said}")
    return last_said



def test_first():
    assert nth_number(2020, "0,3,6") == 436
    assert nth_number(2020, "1,3,2") == 1
    assert nth_number(2020, "2,1,3") == 10
    assert nth_number(2020, "1,2,3") == 27
    assert nth_number(2020, "2,3,1") == 78
    assert nth_number(2020, "3,2,1") == 438
    assert nth_number(2020, "3,1,2") == 1836


def first():
    print(nth_number(2020, "8,13,1,0,18,9"))


def test_second():
    pass
    # assert nth_number(30000000, "0,3,6") == 175594
    # assert nth_number(30000000, "1,3,2") == 2578
    # assert nth_number(30000000, "2,1,3") == 3544142
    # assert nth_number(30000000, "1,2,3") == 261214
    # assert nth_number(30000000, "2,3,1") == 6895259
    # assert nth_number(30000000, "3,2,1") == 18
    # assert nth_number(30000000, "3,1,2") == 362


def second():
    print(nth_number(30000000, "8,13,1,0,18,9"))


if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()