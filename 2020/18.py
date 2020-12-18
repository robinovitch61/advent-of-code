
with open('./18_input.txt', 'r') as f:
    INPUT_TEXT = f.read()


def eval_left_to_right(expr):
    tokens = expr.split(' ')
    while(len(tokens) > 1):
        first = tokens.pop(0)
        op = tokens.pop(0)
        second = tokens.pop(0)
        tokens = [eval(f"{first} {op} {second}")] + tokens
    return int(tokens[0])


def add_first(expr):
    if ('+' not in expr) | ('*' not in expr):
        return eval(expr)
    else:
        tokens = expr.split(' ')
        idx = tokens.index('+')
        first, second = tokens.pop(idx - 1), tokens.pop(idx)
        tokens[idx - 1] = str(eval(f"{first} + {second}"))
        new_expr = ' '.join(tokens)
        return add_first(new_expr)


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


def evaluate(expr, precedence):
    if '(' not in expr:
        return precedence(expr)
    else:
        start = [idx for idx, val in enumerate(list(expr)) if val == '('][-1]  # last left paren in expr
        end = find_closing_paren(expr, start)
        new_expr = expr[:start] + str(precedence(expr[start + 1: end])) + expr[end + 1:]
        return evaluate(new_expr, precedence)



def test_first():
    assert(evaluate("2 * 3 + (4 * 5)", eval_left_to_right) == 26)
    assert(evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)", eval_left_to_right) == 437)
    assert(evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", eval_left_to_right) == 12240)
    assert(evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", eval_left_to_right) == 13632)


def first():
    sum = 0
    for expr in INPUT_TEXT.split('\n'):
        sum += evaluate(expr, eval_left_to_right)
    print(sum)


def test_second():
    assert(add_first("1 + 2 * 3 + 4 * 5 + 6") == 231)
    assert(add_first("1 + 6 + 44") == 51)
    assert(add_first("2 * 3 + 20") == 46)
    assert(evaluate("1 + (2 * 3) + (4 * (5 + 6))", add_first) == 51)
    assert(evaluate("2 * 3 + (4 * 5)", add_first) == 46)
    assert(evaluate("5 + (8 * 3 + 9 + 3 * 4 * 3)", add_first) == 1445)
    assert(evaluate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", add_first) == 669060)
    assert(evaluate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", add_first) == 23340)


def second():
    sum = 0
    for expr in INPUT_TEXT.split('\n'):
        sum += evaluate(expr, add_first)
    print(sum)


if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()
