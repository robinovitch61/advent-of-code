
with open('./16_input.txt', 'r') as f:
    INPUT_TEXT = f.read()
with open('./16_test_input.txt', 'r') as f:
    TEST_INPUT_TEXT = f.read()
with open('./16_test_input2.txt', 'r') as f:
    TEST_INPUT_2_TEXT = f.read()


def gen_valid_field_ranges(ranges):
    field_ranges = {}
    all_set = set()
    for this_range in ranges:
        name, or_vals = this_range.split(': ')
        first, second = or_vals.split(' or ')
        valid = set()
        for val in (first, second):
            lower, upper = val.split('-')
            valid = valid.union(set(range(int(lower), int(upper) + 1)))
        field_ranges[name] = valid
        all_set = all_set.union(valid)
    return field_ranges, all_set


def parse_input(text_input):
    fields = text_input.split('\n\n')
    raw_ranges, ticket, other_tickets = fields[0].split('\n'), fields[1].split(':\n')[1], fields[2].split(':\n')[1]
    ticket = [int(v) for v in ticket.split(',')]
    parsed_tickets = []
    for other_ticket in other_tickets.split('\n'):
        parsed_tickets.append([int(v) for v in other_ticket.split(',')])
    return raw_ranges, ticket, parsed_tickets


def compute_scan_err_rate(ranges, other_tickets):
    _, valid = gen_valid_field_ranges(ranges)
    scan_err_rate = 0
    for other_ticket in other_tickets:
        for val in other_ticket:
            if val not in valid:
                scan_err_rate += val
    return scan_err_rate


def test_first():
    ranges, _, other_tickets = parse_input(TEST_INPUT_TEXT)
    scan_err_rate = compute_scan_err_rate(ranges, other_tickets)
    assert scan_err_rate == 71


def first():
    ranges, _, other_tickets = parse_input(INPUT_TEXT)
    print(compute_scan_err_rate(ranges, other_tickets))


def get_valid_tickets(other_tickets, all_set):
    valid_tickets = []
    for other_ticket in other_tickets:
        valid_tickets.append(other_ticket)
        for val in other_ticket:
            if val not in all_set:
                valid_tickets.pop()
                break
    return valid_tickets


def possible_fields(field_ranges, valid_tickets):
    options = {idx: set(field_ranges.keys()) for idx in range(len(field_ranges))}
    for idx in range(len(field_ranges)):
        valid_ticket_vals = set([t[idx] for t in valid_tickets])
        for field, valid_set in field_ranges.items():
            if valid_ticket_vals.intersection(valid_set) != valid_ticket_vals:
                options[idx].remove(field)
    return options


def eliminate_options(options):
    final = {}
    while len(final) < len(options):
        for idx, fields in options.items():
            if len(fields) == 1:
                field = fields.pop()
                final[field] = idx
                for _, fields in options.items():
                    if field in fields:
                        fields.remove(field)
    return final


def test_second():
    ranges, ticket, other_tickets = parse_input(TEST_INPUT_2_TEXT)
    field_ranges, all_set = gen_valid_field_ranges(ranges)
    valid_tickets = get_valid_tickets(other_tickets, all_set)
    options = possible_fields(field_ranges, valid_tickets)
    final = eliminate_options(options)
    assert (final['row'] == 0) & (final['class'] == 1) & (final['seat'] == 2)


def second():
    ranges, ticket, other_tickets = parse_input(INPUT_TEXT)
    field_ranges, all_set = gen_valid_field_ranges(ranges)
    valid_tickets = get_valid_tickets(other_tickets, all_set)
    options = possible_fields(field_ranges, valid_tickets)
    final = eliminate_options(options)
    answer = 1
    for field, idx in final.items():
        if "departure" in field:
            answer *= ticket[idx]
    print(answer)


if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()
