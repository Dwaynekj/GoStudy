function listPrime(n)
	for (var i =3; i <= n.length; i++)
		n = rmMult(n, i)
	return n

rmMult(arr, val)
	currMult = val + val
	while arr.find(currMul) != -1
		arr.remove(currMul)
		currMul = currMul + val
	return arr

//Cannot be sparse
// n's size is predetermined