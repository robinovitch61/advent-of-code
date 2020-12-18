
with open('./18_input.txt', 'r') as f:
    INPUT_TEXT = f.read()


def eval_left_to_right(expr):
    tokens = [v for v in expr.split(' ') if v != '']
    while(len(tokens) > 1):
        first = tokens.pop(0)
        op = tokens.pop(0)
        second = tokens.pop(0)
        tokens = [eval(f"{first} {op} {second}")] + tokens
    return int(tokens[0])


def find_closing_paran(expr, start_idx):
    count = 1
    while count > 0:
        start_idx += 1
        val = expr[start_idx]
        if val == "(":
            count += 1
        if val == ")":
            count -= 1
    return start_idx


def remove_redundant_parens(expr):
    if '(' not in expr:
        return expr
    start_idx = expr.index('(')
    if start_idx == 0:
        close_idx = find_closing_paran(expr, start_idx)
        new_expr = expr[1:close_idx] + expr[close_idx+1:]
        return remove_redundant_parens(new_expr)
    else:
        return expr


def evaluate(expr):
    expr = remove_redundant_parens(expr)
    start_idxs = [idx for idx, val in enumerate(list(expr)) if val == "("]
    close_idxs = [find_closing_paran(expr, idx) for idx in start_idxs]
    if len(start_idxs) == 0:
        return eval_left_to_right(expr)
    else:
        start, close = start_idxs.pop(0), close_idxs.pop(0)
        return eval_left_to_right(expr[:start] + str(evaluate(expr[start+1:close])) + expr[close+1:])


def test_first():
    assert(evaluate("2 * 3 + (4 * 5)") == 26)
    assert(evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)") == 437)
    assert(evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))") == 12240)
    assert(evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2") == 13632)


def first():
    sum = 0
    for expr in INPUT_TEXT.split('\n'):
        print(expr)
        sum += evaluate(expr)
    print(sum)


def test_second():
    pass


def second():
    pass


if __name__ == "__main__":
    test_first()
    first()
    # test_second()
    # second()
