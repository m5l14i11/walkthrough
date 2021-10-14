# python3

def prime(n):
    is_prime = [1] * n
    is_prime[0] = 0
    is_prime[1] = 0

    for i in range(2, n):
        if is_prime[i] == 1:
            for j in range(i ** 2, n, i):
                is_prime[j] = 0

    return is_prime

def get_nearest(n, is_prime):
    if is_prime[n] == 1:
        return n

    res = n
    max_dist = float('inf')
    
    for i in reversed(range(len(is_prime))):
        if is_prime[i] == 1:
            d = abs(n - i)

            if d < max_dist:
                max_dist = d
                res = i

    return res
    
if __name__ == '__main__':
    n = int(input())

    is_prime = prime(n * 2)

    print(get_nearest(n, is_prime))
