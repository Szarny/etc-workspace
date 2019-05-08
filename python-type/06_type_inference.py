from typing import Iterable, List

def extract_under_threshold(numbers: Iterable[float], threshold: float) -> List[float]:
    extracted_list : List[float] = []

    for number in numbers:
        if number <= threshold:
            extracted_list.append(number)
    
    return extracted_list

print(extract_under_threshold([1,2,3], 2))