import sys


def filename(day):
    return f"./inputs/{str(day).zfill(2)}.txt"


def string(day):
    with open(filename(day), "r") as f:
        return f.read()
