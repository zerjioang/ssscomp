# CL-15

## Abstract

When data maintained in a decentralized fashion needs to be synchronized or ex-changed  between  different  databases,  related  data  sets  usually  get  associated  with  aunique identifier.  While this approach facilitates cross-domain data exchange,  it alsocomes  with  inherent  drawbacks  in  terms  of  controllability.   As  data  records  can  eas-ily  be  linked,  no  central  authority  can  limit  or  control  the  information  flow.   Worse,when  records  contain  sensitive  personal  data,  as  is  for  instance  the  case  in  nationalsocial  security  systems,  such  linkability  poses  a  massive  security  and  privacy  threat.An  alternative  approach  is  to  use  domain-specific  pseudonyms,  where  only  a  centralauthority knows the cross-domain relation between the pseudonyms.  However, currentsolutions require the central authority to be a fully trusted party, as otherwise it canprovide false conversions and exploit the data it learns from the requests.  We proposean (un)linkable pseudonym system that overcomes those limitations, and enables con-trolled yet privacy-friendly exchange of distributed data.  We prove our protocol securein the UC framework and provide an efficient instantiation based on discrete-logarithmrelated assumptions.
## References

* Privacy for Distributed Databases via (Un)linkable Pseudonyms (https://eprint.iacr.org/2017/022.pdf)
* [CL15] Camenisch, Lehmann. (Un)linkable Pseudonyms for Governmental Databases. CCS’15. (https://www.zurich.ibm.com/pdf/csc/pseudonyms_paper.pdf)
* [CL17] Camenisch, Lehmann. Privacy-Preserving User-Auditable Pseudonym Systems. IEEE EuroSP’17 (https://www.researchgate.net/profile/Jan_Camenisch/publication/318125386_Privacy-Preserving_User-Auditable_Pseudonym_Systems/links/5a969ab1aca272140569f0eb/Privacy-Preserving-User-Auditable-Pseudonym-Systems.pdf)
