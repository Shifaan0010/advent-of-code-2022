import sys


def snafu_to_decimal(snafu_no):
    decimal = 0

    for i, digit in enumerate(snafu_no[::-1]):
        if digit == '-':
            val = -1
        elif digit == '=':
            val = -2
        else:
            val = int(digit)

        decimal += 5 ** i * val

    return decimal


def to_base5(number):
    base5 = ''

    while number > 0:
        base5 = str(number % 5) + base5
        number //= 5

    return base5


def decimal_to_snafu(number):
    base5 = to_base5(number)

    # print(base5)

    snafu = ''

    i = len(base5) - 1
    carry = 0
    while carry > 0 or i >= 0:
    # for i, digit in enumerate(base5[::-1]):
        digit = int(base5[i]) + carry if i >= 0 else carry
            
        carry = digit // 5
        digit %= 5

        # print(digit, carry)
        
        if digit <= 2:
            snafu = str(digit) + snafu
        else:
            carry += 1
            snafu = ('=' if digit == 3 else '-') + snafu
        
        i -= 1

    return snafu


snafu_numbers = [line.strip() for line in sys.stdin]
numbers = [snafu_to_decimal(snafu) for snafu in snafu_numbers]

print(decimal_to_snafu(sum(numbers)))
