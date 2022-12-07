from __future__ import annotations

import functools
from dataclasses import dataclass
from typing import Optional, List

import common

PUZZLE = common.string(7)

TEST_PUZZLE = """$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir a
29116 f
2557 g
62596 h.lst
$ cd a
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
"""


@dataclass
class Dir:
    name: str
    parent: Optional[Dir]
    children: List[Dir]
    files: List[MyFile]

    def __repr__(self):
        return f"Dir({self.name}, {self.parent.name if self.parent is not None else ''}, {self.children}, {self.files})"


@dataclass
class MyFile:
    size: int
    self: str

    def __repr__(self):
        return f"MyFile({self.size}, {self.name})"


def generate_filetree(puzzle):
    top = Dir("/", None, [], [])
    current = top
    lines = puzzle.split("\n")[1:-1]
    i = 0
    while i < len(lines):
        i, current = process_line(i, lines, current)
        i += 1
    return top


def process_line(i, lines, current):
    if lines[i] == "$ ls":
        while i + 1 < len(lines) and not lines[i + 1].startswith("$"):
            i += 1
            current = process_output(lines[i], current)
    elif lines[i] == "$ cd ..":
        current = current.parent
    else:  # cd somewhere
        dirname = lines[i].split(" ")[-1]
        current = next(c for c in current.children if c.name == dirname)
    return i, current


def process_output(line, current):
    if line.startswith("dir"):
        _, dirname = line.split(" ")
        current.children.append(Dir(dirname, current, [], []))
    else:
        s, fname = line.split(" ")
        current.files.append(MyFile(int(s), fname))
    return current


def get_all_dir_sizes(top):
    def get_dir_sizes(curr, prev=()):
        prev += (dir_size(curr),)
        for child in curr.children:
            prev = get_dir_sizes(child, prev)
        return prev

    return get_dir_sizes(top)


def dir_size(curr):
    child_dir_sizes = sum(dir_size(c) for c in curr.children)
    return sum(f.size for f in curr.files) + child_dir_sizes


@functools.lru_cache(maxsize=None)
def get_tree_and_dir_sizes(puzzle):
    tree = generate_filetree(puzzle)
    return tree, get_all_dir_sizes(tree)


def first(puzzle):
    _, all_dir_sizes = get_tree_and_dir_sizes(puzzle)
    return sum(x for x in all_dir_sizes if x < 100000)


def second(puzzle):
    tree, all_dir_sizes = get_tree_and_dir_sizes(puzzle)
    total = dir_size(tree)
    reqd = 30000000 - (70000000 - total)
    for d in sorted(all_dir_sizes):
        if d > reqd:
            return d


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 95437
    assert first(PUZZLE) == 1513699
    assert second(TEST_PUZZLE) == 24933642
    assert second(PUZZLE) == 7991939


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
