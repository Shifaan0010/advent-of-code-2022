def split_newline(iterable):
    slice = []
    for line in iterable:
        if line == '\n':
            yield slice
            slice = []
        else:
            slice.append(line.strip())


with open('1.txt') as file:
    print(
        sum(
            sorted(
                [sum(map(int, slice))
                    for slice in split_newline(file)],
                reverse=True
            )[:3]
        )
    )
