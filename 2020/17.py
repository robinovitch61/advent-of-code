
import numpy as np
from scipy.ndimage import convolve

TEST_INPUT = """.#.
..#
###"""

REAL_INPUT = """......##
####.#..
.##....#
.##.#..#
........
.#.#.###
#.##....
####.#.."""

KERNEL_3D = np.ones((3, 3, 3), dtype=int)
KERNEL_3D[1][1][1] = False
KERNEL_4D = np.ones((3, 3, 3, 3), dtype=int)
KERNEL_4D[1][1][1][1] = False


def add_empty_shell(arr):
    ndims = len(arr.shape)
    for idx in range(ndims):
        curr_shape = list(arr.shape)
        curr_shape.pop(idx)
        empty_side = np.expand_dims(np.zeros(curr_shape), axis=idx)
        arr = np.concatenate([empty_side, arr, empty_side], axis=idx)
    return arr


def create_start_arr(input_text, n_extra_dims):
    rows = input_text.split('\n')
    start = np.array([[1 if v == '#' else 0 for v in row] for row in rows])
    for _ in range(n_extra_dims):
        start = np.expand_dims(start, 0)
    return start


def active_after_iter(arr, n_iter, kernel):
    for _ in range(n_iter):
        arr = add_empty_shell(arr)
        adjacent_active = convolve(arr, weights=kernel, mode='constant', cval=0)
        remains_active = (arr == 1) & ((adjacent_active == 2) | (adjacent_active == 3))
        active_to_inactive = ~remains_active
        inactive_to_active = (arr == 0) & (adjacent_active == 3)
        arr[active_to_inactive] = 0
        arr[inactive_to_active] = 1
    return np.sum(arr)


def test_first():
    arr = create_start_arr(TEST_INPUT, 1)
    assert active_after_iter(arr, 6, KERNEL_3D) == 112


def first():
    arr = create_start_arr(REAL_INPUT, 1)
    print(active_after_iter(arr, 6, KERNEL_3D))


def test_second():
    arr = create_start_arr(TEST_INPUT, 2)
    assert active_after_iter(arr, 6, KERNEL_4D) == 848


def second():
    arr = create_start_arr(REAL_INPUT, 2)
    print(active_after_iter(arr, 6, KERNEL_4D))


if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()
