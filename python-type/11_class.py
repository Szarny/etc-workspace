from typing import ClassVar, List

class Car:

    seats: ClassVar[int]
    passengers : ClassVar[List[str]]

    def __init__(self) -> None:
        self.seats = 4
        self.passengers = []

    def get_seats(self) -> int:
        return self.seats