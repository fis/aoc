#! /usr/bin/python3

# part 1

def test(n):
    digits = list(str(n))
    double = False
    for i in range(len(digits)-1):
        a, b = digits[i], digits[i+1]
        if b < a:
            return False
        if b == a:
            double = True
    return double

count = 0
for n in range(256310, 732736+1):
    if test(n):
        count += 1
print(count)

# part 2

def test2(n):
    digits = list(str(n))
    double = False
    for i in range(len(digits)-1):
        a, b = digits[i], digits[i+1]
        if b < a:
            return False
        if b == a and (i == 0 or digits[i-1] != a) and (i == len(digits)-2 or digits[i+2] != b):
            double = True
    return double

count = 0
for n in range(256310, 732736+1):
    if test2(n):
        count += 1
print(count)
