def filename(day):
    return f"./inputs/{str(day).zfill(2)}.txt"


def string(day):
    with open(filename(day), "r") as f:
        return f.read()


def line(day):
    return string(day).strip().splitlines()[0]


def lines(day):
    return string(day).split("\n")[:-1]  # assumes last line is newline
