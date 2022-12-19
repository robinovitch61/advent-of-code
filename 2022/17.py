import common

PUZZLE = common.line(17)

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
SHAPES = tuple(s[::-1] for s in SHAPES)

TEST_PUZZLE = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"


def show(v):
    for l in v:
        print("|" + format(l, "07b").replace("0", ".").replace("1", "#") + "|")
    print("+-------+\n")


def drop(tower, shape, jets):
    # jets push at least 3 times without stopping
    for _ in range(3):
        jet_dir = next(jets)
        if jet_dir == ">":
            at_right_wall = any(s & 1 != 0 for s in shape)
            if not at_right_wall:
                shape = [l // 2 for l in shape]
        else:
            at_left_wall = any(s & (1 << 6) != 0 for s in shape)
            if not at_left_wall:
                shape = [l * 2 for l in shape]

    stopped = False
    overlap = 0
    while not stopped:
        # jet pushes
        if next(jets) == ">":
            at_right_wall = any(s & 1 for s in shape)
            at_stopped = False
            for i, s in enumerate(shape[:overlap]):
                if tower[len(tower) - overlap + i] & (s // 2) != 0:
                    at_stopped = True
                    break
            if not at_right_wall and not at_stopped:
                shape = [s // 2 for s in shape]
        else:
            at_left_wall = any(s & (1 << 6) for s in shape)
            at_stopped = False
            for i, s in enumerate(shape[:overlap]):
                if tower[len(tower) - overlap + i] & (s * 2) != 0:
                    at_stopped = True
                    break
            if not at_left_wall and not at_stopped:
                shape = [s * 2 for s in shape]

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
    for i, final in enumerate(shape):
        tower[len(tower) - len(shape) - max(0, overlap - len(shape)) + i] |= final
    return tower


class Jets:
    def __init__(self, puzzle):
        self.i = -1
        self.puzzle = puzzle
        self.l = len(puzzle)

    def __next__(self):
        self.i += 1
        return self.puzzle[self.i % self.l]


def build_tower(puzzle, drops):
    tower = []
    jets = Jets(puzzle)
    for drop_num in range(int(drops)):
        drop(tower, SHAPES[drop_num % len(SHAPES)], jets)
    return tower


def first(puzzle):
    return len(build_tower(puzzle, 2022))


def second(puzzle, offset):
    tower = []
    jets = Jets(puzzle)
    drop_num = 0
    first_bit = None
    drops_to_first = None
    while True:
        drop(tower, SHAPES[drop_num % len(SHAPES)], jets)
        drop_num += 1
        reset_motion = (jets.i - offset) % jets.l == 0
        if reset_motion:
            if first_bit is None:
                first_bit = len(tower)
                drops_to_first = drop_num
            else:
                diff = len(tower) - first_bit
                drops_per_diff = drop_num - drops_to_first
                break
    mult = (1000000000000 - drop_num) // drops_per_diff
    drop_num += drops_per_diff * mult
    middle_bit = diff * mult
    while drop_num < 1000000000000:
        drop(tower, SHAPES[drop_num % len(SHAPES)], jets)
        drop_num += 1
    return first_bit + middle_bit + (len(tower) - first_bit)


# `pytest -rfEP 17.py`
def test():
    assert first(TEST_PUZZLE) == 3068
    assert first(PUZZLE) == 3166
    # determined offsets manually
    assert second(TEST_PUZZLE, 0) == 1514285714288
    assert second(PUZZLE, 2) == 1577207977186


if __name__ == "__main__":
    print(first(PUZZLE))
    # determined offset manually
    print(second(PUZZLE, 2))
