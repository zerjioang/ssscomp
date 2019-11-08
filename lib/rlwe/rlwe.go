// this file implements the RLWE ( Ring learning with errors)
// lattice based encryption, the algorithm as of writing the source seems to be
// safe from quantum computer attacks
// this file implements a non-constant time gaussian base RLWE with params n = 1024 q 40961
// the above parameters have been proved to deliver minimal 2^256 bit security, upper limit not yet found

package rlwe

//import "math/rand"

/* below starts the code to compute  Number Theoretic Transform f0r multiplicattion in ring F_q[x] / <x^n+1>
n = 1024, q = 40961	*/
var n RINGELT = 1024
var q RINGELT = 40961

/*
We use Gentleman-Sande, decimation-in-frequency FFT, for the forward FFT.
We premultiply x by the 2n'th roots of unity to affect a Discrete Weighted Fourier Transform,
so when we apply pointwise multiplication we obtain the negacyclic convolution, i.e. multiplication
modulo x^n+1.
Note that we will not perform the usual scambling / bit-reversal procedure here because we will invert
the fourier transform using decimation-in-time.
*/
func FftTwistedForward102440961(x []RINGELT) {
	var index, step RINGELT
	var i, j, m RINGELT
	var t0, t1 RINGELT

	//Pre multiplication for twisted FFT
	j = 0
	for i = 0; i < n>>1; i++ {
		//MUL_MOD(x[j], x[j], W[i], q);
		x[j] = RINGELT((FFTLONG(x[j]) * FFTLONG(W[i])) % FFTLONG(q))
		j++
		//MUL_MOD(x[j], x[j], W_sqrt[i], q);
		x[j] = RINGELT((FFTLONG(x[j]) * FFTLONG(WSqrt[i])) % FFTLONG(q))
		j++
	}

	step = 1
	for m = n >> 1; m >= 1; m = m >> 1 {
		index = 0
		for j = 0; j < m; j++ {
			for i = j; i < n; i += m << 1 {
				//ADD_MOD(t0, x[i], x[i+m], q);
				t0 = x[i] + x[i+m]
				if t0 >= q {
					t0 = t0 - q
				}

				//ADD(t1, x[i], q - x[i+m]);
				t1 = x[i] + (q - x[i+m])

				//MUL_MOD(x[i+m], t1, W[index], q);
				x[i+m] = RINGELT((FFTLONG(t1) * FFTLONG(W[index])) % FFTLONG(q))
				x[i] = t0
			}
			//SUB_MODn(index, index, step);
			index = index + (n - (step))
			if index >= n {
				index = index - n
			}

		}
		step = step << 1
	}
}

/*
We use Cooley-Tukey, decimation-in-time FFT, for the inverse FFT.
We postmultiply x by the inverse of the 2n'th roots of unity * n^-1 to affect a Discrete Weighted Fourier Transform,
so when we apply pointwise multiplication we obtain the negacyclic convolution, i.e. multiplication
modulo x^n+1.
Note that we will not perform the usual scambling / bit-reversal procedure here because we will the forward
fourier transform is using decimation-in-frequency.
*/
func FftTwistedBackward102440961(x []RINGELT) {
	var index, step RINGELT
	var i, j, m RINGELT
	var t0, t1 RINGELT

	step = n >> 1
	for m = 1; m < n; m = m << 1 {
		index = 0
		for j = 0; j < m; j++ {
			for i = j; i < n; i += m << 1 {
				t0 = x[i]
				//t0 -= (t0 >= q) ? q : 0;
				if t0 >= q { // do the above step with a branch, this causes  us a branch miss
					t0 = t0 - q
				}
				//MUL_MOD(t1, x[i+m], W_rev[index], q);
				t1 = RINGELT((FFTLONG(x[i+m]) * FFTLONG(WRev[index])) % FFTLONG(q))
				//ADD(x[i], t0, t1);
				x[i] = t0 + t1
				//ADD(x[i+m], t0, q - t1);
				x[i+m] = (t0) + (q - t1)
			}
			//SUB_MODn(index, index, step);
			index = index + (n - (step))
			if index >= n {
				index = index - n
			}

		}
		step = step >> 1
	}

	//Post multiplication for twisted FFT
	j = 0
	for i = 0; i < n>>1; i++ {
		//MUL_MOD(x[j], x[j], W_rev[i], q);
		x[j] = RINGELT((FFTLONG(x[j]) * FFTLONG(WRev[i])) % FFTLONG(q))
		j++
		//MUL_MOD(x[j], x[j], W_sqrt_rev[i], q);
		x[j] = RINGELT((FFTLONG(x[j]) * FFTLONG(WSqrtRev[i])) % FFTLONG(q))
		j++
	}
}

