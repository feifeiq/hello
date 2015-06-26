inc x=x+1
twice x y=x*2+y*2
small x y= if x<y then x else y
conanO'Brien = "It's a-me, Conan O'Brien!"  

maxi'::(Ord a)=>[a]->a
maxi' []=error "empty list"
maxi' [x]=x 
maxi' (x:xs)
	|x>maxTail=x 
	|otherwise=maxTail
	where maxTail=maxi' xs
	
times'::(Num i,Ord i)=>i->a->[a]
times' n x 
	|n<=0=[]
	|otherwise=x:times' (n-1) x 

take'::(Num i,Ord i)=>i->[a]->[a]
take' n _
	|n<=0=[]
take' _ []=[]
take' n (x:y)=x:take' (n-1) y 

reverse'::[a]->[a]
reverse' []=[]
reverse' (x:xs)=reverse' xs++[x]

--repeat'' x=x:repeat' x
-- repeat'::a->a
-- repeat' []=[]
-- repeat' [x]=[x]++repeat' [x]
-- repeat' n=n:repeat' n

in'::(Eq a)=>a->[a]->Bool
in' a []=False
in' a (x:xs)|a==x=True|otherwise=a `in'` xs

len' xs=sum [1|_<-xs]

-- qsort::(Ord a)=>[a]->[a]
-- qsort []=[]
-- qsort (x:xs)=
	-- let	smaller=[a|a<-xs,a<x]
		-- bigger=[a|a<-xs,a>x]
	-- in qsort smaller++[x]++qsort bigger

qsort::(Ord a)=>[a]->[a]
qsort []=[]
qsort (x:xs)=qsort smaller++[x]++qsort bigger
	where	
		smaller=[a|a<-xs,a<x]
		bigger=[a|a<-xs,a>x]


-- isPrime x=[a|a<-[2..y],a>1,x>1,mod x a==0]==[]
	-- where y=(floor (sqrt $ fromIntegral x)) 

		
primes::(Integral a)=>[a]->[a]
primes []=[]
primes (x:xs)=prime x++primes xs
	where prime x=if isPrime x then [x] else []
		where isPrime x=[a|a<-[2..y],a>1,mod x a==0]==[] && x>1
			where y=(floor (sqrt $ fromIntegral x)) 

			
	
		
		

	





