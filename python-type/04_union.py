from typing import Union

def add_prefix_to_id(_id: Union[int, str]) -> str:
    if isinstance(_id, int):
        return "ID-{}".format(_id)
    else:
        return _id
    
print(add_prefix_to_id(10))
print(add_prefix_to_id("ID-99"))