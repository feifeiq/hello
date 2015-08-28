pascal''::(Integral a)=>a->a->a
pascal'' _ 1=1
pascal'' x y
	|x==y=1
	|x>0&&y>0&&x>y=pascal'' (x-1) (y-1)+pascal'' (x-1) y
		
pascal'::(Integral a)=>a->[a]
pascal' 1=[1] 
pascal' x=map (pascal'' x) [1..x]


--pascal::(Integral a)=>a->IO()
--pascal x=putStr $ concat $ map ((++ "\n") . show . pascal') [1..x]
pascal :: (Num a, Eq a, Enum a) => a -> IO ()
pascal x=putStr $ concat $ map ((++ "\n") . show . fx) [1..x]
	where fx x=if x==1 then [1] else map (fy x) [1..x] 
		where fy x y=if (x==y||y==1)then 1 else fy (x-1) (y-1)+fy (x-1) y
	
	







