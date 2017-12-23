#!/usr/bin/env python3


def missing_numbers(stdin):
    A, B = ([int(n) for n in line.split()] for line in stdin.readlines()[1::2])
    results = {A.remove(n) if n in A else n for n in B}
    results.remove(None)
    print(' '.join(str(i) for i in sorted(results)))


if __name__ == '__main__':
    import sys
    missing_numbers(sys.stdin)
elif __name__ == '__live_coding__':
    import io

    INPUT = """10
203 204 205 206 207 208 203 204 205 206
13
203 204 204 205 206 207 205 208 203 206 205 206 204 204"""

    missing_numbers(io.StringIO(INPUT))
