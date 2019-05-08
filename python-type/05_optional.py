from typing import Optional

def greeting(name: Optional[str] = None) -> str:
    if name is None:
        name = "Stranger"
    
    return "Hello, " + name


print(greeting())
print(greeting("tsubasa"))