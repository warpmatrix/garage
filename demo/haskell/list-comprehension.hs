-- generator && multi-generator
a = [x^2 | x <- [1..5]]
b = [(x, y) | x <- [1..3], y <- [4,5]]
c = [(x, y) | y <- [4,5] ,x <- [1..3]]

-- dependant generators
d = [(x, y) | x <- [1..3], y <- [x..5]]
myconcat xss = [x | xs <- xss, x <- xs]

-- guards
e = [x | x <- [1..10], even x]

-- find primes
factors n = [x | x <- [1..n], n `mod` x == 0]
prime n = factors n == [1,n]
primes n = [x | x <- [2..n], prime x]
-- faster version
allprimes = sieve [2..]
sieve (p:xs) = p : sieve [x | x <- xs, x `mod` p /= 0]

-- sorted
pairs xs = zip xs (tail xs)
sorted xs = and [x <= y | (x,y) <- pairs xs]
-- positions
positions x xs = [loc | (x', loc) <- zip xs [0..], x == x']
-- count
count target xs = length [x | x <- xs, x == target]


-- exercise: pythagorean
pyths n = [(x, y, z) | x <- [1..n], y <- [1..n], z <- [1..n], x^2 + y^2 == z^2]
-- exercise: perfect number
perfect n = n == sum (init (factors n))
perfects n = [x | x <- [1..n], perfect x]
-- exercise: scalar product
sp xs ys = sum [x*y | (x,y) <- zip xs ys]
