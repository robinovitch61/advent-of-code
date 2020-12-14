
TEST_INPUT = "939\n7,13,x,x,59,x,31,19"
REAL_INPUT = "1000507\n29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,631,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,383,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17"

def split_input(input):
    num, vals = input.split("\n")
    num = int(num)
    vals = vals.split(",")
    return num, vals


def ignore_x(input):
    num, vals = split_input(input)
    valid = [int(v) for v in vals if v != "x"]
    closest_factors = [(num + val) // val * val for val in valid]
    bus_time = min(closest_factors)
    bus_id = valid[closest_factors.index(bus_time)]
    return (bus_time - num) * bus_id


def include_x(input):
    num, vals = split_input(input)
    id_to_idx = [(int(bus_id), idx) for (idx, bus_id) in enumerate(vals) if bus_id != 'x']
    # https://www.youtube.com/watch?v=z5hR01EmgtM very helpful explanation
    min_t = 0
    can_add = 1
    for id, idx in id_to_idx:
        # print(f"t + {idx} === 0 % {id}")
        while (min_t + idx) % id != 0:
            min_t += can_add
        #  can add the product of all prior ids for the next contraint and still satisfy prior contraints
        can_add *= id
    return min_t


def test_first():
    answer = ignore_x(TEST_INPUT)
    assert(answer == 295)


def first():
    print(ignore_x(REAL_INPUT))


def test_second():
    assert(include_x(TEST_INPUT) == 1068781)


def second():
    print(include_x(REAL_INPUT))

if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()