func FftForward(x []RINGELT) {
	FftTwistedForward102440961(x)
}

func FftBackward(x []RINGELT) {
	FftTwistedBackward102440961(x)
	for i := RINGELT(0); i < m; i++ {
		//MUL_MOD((_x)[_i], (_x)[_i], 40921, (q));\
		x[i] = RINGELT((FFTLONG(x[i]) * FFTLONG(40921)) % FFTLONG(q))
	}
}

/* Auxiliary functions for constant-time comparison
   these should be replaced with faster versions also teh system is partiallly constant time
   so as leaks are minimal and unpredictable
*/

/*
 * Returns 1 if x != 0
 * Returns 0 if x == 0
 * x and y are arbitrary unsigned 64-bit integers
 */
func ctIsnonzeroU64(x uint64) uint64 {
	return (x | -x) >> 63
}

/*
 * Returns 1 if x != y
 * Returns 0 if x == y
 * x and y are arbitrary unsigned 64-bit integers
 */
func ctNeU64(x, y uint64) uint64 {
	return ((x - y) | (y - x)) >> 63
}

/* Returns 1 if x < y
 * Returns 0 if x >= y
 * x and y are arbitrary unsigned 64-bit integers
 */
func ctLtU64(x, y uint64) uint64 {
	return (x ^ ((x ^ y) | ((x - y) ^ y))) >> 63
}

/* Returns 0xFFFF..FFFF if bit != 0
 * Returns            0 if bit == 0
 */
func ctMaskU64(bit uint64) uint64 {
	return 0 - ctIsnonzeroU64(bit)
}

/* Returns 0 if a >= b
 * Returns 1 if a < b
 * Where a and b are both 3-limb 64-bit integers.
 * This function runs in constant time.
 */
func cmpltCt(a, b [3]uint64) int {

	r := uint64(0)  /* result */
	_m := uint64(0) /* mask   */
	i := int(0)
	for i = 2; i >= 0; i-- {
		r |= ctLtU64(a[i], b[i]) & ^_m       // bit wise complement operator in go is ^
		_m |= ctMaskU64(ctNeU64(a[i], b[i])) /* stop when a[i] != b[i] */
	}
	return int(r & 1)
}

func singleSample(in [3]uint64) uint32 {
	lowerIndex := uint32(0)
	thisIndex := uint32(32)
	upperIndex := uint32(64)
	i := int(0)
	for i = 0; i <= 6; i++ {
		//fmt.Printf("%d : %d\n",i,rlwe_table[this_index][0])
		if cmpltCt(in, rlweTable[thisIndex]) == 1 {
			upperIndex = thisIndex
		} else {
			lowerIndex = thisIndex
		}
		thisIndex = (lowerIndex + upperIndex) / 2
	}
	return upperIndex
}

/*v = e0*b+e1, multiply and add in the ring. All done in the FFT / CRT domain, so point-wise multiplication and addition*/
/*#define POINTWISE_MUL_ADD(v, b, e0, e1) \
do {\
	for (uint16_t _i = 0; _i < m; ++_i) {\
		MUL_MOD((v)[_i], (e0)[_i], (b)[_i], (q));\
		ADD_MOD((v)[_i], (v)[_i], (e1)[_i], (q));\
	}\
} while(0)
POINTWISE_MUL_ADD(b, a, s+m, s);
*/
/*v = e0*b+e1, multiply and add in the ring. All done in the FFT / CRT domain, so point-wise multiplication and addition
 */
func PointwiseMulAdd(noise, private, public []RINGELT) {
	for i := RINGELT(0); i < m; i++ {
		//MUL_MOD(x[j], x[j], W_rev[i], q);
		public[i] = RINGELT((FFTLONG(private[i]) * FFTLONG(a[i])) % FFTLONG(q))

		//ADD_MOD(t0, x[i], x[i+m], q);
		public[i] = public[i] + noise[i]
		if public[i] >= q {
			public[i] = public[i] - q
		}
	}
}

/*v = e0*b, multiply and add in the ring. All done in the FFT / CRT domain, so point-wise multiplication and addition*/
func PointwiseMul(v, b, e0 []RINGELT) {
	for i := RINGELT(0); i < m; i++ {
		v[i] = RINGELT((FFTLONG(e0[i]) * FFTLONG(b[i])) % FFTLONG(q))
	}

}

