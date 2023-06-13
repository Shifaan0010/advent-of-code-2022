import itertools


def priority(letter):
    return 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'.index(letter) + 1


with open('3.txt', 'r') as file:
    print(
        sum(
            sum(map(
                priority, 
                set.intersection(*(set(line.strip()) for line in group))
            ))
            for group in iter(lambda: list(itertools.islice(file, 3)), [])
        )
    )
