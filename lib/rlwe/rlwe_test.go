package rlwe_test

import (
	"fmt"
	"github.com/zerjioang/ssscomp/lib/rlwe"
	"testing"
)

func TestRLWE(t *testing.T) {
	fmt.Printf("Init RWLE\n")

	/*Exclusively For Alice*/
	var sAlice [2 * rlwe.M]rlwe.RINGELT /* Alice's Private Key */
	var muAlice [rlwe.Muwords]uint64    /* Alice's recovered mu */

	/*Exclusively For Bob*/
	var muBob [rlwe.Muwords]uint64 /* Bob's version of mu */

	/*Information that gets shared by Alice and Bob*/
	var bAlice [rlwe.M]rlwe.RINGELT /* Alice's Public Key */
	var u [rlwe.M]rlwe.RINGELT      /* Bob's Ring Element from Encapsulation */
	var crV [rlwe.Muwords]uint64    /* Cross Rounding of v */
	/*for i:=0 ; i < 100;i++{
		fmt.Printf("%4d: %d\n",i,RANDOM8())
	}*/
	rlwe.Kem1Generate(&sAlice, &bAlice)
	// KEM1_Generate(s_alice,b_alice)
	for i := 1000; i < 1024; i++ {
		fmt.Printf("%4d: %8d  %8d  %d  %d\n", i, sAlice[i], sAlice[i+1024], bAlice[i], len(sAlice))
	}
	fmt.Printf("Keys initialised\n")
	publicAlice := bAlice[:1024]
	rlwe.Kem1Encapsulate(&u, &crV, &muBob, publicAlice)
	/* 	 for i:=0 ; i < 16;i++{
	   	 	fmt.Printf("%4d: %d %d %dx\n",i,u[i],cr_v[i],mu_bob[i])
	   	 }
	*/
	privateAlice := sAlice[1024:]
	uCopy := u[:]
	rlwe.Kem1Decapsulate(&muAlice, uCopy, privateAlice, crV)
	for i := 0; i < 16; i++ {
		fmt.Printf("%4d: %d %d \n", i, muBob[i], muAlice[i])
	}
}