/*v = e0+b, multiply and add in the ring. All done in the FFT / CRT domain, so point-wise multiplication and addition*/
func PointwiseAdd(v, b, e0 []RINGELT) {
	for i := RINGELT(0); i < m; i++ {
		v[i] = e0[i] + b[i]
		if v[i] >= q {
			v[i] = v[i] - q
		}
	}

}

/*Set the m'th coefficient to be 0 in the prime case*/
func sampleSecret(s []RINGELT) {

	r := uint64(0)
	var rnd [3]uint64
	ind := RINGELT(0)
	//#if MISPOWEROFTWO
	for ind < m {
		//#else
		//	while (ind < m-1) {
		//#endif
		if (ind & 0x3F) == 0 {
			//r = RANDOM64
			r = RandomSafeUint64()
			//r = uint64(rand.Uint32()) << 32 | uint64(rand.Uint32())

			//fmt.Printf("r = %d\n",r)
		}
		//RANDOM192(rnd);
		rnd[0] = RandomSafeUint64() //uint64(rand.Uint32()) << 32 | uint64(rand.Uint32())
		rnd[1] = RandomSafeUint64() //uint64(rand.Uint32()) << 32 | uint64(rand.Uint32())
		rnd[2] = RandomSafeUint64() //uint64(rand.Uint32()) << 32 | uint64(rand.Uint32())

		s[ind] = RINGELT(singleSample(rnd))

		if s[ind] > 0 {
			if (r & 1) == 1 {
				//fmt.Printf("%d : %d %d \n",ind,s[ind], RINGELT(q) - s[ind])
				s[ind] = RINGELT(q) - s[ind]
			}
		}
		r >>= 1
		ind++
	}
	//#if !MISPOWEROFTWO
	//	s[m-1] = 0;
	//#endif
}

/* Round and cross-round */
func roundAndCrossRound(modularRnd *[muwords]uint64, crossRnd *[muwords]uint64, v []RINGELT) {
	//RANDOM_VARS;
	i := RINGELT(0)
	r := RandomSafeUint64()
	var word, pos, rbit, val RINGELT = 0, 0, 0, 0

	//memset((void *) modular_rnd, 0, muwords*sizeof(uint64_t));
	//memset((void *) cross_rnd, 0, muwords*sizeof(uint64_t));

	for i = 0; i < m; i++ {
		val = v[i]
		/*Randomize rounding procedure - probabilistic nudge*/
		if qmod4 == 1 {
			if val == 0 {
				if (r & 1) == 1 {
					val = q - 1
				}
				rbit++
				if rbit >= 64 {
					r = RandomSafeUint64()
					rbit = 0
				} else {
					r = r >> 1
				}
			} else if val == q_1_4-1 {
				if (r & 1) == 1 {
					val = q_1_4
				}
				rbit++
				if rbit >= 64 {
					r = RandomSafeUint64()
					rbit = 0
				} else {
					r = r >> 1
				}
			}
		} else {
			if val == 0 {
				if (r & 1) == 1 {
					val = q - 1
				}
				rbit++
				if rbit >= 64 {
					r = RandomSafeUint64()
					rbit = 0
				} else {
					r = r >> 1
				}
			} else if val == q_3_4-1 {
				if (r & 1) == 1 {
					val = q_3_4
				}
				rbit++
				if rbit >= 64 {
					r = RandomSafeUint64()
					rbit = 0
				} else {
					r = r >> 1
				}
			}
		}

		/*Modular rounding process*/
		if val > q_1_4 && val < q_3_4 {
			modularRnd[word] |= uint64(1) << pos
		}

		/*Cross Rounding process*/
		if (val > q_1_4 && val <= q_2_4) || val >= q_3_4 {
			crossRnd[word] |= uint64(1) << pos
		}

		pos++
		if pos == 64 {
			word++
			pos = 0
		}

	}

}

/* Reconcile */
func rec(r *[muwords]uint64, w []RINGELT, b [muwords]uint64) {
	i := RINGELT(0)
	word := RINGELT(0)
	pos := RINGELT(0)

	//memset((void *) r, 0, muwords*sizeof(uint64_t));

	for i = 0; i < m; i++ {
		if ((b[word] >> pos) & 1) == 1 {
			if w[i] > r1_l && w[i] < r1_u {
				r[word] |= uint64(1) << pos
			}
		} else {
			if w[i] > r0_l && w[i] < r0_u {
				r[word] |= uint64(1) << pos
			}
		}
		pos++
		if pos == 64 {
			word++
			pos = 0
		}
	}
}

