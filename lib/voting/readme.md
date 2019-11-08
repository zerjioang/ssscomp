# Fair and Honest voting

## Problem description
So how can be created a trusted voting system? With this we want to make sure that Alice, Bob, Carol, Dave and Eve can vote in an election, but keep their votes secret? But none of them trust Trent to count the votes, so can we find a way for them all to vote, and broadcast their votes to everyone else, and then for each of them to know the result of the vote, without revealing anyones vote? 

## Method

So how can be created a trusted voting system? With this we want to make sure that Alice, Bob, Carol, Dave and Eve can vote in an election, but keep their votes secret. But none of them trust Trent to count the votes, so can we find a way for them all to vote, and broadcast their votes to everyone else, and then for each of them to know the result of the vote, without revealing anyones vote? 

One of the advantages of using Blockchain is the application of smart contracts which can be used for self-tallying voting. Within [1] the authors define a smart contract method within Ethereum, and where there is no need for a trust infrastructure and where the privacy of the voter is preserved. The only way to breach the vote is where all the voters collude against the system. Zero-knowledge proof using is implemented using the Schnorr proof [2] and is made non-interactive using the Fiat-Shamir heuristic [3]. In their solution (as illustrated in Figure 1), voters register their voting key with (for voter i)

## References

[1] P. McCorry, S. F. Shahandashti, and F. Hao, “A smart contract for board- room voting with maximum voter privacy,” in International Conference on Financial Cryptography and Data Security. Springer, 2017, pp.357–375.

[2] C.-P. Schnorr, “Efficient signature generation by smart cards,” Journal of cryptology , vol. 4, no. 3, pp. 161–174, 1991.

[3] A. Fiat and A. Shamir, “How to prove yourself: Practical solutions to identification and signature problems,” in Conference on the Theory and Application of Cryptographic Techniques. Springer, 1986, pp. 186–194.