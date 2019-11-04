# `Secret Sharing` & `Secure Computation` Library for Golang

<p align="center">
  <img alt="ssscomp" src="images/header.png" width="auto"></img>
  <h3 align="center"><b>Secret Sharing & Secure Computation Library for Go</b></h3>
</p>

<p align="center">
    <a href="https://travis-ci.org/zerjioang/ssscomp">
      <img alt="Build Status" src="https://travis-ci.org/zerjioang/ssscomp.svg?branch=master">
    </a>
    <a href="https://goreportcard.com/report/github.com/zerjioang/ssscomp">
       <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/zerjioang/ssscomp">
    </a>
    <a href="https://github.com/zerjioang/ssscomp/blob/master/LICENSE">
        <img alt="Software License" src="http://img.shields.io/:license-GPL3-brightgreen.svg?style=flat-square">
    </a>
    <a href="https://godoc.org/github.com/zerjioang/ssscomp">
       <img alt="Build Status" src="https://godoc.org/github.com/zerjioang/ssscomp?status.svg">
    </a>
</p>

Package **ssscomp** is a High Performance, Pure Go Secret Sharing & Secure Computation Library including support for several HE (Homomorphic Encryption) algorithms. It provides a Development Kit for Python too!

## Install

```bash
Copyright (C) 2019 Sergio Anguita
See Contributors.md for a complete list of contributors.  
Licensed under the MIT License.  
```

```bash
go get github.com/zerjioang/ssscomp
```

## Install for Python

```bash
Copyright (C) 2019 Sergio Anguita
See Contributors.md for a complete list of contributors.  
Licensed under the MIT License.  
```


This package is designed to be compatible with current Machine Learning frameworks such as: `pandas`, `keras` or `tensorflow`
Use `pip install ssscomp` to download the package to you environment. Development Python SDK is provided with examples at [sdk/python/example](./sdk/python/example) 

## Features

| Supported Features           	| Properties              	| Limitations       	| Reference papers 	|
|-------------------------------|---------------------------|-----------------------|------------------:|
| Additive Sharing Schema      	| Secret Sharing for SMPC 	|                   	|                  	|
| Shamir Secret Sharing Schema 	| Secret Sharing for SMPC 	|                   	|                  	|
| Unpadded RSA                 	| Homomorphic Encryption  	| `mul` only        	|                  	|
| ElGamal                   	| Homomorphic Encryption  	| `add` only        	|                  	|
| BGN Cryptosystem             	| Homomorphic Encryption  	| `mul`, `add` only  	|                  	|

### Supported Secret Sharing

#### Description

The most common examples of finite fields are given by the integers mod p when p is a prime number. 

#### Packed sharing
If many secrets are to be secret shared, it may be beneficial to use the packed scheme where several secrets are packed into each share. While still very computational efficient, one downside is that the parameters are somewhat restricted.

Specifically, the parameters are split in *scheme parameters* and *implementation parameters*:
- the former, like in Shamir sharing, determines the abstract properties of the scheme, yet now also with a `secret_count` specifying how many secrets are to be packed into each share; the reconstruction limit is implicitly defined as `secret_count + threshold + 1`
- the latter is related to the implementation (currently based on the Fast Fourier Transform) and requires not only a `prime` specifying the field, but also two principal roots of unity within that field, which must be respectively a power of 2 and a power of 3

Due to this increased complexity, providing helper functions for finding suitable parameters are in progress. For now, a few fixed fields are included in the `packed` module as illustrated in the example below:

- `PSS_4_8_3`, `PSS_4_26_3`, `PSS_155_728_100`, `PSS_155_19682_100`

with format `PSS_T_N_D` for sharing `D` secrets into `N` shares with a threshold of `T`.

#### Homomorphic properties

Both the Shamir and the packed scheme enjoy certain homomorphic properties: shared secrets can be transformed by manipulating the shares. Both addition and multiplications work, yet notice that the reconstruction limit in the case of multiplication goes up by a factor of two for each application.

#### Other SSS Implementations

* https://github.com/itslab-kyushu/sss
* https://github.com/amousa11/sss
* https://github.com/sam701/secret-sharing

### Supported Homomorphic Encryption Algorithms

This library has support for current homomorphic algorithms:

* Unpadded RSA
* ElGamal
* Paillier
* DGHV