/* Construct Alice's private / public key pair. Return all elements in the Fourier Domain
 * input:  none
 * output: private key s_1=s[n]...s[2*n-1] in Fourier Domain
 *         noise term s_0=s[0]...s[n-1] in Fourier Domain, not needed again
 *         public key b in Fourier Domain
 */
func Kem1Generate(s *[2 * m]RINGELT, b *[m]RINGELT) {

	// these arrays just refer to underlying slices
	noise := s[:1024]
	private := s[1024:]
	public := b[:1024]

	sampleSecret(noise)
	sampleSecret(private)
	FftForward(noise)
	FftForward(private)

	// generate public key in public
	PointwiseMulAdd(noise, private, public)

}

/* Encapsulation routine. Returns an element in R_q x R_2
 * input:  Alice's public key b in Fourier Domain
 * output: Bob's public key u in Fourier Domain
 *         reconciliation data cr_v
 *         shared secret mu
 */

func Kem1Encapsulate(u *[m]RINGELT, crV *[muwords]uint64, mu *[muwords]uint64, b []RINGELT) {
	//  create bob's noise, private key, public key
	var noise = make([]RINGELT, 1024, 1024)
	var private = make([]RINGELT, 1024, 1024)
	//var public = make([]RINGELT,1024,1024);
	public := u[:]
	var v = make([]RINGELT, 1024, 1024)

	var e2 = make([]RINGELT, 1024, 1024)

	sampleSecret(noise)
	sampleSecret(private)
	FftForward(noise)
	FftForward(private)

	sampleSecret(e2)

	// generate public key in public
	// remember bob's public key is differently generated
	// alice's public key is generated by  POINTWISE_MUL_ADD(noise,private,public)
	// however bob's public is generated using  POINTWISE_MUL_ADD(private,noise,public)
	// note that private and noise are interchanged
	// verify the crypo more throughly as there might be inherent weakness
	// however, this might be the basis of the strength also
	PointwiseMulAdd(private, noise, public)

	/*	for i:=1000 ; i < 1024;i++{
		 fmt.Printf("bob %4d: %8d  %8d  %d \n",i,noise[i],private[i], public[i])
		}
	*/

	//copy(u, public[:] )
	/*for i:=0 ; i < 1024;i++{
		u[i]= public[i]
	}*/

	PointwiseMul(v, b, noise)

	FftBackward(v)
	//FFT_twisted_backward_1024_40961(v)

	//MAPTOCYCLOTOMIC(v);   since we are already power of 2, it's not required and is thus a NIL operation

	PointwiseAdd(v, v, e2)

	roundAndCrossRound(mu, crV, v)

	/*
	   	for i:=0 ; i < 16;i++{
	    	 	fmt.Printf("%4d: %d %d   %d\n",i,private[i], mu[i],cr_v[i])
	    	 }


	    	 {
	    	 	total := uint64(0)
	    	 	for i:=0 ; i < 1024;i++{
	    	 		total = total + uint64(v[i])

	    	 }
	    	 fmt.Printf("Total %d\n",total)

	    	 }
	*/

}

/* Decapsulation routine.
 * input:  Bob's public key u in Fourier Domain
 *         Alice's private key s_1 in Fourier Domain
 *         reconciliation data cr_v
 * output: shared secret mu

void KEM1_Decapsulate(uint64_t mu[muwords], RINGELT u[m], RINGELT s_1[m], uint64_t cr_v[muwords]) {
	RINGELT w[m];

	POINTWISE_MUL(w, s_1, u); //Create w = s1*u
	FFT_backward(w); //Undo the Fourier Transform
	MAPTOCYCLOTOMIC(w);
	rec(mu, w, cr_v);
}
*/

func Kem1Decapsulate(mu *[muwords]uint64, u []RINGELT, s1 []RINGELT, crV [muwords]uint64) {
	w := make([]RINGELT, m, m)
	PointwiseMul(w, s1, u)
	FftBackward(w)

	rec(mu, w, crV)

	/* {
		total := uint64(0)
		for i:=0 ; i < 1024;i++{
			total = total + uint64(w[i])

	}
	fmt.Printf("Total w %d\n",total)

	}*/

}
