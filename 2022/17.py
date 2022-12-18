import common

PUZZLE = common.string(17)

SHAPES = (
    (
        int(b"0011110", 2),
    ),
    (
        int(b"0001000", 2),
        int(b"0011100", 2),
        int(b"0001000", 2),
    ),
    (
        int(b"0000100", 2),
        int(b"0000100", 2),
        int(b"0011100", 2),
    ),
    (
        int(b"0010000", 2),
        int(b"0010000", 2),
        int(b"0010000", 2),
        int(b"0010000", 2),
    ),
    (
        int(b"0011000", 2),
        int(b"0011000", 2),
    ),
)

TEST_PUZZLE = """>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
"""


# def print(a):
#     pass


def show(v):
    return
    for l in v:
        print("|" + format(l, "07b").replace("0", ".").replace("1", "#") + "|")
    print("+-------+")
    print("")


# shape
# ...#... s2
# ..###.. s1
# ...#... s0

# starting tower
# ..####. t0

# intermediate tower
# ....... t3
# ....... t2
# ....... t1
# ..####. t0

# end tower
# ...#... t3
# ..###.. t2
# ...#... t1
# ..####. t0

def drop(tower, shape, jets):
    shape = list(shape[::-1])
    # show(shape[::-1])

    # jets push at least 3 times without stopping
    for _ in range(3):
        jet_dir = next(jets)
        # print(jet_dir)
        if jet_dir == ">":
            at_right_wall = any(s & 1 != 0 for s in shape)
            if not at_right_wall:
                shape = [l // 2 for l in shape]
        else:
            at_left_wall = any(s & (1 << 6) != 0 for s in shape)
            if not at_left_wall:
                shape = [l * 2 for l in shape]

    show(shape[::-1])

    stopped = False
    overlap = 0
    while not stopped:
        # jet pushes
        jet_dir = next(jets)
        print(jet_dir)
        # print(overlap)
        if jet_dir == ">":
            at_right_wall = any(s & 1 for s in shape)
            at_stopped = False
            for i, s in enumerate(shape[:overlap]):
                # show([tower[len(tower) - overlap + i], s])
                if tower[len(tower) - overlap + i] & (s // 2) != 0:
                    at_stopped = True
                    break
            if not at_right_wall and not at_stopped:
                shape = [s // 2 for s in shape]
        else:
            at_left_wall = any(s & (1 << 6) for s in shape)
            at_stopped = False
            for i, s in enumerate(shape[:overlap]):
                # show([tower[len(tower) - overlap + i], s])
                if tower[len(tower) - overlap + i] & (s * 2) != 0:
                    at_stopped = True
                    break
            if not at_left_wall and not at_stopped:
                shape = [s * 2 for s in shape]

        show(shape[::-1])
        show(tower[::-1])
        # drop
        overlap += 1

        if not len(tower) or overlap == len(tower) + 1:
            overlap -= 1
            stopped = True
        else:
            for t, s in zip(tower[len(tower) - overlap:], shape[:overlap]):
                if t & s != 0:
                    overlap -= 1
                    stopped = True
                    break

    # add shape to tower
    for _ in range(len(shape) - overlap):
        tower.append(0)
    # print(overlap)
    # show(tower[::-1])
    for i, final in enumerate(shape):
        # print((len(tower) - 1) - (len(shape) - 1) - max(0, overlap - len(shape)) + i)
        tower[(len(tower) - 1) - (len(shape) - 1) - max(0, overlap - len(shape)) + i] |= final
    return tower


def first(puzzle):
    puzzle = puzzle.strip()
    tower = []

    def jet_gen():
        i = -1
        l = len(puzzle)
        while True:
            i += 1
            yield puzzle[i % l]

    jets = jet_gen()
    for drop_num in range(2022):
    #     print(drop_num)
    # for drop_num in range(4):
        drop(tower, SHAPES[drop_num % len(SHAPES)], jets)
        show(tower[::-1])

    return len(tower)


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 3068
    assert first(PUZZLE) > 3164
    # assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    # print(second(PUZZLE))
