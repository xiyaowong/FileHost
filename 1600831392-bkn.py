Active code page: 65001
def bkn(skey: str):
    hash = 5381
    for i in skey:
        hash += (hash << 5) + ord(i)
    return hash & 2147483647


print(bkn(input('skey: ')))
