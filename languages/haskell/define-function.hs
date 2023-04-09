-- mysignum n = if n < 0 then -1 else
--     if n == 0 then 0 else 1
mysignum n | n < 0 = -1
    | n == 0 = 0
    | otherwise = 1

myabs n | n >= 0 = n
    | otherwise = -n

(&&&) :: Bool -> Bool -> Bool
True &&& b = b
False &&& _ = False

odds n = map f [0..n-1]
    where f x = x * 2 + 1

-- version 1
-- safetail xs = if null xs then []
--     else tail xs
--
-- version 2
-- safetail xs | null xs = []
--     | otherwise = tail xs
--
-- version 3
safetail [] = []
safetail (x : xs) = xs
