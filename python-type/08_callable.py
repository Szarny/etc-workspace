from typing import Callable, List


def add(a: int, b: int) -> int:
    return a + b


def sub(a: int, b: int) -> int:
    return a - b


def double(a: int) -> int:
    return a ** 2


def call_twice(f: Callable[[int, int], int], g: Callable[[int, int], int], nums: List[int]) -> int:
    return f(g(nums[0], nums[1]), nums[2])


def fourth(f: Callable[[int], int], n) -> int:
    return f(f(n))


print(call_twice(add, sub, [1,2,3]))
print(fourth(double, 2))