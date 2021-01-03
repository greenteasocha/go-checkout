from typing import List

"""
整数i -> 二値変数n個への変換
ex: 
 7 -> [0, 1, 1, 1]  
 13 -> [1, 1, 0, 1] など
大体2進数への変換と同じ
"""


def convert_mask(num: int, n_elements: int) -> List:
    # ビットマスクを使う方法
    res = []
    mask = 1
    for i in range(n_elements):
        if num & mask != 0:
            res.append(1)
        else:
            res.append(0)

        mask <<= 1

    res.reverse()
    return res


def convert_truncate(num: int, n_elements: int) -> List:
    # num自体を右シフト(切り捨て除算)して詰める方法
    res = []
    for i in range(n_elements):
        if num & 1 == 1:
            res.append(1)
        else:
            res.append(0)

        num >>= 1

    res.reverse()
    return res


def convert_base_k(num: int, k: int, n_elements: int) -> List:
    """
    k値変数n個の状態への変換
    ex: 0 or 1 or 2 の変数の組み合わせ(k=3)
     10 -> [1, 0, 1]  10(10) = 101(3)
     26 -> [2, 2, 2]  26(10) = 222(3)
    """
    res = []
    for i in range(n_elements):
        res.append(num % k)
        num //= k

    res.reverse()
    return res


def example1():
    """
    二値変数n個の組み合わせを列挙
    """
    n_elements = 4
    for i in range(2 << n_elements):  # 0 ~ 15まで
        print(convert_mask(i, n_elements))
        # print(convert_truncate(i, n_elements)) # どちらでも


def example2():
    """
    ビットマスクで絞り込みの条件を追加
    """
    n_elements = 4
    condition_mask = int('1010', 2)  # 左から1, 3つ目のフラグが立っている状態のみが列挙される
    for i in range(2 << n_elements):
        if i & condition_mask == condition_mask:
            print(convert_mask(i, n_elements))
        else:
            continue


def example3():
    """
    二値変数ではなくk値変数の場合(k進数と大体同じ)
    """
    base = 3
    n_elements = 4
    for i in range(base ** n_elements):  # 0 ~ 80まで
        print(convert_base_k(i, base, n_elements))


def main():
    example1()
    # example2()
    # example3()


if __name__ == "__main__":
    main()
