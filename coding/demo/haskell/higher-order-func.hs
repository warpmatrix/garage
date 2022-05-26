-- higher-order function: take function as input
twice f x = f (f x)

-- map: version 1
-- mymap f xs = [f x | x <- xs]
-- map: version 2
mymap f [] = []
mymap f (x:xs) = f x : mymap f xs

-- filter: version 1
-- myfilter f xs = [x | x <- xs, f x]
-- filter: version 2
myfilter _ [] = []
myfilter f (x:xs)
    | f x = x : myfilter f xs
    | otherwise = myfilter f xs

-- foldr
myfoldr f v [] = v
myfoldr f v (x:xs) = f x (myfoldr f v xs)
-- length impl by foldr
mylength = myfoldr (\_ x -> 1 + x) 0
-- reverse impl by foldr
myreverse = myfoldr (\x xs -> xs ++ [x]) []

-- .
-- f . g = \x -> f (g x)
mydot f g = \x -> f (g x)

-- all
myall f xs = and [f x | x <- xs]
-- any
myany f xs = or [f x | x <- xs]
-- takeWhile
mytakeWhile f [] = []
mytakeWhile f (x: xs)
    | f x = x : mytakeWhile f xs
    | otherwise = []
-- dropWhile
mydropWhile f [] = []
mydropWhile f (x: xs)
    | f x = mydropWhile f xs
    | otherwise = (x : xs)
