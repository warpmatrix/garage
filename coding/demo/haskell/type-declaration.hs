-- type declaration
type MyString = [Char]
a :: MyString
a = "abc"

-- type with parameters
type Pair a = (a,a)
b :: Pair Int
b = (10,10)
-- type can be nested but cannot be recursive
-- type Tree = (Int, [Tree])


-- data declaration with specifying values
data MyBool = True | False

-- constructor
data Shape = Rect Float Float
    | Circle Float
square :: Float -> Shape
square n = Rect n n
-- using function to decompose the constructor
area (Rect x y) = x * y

-- data with parameters
data MyMaybe n = MyNothing | MyJust n
safediv _ 0 = MyNothing
safediv m n = MyJust (m `div` n)

-- recursive types
data Nat = Zero | Succ Nat

nat2int Zero = 0
nat2int (Succ n) = 1 + nat2int n

add Zero n = n
add (Succ m) n = Succ (add m n)


-- exercise

-- mul with natural number
mul Zero n = Zero
mul (Succ m) n = add m (mul m n)

-- folde, eval, size
data Expr = Val Int
    | Add Expr Expr
    | Mul Expr Expr
myfolde id f g (Val n) = id n
myfolde id f g (Add x y) = f (myfolde id f g x) (myfolde id f g y)
myfolde id f g (Mul x y) = g (myfolde id f g x) (myfolde id f g y)

eval = myfolde id (+) (*)
size = myfolde (\_ -> 1) (+) (+)

-- Tree
data Tree a = Leaf a
    | Node (Tree a) (Tree a)
