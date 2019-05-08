def greet_with_fun(name: str, is_excite: bool = False) -> None:
    if is_excite:
        print(name + "!!!")
    else:
        print(name + "...")

greet_with_fun("tsubasa", True)