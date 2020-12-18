
with open('./18_input.txt', 'r') as f:
    INPUT_TEXT = f.read()


def eval_left_to_right(expr):
    tokens = [v for v in expr.split(' ') if v != '']
    while(len(tokens) > 1):
        first = tokens.pop(0)
        op = tokens.pop(0)
        second = tokens.pop(0)
        tokens = [eval(f"{first} {op} {second}")] + tokens
    print(f"computed: {int(tokens[0])}")
    return int(tokens[0])


def find_closing_paren(expr, start_idx):
    count = 1
    while count > 0:
        start_idx += 1
        val = expr[start_idx]
        if val == "(":
            count += 1
        if val == ")":
            count -= 1
    return start_idx


def reduce_expr(expr):
    if '(' not in expr:
        return eval_left_to_right(expr)

    # first_left_paren = expr.index('(')
    close = find_closing_paren(expr, first_left_paren)
    if first_left_paren == 0:
        prefix, prefix_op = None, None
    else:
        prefix = expr[:first_left_paren - 3]
        prefix_op = expr[first_left_paren - 2]

    in_parens = expr[first_left_paren + 1: close]

    if close == len(expr) - 1:
        postfix, postfix_op = None, None
    else:
        postfix = expr[close + 4:]
        postfix_op = expr[close + 2]

    print(f"expr: {expr}")
    print(f"prefix: {prefix}")
    print(f"prefix_op: {prefix_op}")
    print(f"in_parens: {in_parens}")
    print(f"postfix_op: {postfix_op}")
    print(f"postfix: {postfix}")
    final_expr = ''
    if prefix:
        final_expr += f"{reduce_expr(prefix)} {prefix_op} "
    final_expr += f"{reduce_expr(in_parens)}"
    if postfix:
        final_expr += f" {postfix_op} {reduce_expr(postfix)}"

    print(f"final_expr: {final_expr}")

    return reduce_expr(final_expr)


def evaluate(expr):
    return reduce_expr(expr)


def test_first():
    # assert(evaluate("2 * 3 + (4 * 5)") == 26)
    # assert(evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)") == 437)
    # assert(evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))") == 12240)
    assert(evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2") == 13632)


def first():
    sum = 0
    for expr in INPUT_TEXT.split('\n'):
        sum += evaluate(expr)
    print(sum)


def test_second():
    pass


def second():
    pass


if __name__ == "__main__":
    test_first()
    # first()
    # test_second()
    # second()
