-- exercise: and
myand [] = True
myand (b: bs) = b && myand bs

-- exercise: concat
myconcat [] = []
myconcat (xs: xss) = xs ++ myconcat xss

-- exercise: replicate
myreplicate 0 _ = []
myreplicate n x = x : myreplicate (n-1) x

-- exercise: !!
(x: _) !!! 0 = x
(_: xs) !!! idx = xs !!! (idx-1)

-- exercise: insertion sort
insert x [] = [x]
insert x (y: ys) | x <= y = x : y : ys
    | otherwise = y : insert x ys

isort [] = []
isort (x: xs) = insert x (isort xs)

-- exercise: merge sort
merge [] ys = ys
merge xs [] = xs
merge (x:xs) (y:ys) | x <= y = x : merge xs (y:ys)
    | y < x = y : merge (x:xs) ys

halve xs = (take n xs, drop n xs)
    where n = length xs `div` 2

msort [] = []
msort [x] = [x]
msort xs = merge (msort lhs) (msort rhs)
    where (lhs, rhs) = halve xs

qsort [] = []
qsort (x:xs) = qsort lhs ++ [x] ++ qsort rhs
    where 
        lhs = [x' | x' <- xs, x' <= x]
        rhs = [x' | x' <- xs, x' > x]
