divchk ::(Integral a)=>a->a->Bool
divchk x y
   |x<2=False
   |y<2=False
   |otherwise=x `mod` y ==0 || (divchk x (y-1))
   
isPrime ::(Integral a)=>a->Bool   
isPrime x
	|x<2=False 
	|otherwise=(divchk x (floor (sqrt $ fromIntegral x))) 
 