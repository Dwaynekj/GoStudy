//Recursive Cartesian GO

function cartProd(x, y){
	i = len (x)
	j = len(y)
	result = []

	function A(i, j, result){
	// Inner  "iteration on i"
		// Inner  "iteration on j"
		function B(i, j, r) {

			if (i < 0) {
				return r
			}

			r.push(  [x[i], y[j--]])
			
			return B(i, j, r)
		}

		if (i < 0) {
			return result
		} else {
			return A(B(i--, j, result))
		}
		
	}

	return A(i, j, result)

	//Watch out for stack overflow
	// Variable shadowing
}