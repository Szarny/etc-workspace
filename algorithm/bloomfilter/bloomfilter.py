import hashlib

SELECT_ADD    = 1
SELECT_SEARCH = 2
SELECT_REMOVE = 3
SELECT_QUIT   = 4

class Bloomfilter:
    def __init__(self):
        self.table = [0] * (2**16)

    # Use md5, sha256, sha384 hash value to determine
    def get_hashvalue(value):
        md5value    = int(hashlib.md5(value.encode()).hexdigest(), 16) % (2**16)
        sha256value = int(hashlib.sha256(value.encode()).hexdigest(), 16) % (2**16)
        sha384value = int(hashlib.sha384(value.encode()).hexdigest(), 16) % (2**16)

        return md5value, sha256value, sha384value

    def add(self, value):
        md5value, sha256value, sha384value = Bloomfilter.get_hashvalue(value)
        self.table[md5value]    = 1
        self.table[sha256value] = 1
        self.table[sha384value] = 1

    def is_exist(self, value):
        md5value, sha256value, sha384value = Bloomfilter.get_hashvalue(value)
        return (self.table[md5value] and self.table[sha256value] and self.table[sha384value])


class Countingfilter(Bloomfilter):
    # Override
    def add(self, value):
        md5value, sha256value, sha384value = Bloomfilter.get_hashvalue(value)
        self.table[md5value]    += 1
        self.table[sha256value] += 1
        self.table[sha384value] += 1

    
    def remove(self, value):
        md5value, sha256value, sha384value = Bloomfilter.get_hashvalue(value)
        self.table[md5value]    -= 1
        self.table[sha256value] -= 1
        self.table[sha384value] -= 1


def menu():
    while True:
        user_input = input("[?] 1:Add value  2:Search value  3:Remove value  4:Quit >> ")
        if user_input.isdigit() and SELECT_ADD <= int(user_input) <= SELECT_QUIT:
            return int(user_input)


def main():
    user_input = ""
    countingfilter = Countingfilter()

    while True:
        user_input = menu()

        if user_input == SELECT_ADD:
            add_value = input("[?] Value to add >> ")
            countingfilter.add(add_value)
            print("[*] '{}' Successfully added.".format(add_value))

        if user_input == SELECT_SEARCH:
            search_value = input("[?] Value to search >> ")
            
            if countingfilter.is_exist(search_value):
                print("[*] '{}' may be stored.".format(search_value))
            else:
                print("[*] '{}' is NOT stored.".format(search_value))

        if user_input == SELECT_REMOVE:
            remove_value = input("[?] Value to remove >> ")
            
            if countingfilter.is_exist(remove_value):
                countingfilter.remove(remove_value)
                print("[*] '{}' is Successfully removed.".format(remove_value))
            else:
                print("[*] '{}' is NOT stored. Failed to remove.".format(remove_value))

        if user_input == SELECT_QUIT:
            break
        
        print()

    print("[*] Quit.")


if __name__ == '__main__':
    main()