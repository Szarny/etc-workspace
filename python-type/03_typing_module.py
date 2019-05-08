from typing import List, Iterable

def greet_all(names: List[str]) -> None:
    for name in names:
        print("Hello, " + name)


def greet_all_iterable(names: Iterable[str]) -> None:
    for name in names:
        print("Hello, " + name)

greet_all(["tsubasa", "foo", "bar"])
# greet_all([1,2,3])
# greet_all(("foo", "bar"))

greet_all_iterable(["tsubasa", "foo", "bar"])
greet_all_iterable(("foo", "bar"))