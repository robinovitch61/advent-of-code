import functools

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


class Dir:
    def __init__(self, name, parent, children, files):
        self.name = name
        self.parent = parent
        self.children = children
        self.files = files

    def __repr__(self):
        return f"Dir({self.name}, {self.parent.name if self.parent is not None else 'NONE'}, {self.children}, {self.files})"


class MyFile:
    def __init__(self, size, name):
        self.size = size
        self.name = name

    def __repr__(self):
        return f"MyFile({self.size}, {self.name})"


def generate_filetree(puzzle):
    top = Dir("/", None, [], [])
    current = top
    lines = puzzle.split("\n")[1:-1]
    i = 0
    while i < len(lines):
        if lines[i] == "$ ls":
            while i + 1 < len(lines) and not lines[i + 1].startswith("$"):
                i += 1
                if lines[i].startswith("dir"):
                    _, dirname = lines[i].split(" ")
                    current.children.append(Dir(dirname, current, [], []))
                else:
                    s, fname = lines[i].split(" ")
                    current.files.append(MyFile(int(s), fname))
        elif lines[i].strip() == "$ cd ..":
            current = current.parent
        else:  # cd somewhere
            dirname = lines[i].split(" ")[-1]
            current = [c for c in current.children if c.name == dirname][0]
        i += 1
    return top


def dir_size(curr):
    res = sum(dir_size(c) for c in curr.children)
    if not len(curr.files):
        return res
    return sum([f.size for f in curr.files]) + res


def get_all_dir_sizes(top):
    def get_dir_sizes(curr, prev=()):
        prev += (dir_size(curr),)
        for child in curr.children:
            prev = get_dir_sizes(child, prev)
        return prev

    return get_dir_sizes(top)


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
