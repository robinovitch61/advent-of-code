
import numpy as np
from scipy import signal

TEST = 'L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL'

KERNEL = np.array([[1, 1, 1], [1, 0, 1], [1, 1, 1]])

with open("11_input.txt", "r") as f:
    INPUT_LINES = [line.strip() for line in f.readlines()]


def update_seats(occupied, floor):
    occupied_adj = signal.convolve2d((occupied & ~floor), KERNEL, mode='same')
    becomes_occupied = ~floor & ~occupied & (occupied_adj == 0)
    becomes_empty = ~floor & occupied & (occupied_adj >= 4)
    if np.any(becomes_empty) | np.any(becomes_occupied):
        occupied[becomes_empty] = False
        occupied[becomes_occupied] = True
        return True
    else:
        return False


def update_seats_2(occupied, floor):
    def num_visible_adj(occupied, floor, row, col):
        total_visible_adjacent = 0
        height, width = occupied.shape
        for row_step in (-1, 0, 1):
            for col_step in (-1, 0, 1):
                if (row_step == 0) & (col_step == 0):
                    continue
                visible_in_direction = 0
                row_idx, col_idx = row, col
                while True:
                    row_idx, col_idx = row_idx + row_step, col_idx + col_step
                    if (row_idx < 0) | (row_idx >= height) | (col_idx < 0) | (col_idx >= width):
                        break
                    elif floor[row_idx, col_idx]:
                        continue
                    elif occupied[row_idx, col_idx]:
                        visible_in_direction = 1
                        break
                    else:
                        break
                total_visible_adjacent += visible_in_direction
        return total_visible_adjacent

    not_floor = ~floor
    seat_idxs = np.where(not_floor)
    occupied_adj = np.zeros_like(occupied, dtype=int)
    for row, col in zip(*seat_idxs):
        occupied_adj[row, col] = num_visible_adj(occupied, floor, row, col)

    becomes_occupied = not_floor & ~occupied & (occupied_adj == 0)
    becomes_empty = not_floor & occupied & (occupied_adj >= 5)
    if np.any(becomes_empty) | np.any(becomes_occupied):
        occupied[becomes_empty] = False
        occupied[becomes_occupied] = True
        return True
    else:
        return False


def iterate_until_static(data, updater):
    floor = np.array([[seat == '.' for seat in list(row)] for row in data])
    occupied = np.zeros_like(floor, dtype=bool)
    did_update = True
    while did_update:
        did_update = updater(occupied, floor)
    return occupied


def test():
    final = iterate_until_static(TEST.split(), update_seats)
    assert(np.sum(final) == 37)

def first():
    final = iterate_until_static(INPUT_LINES, update_seats)
    print(np.sum(final))

def test2():
    final = iterate_until_static(TEST.split(), update_seats_2)
    assert(np.sum(final) == 26)

def second():
    final = iterate_until_static(INPUT_LINES, update_seats_2)
    print(np.sum(final))

if __name__ == "__main__":
    test()
    first()
    test2()
    second()
