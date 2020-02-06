package rlwe

//type RINGELT uint32
type RINGELT uint32
type FFTLONG uint64

var W = [1024]RINGELT{1, 40, 1600, 23039, 20418, 38461, 22883, 14178, 34627, 33367, 23928, 15017, 27226, 24054, 20057, 24021, 18737, 12182, 36709, 34725, 37287, 16884, 19984, 21101, 24820, 9736, 20791, 12420, 5268, 5915, 31795, 2009, 39399, 19442, 40382, 17801, 15703, 13705, 15707, 13865, 22107, 24099, 21857, 14099, 31467, 29850, 6131, 40435, 19921, 18581, 5942, 32875, 4248, 6076, 38235, 13843, 21227, 29860, 6531, 15474, 4545, 17956, 21903, 15939, 23145, 24658, 3256, 7357, 7553, 15393, 1305, 11239, 39950, 521, 20840, 14380, 1746, 28879, 8252, 2392, 13758, 17827, 16743, 14344, 306, 12240, 39029, 4642, 21836, 13259, 38828, 37563, 27924, 11013, 30910, 7570, 16073, 28505, 34253, 18407, 39943, 241, 9640, 16951, 22664, 5418, 11915, 26029, 17135, 30024, 13091, 32108, 14529, 7706, 21513, 339, 13560, 9907, 27631, 40254, 12681, 15708, 13905, 23707, 6177, 1314, 11599, 13389, 3067, 40758, 32841, 2888, 33598, 33168, 15968, 24305, 30097, 16011, 26025, 16975, 23624, 2857, 32358, 24529, 39057, 5762, 25675, 2975, 37078, 8524, 13272, 39348, 17402, 40704, 30681, 39371, 18322, 36543, 28085, 17453, 1783, 30359, 26491, 35615, 31926, 7249, 3233, 6437, 11714, 17989, 23223, 27778, 5173, 2115, 2678, 25198, 24856, 11176, 37430, 22604, 3018, 38798, 36363, 20885, 16180, 32785, 648, 25920, 12775, 19468, 461, 18440, 302, 12080, 32629, 35369, 22086, 23259, 29218, 21812, 12299, 428, 17120, 29424, 30052, 14211, 35947, 4245, 5956, 33435, 26648, 934, 37360, 19804, 13901, 23547, 40738, 32041, 11849, 23389, 34418, 25007, 17216, 33264, 19808, 14061, 29947, 10011, 31791, 1849, 32999, 9208, 40632, 27801, 6093, 38915, 82, 3280, 8317, 4992, 35836, 40766, 33161, 15688, 13105, 32668, 36929, 2564, 20638, 6300, 6234, 3594, 20877, 15860, 19985, 21141, 26420, 32775, 248, 9920, 28151, 20093, 25461, 35376, 22366, 34459, 26647, 894, 35760, 37726, 34444, 26047, 17855, 17863, 18183, 30983, 10490, 9990, 30951, 9210, 40712, 31001, 11210, 38790, 36043, 8085, 36673, 33285, 20648, 6700, 22234, 29179, 20252, 31821, 3049, 40038, 4041, 38757, 34723, 37207, 13684, 14867, 21226, 29820, 4931, 33396, 25088, 20456, 39981, 1761, 29479, 32252, 20289, 33301, 21288, 32300, 22209, 28179, 21213, 29300, 25092, 20616, 5420, 11995, 29229, 22252, 29899, 8091, 36913, 1924, 35999, 6325, 7234, 2633, 23398, 34778, 39407, 19762, 12221, 38269, 15203, 34666, 34927, 4406, 12396, 4308, 8476, 11352, 3509, 17477, 2743, 27798, 5973, 34115, 12887, 23948, 15817, 18265, 34263, 18807, 14982, 25826, 9015, 32912, 5728, 24315, 30497, 32011, 10649, 16350, 39585, 26882, 10294, 2150, 4078, 40237, 12001, 29469, 31852, 4289, 7716, 21913, 16339, 39145, 9282, 2631, 23318, 31578, 34290, 19887, 17221, 33464, 27808, 6373, 9154, 38472, 23323, 31778, 1329, 12199, 37389, 20964, 19340, 36302, 18445, 502, 20080, 24941, 14576, 9586, 14791, 18186, 31103, 15290, 38146, 10283, 1710, 27439, 32574, 33169, 16008, 25905, 12175, 36429, 23525, 39858, 37802, 37484, 24764, 7496, 13113, 32988, 8768, 23032, 20138, 27261, 25454, 35096, 11166, 37030, 6604, 18394, 39423, 20402, 37821, 38244, 14203, 35627, 32406, 26449, 33935, 5687, 22675, 5858, 29515, 33692, 36928, 2524, 19038, 24222, 26777, 6094, 38955, 1682, 26319, 28735, 2492, 17758, 13983, 26827, 8094, 37033, 6724, 23194, 26618, 40695, 30321, 24971, 15776, 16625, 9624, 16311, 38025, 5443, 12915, 25068, 19656, 7981, 32513, 30729, 330, 13200, 36468, 25085, 20336, 35181, 14566, 9186, 39752, 33562, 31728, 40290, 14121, 32347, 24089, 21457, 39060, 5882, 30475, 31131, 16410, 1024, 40960, 40921, 39361, 17922, 20543, 2500, 18078, 26783, 6334, 7594, 17033, 25944, 13735, 16907, 20904, 16940, 22224, 28779, 4252, 6236, 3674, 24077, 20977, 19860, 16141, 31225, 20170, 28541, 35693, 35046, 9166, 38952, 1562, 21519, 579, 23160, 25258, 27256, 25254, 27096, 18854, 16862, 19104, 26862, 9494, 11111, 34830, 526, 21040, 22380, 35019, 8086, 36713, 34885, 2726, 27118, 19734, 11101, 34430, 25487, 36416, 23005, 19058, 25022, 17816, 16303, 37705, 33604, 33408, 25568, 39656, 29722, 1011, 40440, 20121, 26581, 39215, 12082, 32709, 38569, 27203, 23134, 24218, 26617, 40655, 28721, 1932, 36319, 19125, 27702, 2133, 3398, 13037, 29948, 10051, 33391, 24888, 12456, 6708, 22554, 1018, 40720, 31321, 24010, 18297, 35543, 29046, 14932, 23826, 10937, 27870, 8853, 26432, 33255, 19448, 40622, 27401, 31054, 13330, 707, 28280, 25253, 27056, 17254, 34784, 39647, 29362, 27572, 37894, 203, 8120, 38073, 7363, 7793, 24993, 16656, 10864, 24950, 14936, 23986, 17337, 38104, 8603, 16432, 1904, 35199, 15286, 37986, 3883, 32437, 27689, 1613, 23559, 257, 10280, 1590, 22639, 4418, 12876, 23508, 39178, 10602, 14470, 5346, 9035, 33712, 37728, 34524, 29247, 22972, 17738, 13183, 35788, 38846, 38283, 15763, 16105, 29785, 3531, 18357, 37943, 2163, 4598, 20076, 24781, 8176, 40313, 15041, 28186, 21493, 40500, 22521, 40659, 28881, 8332, 5592, 18875, 17702, 11743, 19149, 28662, 40533, 23841, 11537, 10909, 26750, 5014, 36716, 35005, 7526, 14313, 40027, 3601, 21157, 27060, 17414, 223, 8920, 29112, 17572, 6543, 15954, 23745, 7697, 21153, 26900, 11014, 30950, 9170, 39112, 7962, 31753, 329, 13160, 34868, 2046, 40879, 37681, 32644, 35969, 5125, 195, 7800, 25273, 27856, 8293, 4032, 38397, 20323, 34661, 34727, 37367, 20084, 25101, 20976, 19820, 14541, 8186, 40713, 31041, 12810, 20868, 15500, 5585, 18595, 6502, 14314, 40067, 5201, 3235, 6517, 14914, 23106, 23098, 22778, 9978, 30471, 30971, 10010, 31751, 249, 9960, 29751, 2171, 4918, 32876, 4288, 7676, 20313, 34261, 18727, 11782, 20709, 9140, 37912, 923, 36920, 2204, 6238, 3754, 27277, 26094, 19735, 11141, 36030, 7565, 15873, 20505, 980, 39200, 11482, 8709, 20672, 7660, 19673, 8661, 18752, 12782, 19748, 11661, 15869, 20345, 35541, 28966, 11732, 18709, 11062, 32870, 4048, 39037, 4962, 34636, 33727, 38328, 17563, 6183, 1554, 21199, 28740, 2692, 25758, 6295, 6034, 36555, 28565, 36653, 32485, 29609, 37452, 23484, 38218, 13163, 34988, 6846, 28074, 17013, 25144, 22696, 6698, 22154, 25979, 15135, 31946, 8049, 35233, 16646, 10464, 8950, 30312, 24611, 1376, 14079, 30667, 38811, 36883, 724, 28960, 11492, 9109, 36672, 33245, 19048, 24622, 1816, 31679, 38330, 17643, 9383, 6671, 21074, 23740, 7497, 13153, 34588, 31807, 2489, 17638, 9183, 39632, 28762, 3572, 19997, 21621, 4659, 22516, 40459, 20881, 16020, 26385, 31375, 26170, 22775, 9858, 25671, 2815, 30678, 39251, 13522, 8387, 7792, 24953, 15056, 28786, 4532, 17436, 1103, 3159, 3477, 16197, 33465, 27848, 7973, 32193, 17929, 20823, 13700, 15507, 5865, 29795, 3931, 34357, 22567, 1538, 20559, 3140, 2717, 26758, 5334, 8555, 14512, 7026, 35274, 18286, 35103, 11446, 7269, 4033, 38437, 21923, 16739, 14184, 34867, 2006, 39279, 14642, 12226, 38469, 23203, 26978, 14134, 32867, 3928, 34237, 17767, 14343, 266, 10640, 15990, 25185, 24336, 31337, 24650, 2936, 35518, 28046, 15893, 21305, 32980, 8448, 10232, 40631, 27761, 4493, 15876, 20625, 5780, 26395, 31775, 1209, 7399, 9233, 671, 26840, 8614, 16872, 19504, 1901, 35079, 10486, 9830, 24551, 39937}
var WRev = [1024]RINGELT{1, 39937, 24551, 9830, 10486, 35079, 1901, 19504, 16872, 8614, 26840, 671, 9233, 7399, 1209, 31775, 26395, 5780, 20625, 15876, 4493, 27761, 40631, 10232, 8448, 32980, 21305, 15893, 28046, 35518, 2936, 24650, 31337, 24336, 25185, 15990, 10640, 266, 14343, 17767, 34237, 3928, 32867, 14134, 26978, 23203, 38469, 12226, 14642, 39279, 2006, 34867, 14184, 16739, 21923, 38437, 4033, 7269, 11446, 35103, 18286, 35274, 7026, 14512, 8555, 5334, 26758, 2717, 3140, 20559, 1538, 22567, 34357, 3931, 29795, 5865, 15507, 13700, 20823, 17929, 32193, 7973, 27848, 33465, 16197, 3477, 3159, 1103, 17436, 4532, 28786, 15056, 24953, 7792, 8387, 13522, 39251, 30678, 2815, 25671, 9858, 22775, 26170, 31375, 26385, 16020, 20881, 40459, 22516, 4659, 21621, 19997, 3572, 28762, 39632, 9183, 17638, 2489, 31807, 34588, 13153, 7497, 23740, 21074, 6671, 9383, 17643, 38330, 31679, 1816, 24622, 19048, 33245, 36672, 9109, 11492, 28960, 724, 36883, 38811, 30667, 14079, 1376, 24611, 30312, 8950, 10464, 16646, 35233, 8049, 31946, 15135, 25979, 22154, 6698, 22696, 25144, 17013, 28074, 6846, 34988, 13163, 38218, 23484, 37452, 29609, 32485, 36653, 28565, 36555, 6034, 6295, 25758, 2692, 28740, 21199, 1554, 6183, 17563, 38328, 33727, 34636, 4962, 39037, 4048, 32870, 11062, 18709, 11732, 28966, 35541, 20345, 15869, 11661, 19748, 12782, 18752, 8661, 19673, 7660, 20672, 8709, 11482, 39200, 980, 20505, 15873, 7565, 36030, 11141, 19735, 26094, 27277, 3754, 6238, 2204, 36920, 923, 37912, 9140, 20709, 11782, 18727, 34261, 20313, 7676, 4288, 32876, 4918, 2171, 29751, 9960, 249, 31751, 10010, 30971, 30471, 9978, 22778, 23098, 23106, 14914, 6517, 3235, 5201, 40067, 14314, 6502, 18595, 5585, 15500, 20868, 12810, 31041, 40713, 8186, 14541, 19820, 20976, 25101, 20084, 37367, 34727, 34661, 20323, 38397, 4032, 8293, 27856, 25273, 7800, 195, 5125, 35969, 32644, 37681, 40879, 2046, 34868, 13160, 329, 31753, 7962, 39112, 9170, 30950, 11014, 26900, 21153, 7697, 23745, 15954, 6543, 17572, 29112, 8920, 223, 17414, 27060, 21157, 3601, 40027, 14313, 7526, 35005, 36716, 5014, 26750, 10909, 11537, 23841, 40533, 28662, 19149, 11743, 17702, 18875, 5592, 8332, 28881, 40659, 22521, 40500, 21493, 28186, 15041, 40313, 8176, 24781, 20076, 4598, 2163, 37943, 18357, 3531, 29785, 16105, 15763, 38283, 38846, 35788, 13183, 17738, 22972, 29247, 34524, 37728, 33712, 9035, 5346, 14470, 10602, 39178, 23508, 12876, 4418, 22639, 1590, 10280, 257, 23559, 1613, 27689, 32437, 3883, 37986, 15286, 35199, 1904, 16432, 8603, 38104, 17337, 23986, 14936, 24950, 10864, 16656, 24993, 7793, 7363, 38073, 8120, 203, 37894, 27572, 29362, 39647, 34784, 17254, 27056, 25253, 28280, 707, 13330, 31054, 27401, 40622, 19448, 33255, 26432, 8853, 27870, 10937, 23826, 14932, 29046, 35543, 18297, 24010, 31321, 40720, 1018, 22554, 6708, 12456, 24888, 33391, 10051, 29948, 13037, 3398, 2133, 27702, 19125, 36319, 1932, 28721, 40655, 26617, 24218, 23134, 27203, 38569, 32709, 12082, 39215, 26581, 20121, 40440, 1011, 29722, 39656, 25568, 33408, 33604, 37705, 16303, 17816, 25022, 19058, 23005, 36416, 25487, 34430, 11101, 19734, 27118, 2726, 34885, 36713, 8086, 35019, 22380, 21040, 526, 34830, 11111, 9494, 26862, 19104, 16862, 18854, 27096, 25254, 27256, 25258, 23160, 579, 21519, 1562, 38952, 9166, 35046, 35693, 28541, 20170, 31225, 16141, 19860, 20977, 24077, 3674, 6236, 4252, 28779, 22224, 16940, 20904, 16907, 13735, 25944, 17033, 7594, 6334, 26783, 18078, 2500, 20543, 17922, 39361, 40921, 40960, 1024, 16410, 31131, 30475, 5882, 39060, 21457, 24089, 32347, 14121, 40290, 31728, 33562, 39752, 9186, 14566, 35181, 20336, 25085, 36468, 13200, 330, 30729, 32513, 7981, 19656, 25068, 12915, 5443, 38025, 16311, 9624, 16625, 15776, 24971, 30321, 40695, 26618, 23194, 6724, 37033, 8094, 26827, 13983, 17758, 2492, 28735, 26319, 1682, 38955, 6094, 26777, 24222, 19038, 2524, 36928, 33692, 29515, 5858, 22675, 5687, 33935, 26449, 32406, 35627, 14203, 38244, 37821, 20402, 39423, 18394, 6604, 37030, 11166, 35096, 25454, 27261, 20138, 23032, 8768, 32988, 13113, 7496, 24764, 37484, 37802, 39858, 23525, 36429, 12175, 25905, 16008, 33169, 32574, 27439, 1710, 10283, 38146, 15290, 31103, 18186, 14791, 9586, 14576, 24941, 20080, 502, 18445, 36302, 19340, 20964, 37389, 12199, 1329, 31778, 23323, 38472, 9154, 6373, 27808, 33464, 17221, 19887, 34290, 31578, 23318, 2631, 9282, 39145, 16339, 21913, 7716, 4289, 31852, 29469, 12001, 40237, 4078, 2150, 10294, 26882, 39585, 16350, 10649, 32011, 30497, 24315, 5728, 32912, 9015, 25826, 14982, 18807, 34263, 18265, 15817, 23948, 12887, 34115, 5973, 27798, 2743, 17477, 3509, 11352, 8476, 4308, 12396, 4406, 34927, 34666, 15203, 38269, 12221, 19762, 39407, 34778, 23398, 2633, 7234, 6325, 35999, 1924, 36913, 8091, 29899, 22252, 29229, 11995, 5420, 20616, 25092, 29300, 21213, 28179, 22209, 32300, 21288, 33301, 20289, 32252, 29479, 1761, 39981, 20456, 25088, 33396, 4931, 29820, 21226, 14867, 13684, 37207, 34723, 38757, 4041, 40038, 3049, 31821, 20252, 29179, 22234, 6700, 20648, 33285, 36673, 8085, 36043, 38790, 11210, 31001, 40712, 9210, 30951, 9990, 10490, 30983, 18183, 17863, 17855, 26047, 34444, 37726, 35760, 894, 26647, 34459, 22366, 35376, 25461, 20093, 28151, 9920, 248, 32775, 26420, 21141, 19985, 15860, 20877, 3594, 6234, 6300, 20638, 2564, 36929, 32668, 13105, 15688, 33161, 40766, 35836, 4992, 8317, 3280, 82, 38915, 6093, 27801, 40632, 9208, 32999, 1849, 31791, 10011, 29947, 14061, 19808, 33264, 17216, 25007, 34418, 23389, 11849, 32041, 40738, 23547, 13901, 19804, 37360, 934, 26648, 33435, 5956, 4245, 35947, 14211, 30052, 29424, 17120, 428, 12299, 21812, 29218, 23259, 22086, 35369, 32629, 12080, 302, 18440, 461, 19468, 12775, 25920, 648, 32785, 16180, 20885, 36363, 38798, 3018, 22604, 37430, 11176, 24856, 25198, 2678, 2115, 5173, 27778, 23223, 17989, 11714, 6437, 3233, 7249, 31926, 35615, 26491, 30359, 1783, 17453, 28085, 36543, 18322, 39371, 30681, 40704, 17402, 39348, 13272, 8524, 37078, 2975, 25675, 5762, 39057, 24529, 32358, 2857, 23624, 16975, 26025, 16011, 30097, 24305, 15968, 33168, 33598, 2888, 32841, 40758, 3067, 13389, 11599, 1314, 6177, 23707, 13905, 15708, 12681, 40254, 27631, 9907, 13560, 339, 21513, 7706, 14529, 32108, 13091, 30024, 17135, 26029, 11915, 5418, 22664, 16951, 9640, 241, 39943, 18407, 34253, 28505, 16073, 7570, 30910, 11013, 27924, 37563, 38828, 13259, 21836, 4642, 39029, 12240, 306, 14344, 16743, 17827, 13758, 2392, 8252, 28879, 1746, 14380, 20840, 521, 39950, 11239, 1305, 15393, 7553, 7357, 3256, 24658, 23145, 15939, 21903, 17956, 4545, 15474, 6531, 29860, 21227, 13843, 38235, 6076, 4248, 32875, 5942, 18581, 19921, 40435, 6131, 29850, 31467, 14099, 21857, 24099, 22107, 13865, 15707, 13705, 15703, 17801, 40382, 19442, 39399, 2009, 31795, 5915, 5268, 12420, 20791, 9736, 24820, 21101, 19984, 16884, 37287, 34725, 36709, 12182, 18737, 24021, 20057, 24054, 27226, 15017, 23928, 33367, 34627, 14178, 22883, 38461, 20418, 23039, 1600, 40}
var WSqrt = [512]RINGELT{16186, 33025, 10248, 310, 12400, 4468, 14876, 21586, 3259, 7477, 12353, 2588, 21598, 3739, 26677, 2094, 1838, 32559, 32569, 32969, 8008, 33593, 32968, 7968, 31993, 9929, 28511, 34493, 28007, 14333, 40827, 35601, 31366, 25810, 8375, 7312, 5753, 25315, 29536, 34532, 29567, 35772, 38206, 12683, 15788, 17105, 28824, 6052, 37275, 16404, 784, 31360, 25570, 39736, 32922, 6128, 40315, 15121, 31386, 26610, 40375, 17521, 4503, 16276, 36625, 31365, 25770, 6775, 25234, 26296, 27815, 6653, 20354, 35901, 2405, 14278, 38627, 29523, 34012, 8767, 22992, 18538, 4222, 5036, 37596, 29244, 22852, 12938, 25988, 15495, 5385, 10595, 14190, 35107, 11606, 13669, 14267, 38187, 11923, 26349, 29935, 9531, 12591, 12108, 33749, 39208, 11802, 21509, 179, 7160, 40634, 27881, 9293, 3071, 40918, 39241, 13122, 33348, 23168, 25578, 40056, 4761, 26596, 39815, 36082, 9645, 17151, 30664, 38691, 32083, 13529, 8667, 18992, 22382, 35099, 11286, 869, 34760, 38687, 31923, 7129, 39394, 19242, 32382, 25489, 36496, 26205, 24175, 24897, 12816, 21108, 25100, 20936, 18220, 32463, 28729, 2252, 8158, 39593, 27202, 23094, 22618, 3578, 20237, 31221, 20010, 22141, 25459, 35296, 19166, 29342, 26772, 5894, 30955, 9370, 6151, 274, 10960, 28790, 4692, 23836, 11337, 2909, 34438, 25807, 8255, 2512, 18558, 5022, 37036, 6844, 27994, 13813, 20027, 22821, 11698, 17349, 38584, 27803, 6173, 1154, 5199, 3155, 3317, 9797, 23231, 28098, 17973, 22583, 2178, 5198, 3115, 1717, 27719, 2813, 30598, 36051, 8405, 8512, 12792, 20148, 27661, 493, 19720, 10541, 12030, 30629, 37291, 17044, 26384, 31335, 24570, 40697, 30401, 28171, 20893, 16500, 4624, 21116, 25420, 33736, 38688, 31963, 8729, 21472, 39660, 29882, 7411, 9713, 19871, 16581, 7864, 27833, 7373, 8193, 32, 1280, 10239, 40911, 38961, 1922, 35919, 3125, 2117, 2758, 28398, 29973, 11051, 32430, 27409, 31374, 26130, 21175, 27780, 5253, 5315, 7795, 25073, 19856, 15981, 24825, 9936, 28791, 4732, 25436, 34376, 23327, 31938, 7729, 22433, 37139, 10964, 28950, 11092, 34070, 11087, 33870, 3087, 597, 23880, 13097, 32348, 24129, 23057, 21138, 26300, 27975, 13053, 30588, 35651, 33366, 23888, 13417, 4187, 3636, 22557, 1138, 4559, 18516, 3342, 10797, 22270, 30619, 36891, 1044, 799, 31960, 8609, 16672, 11504, 9589, 14911, 22986, 18298, 35583, 30646, 37971, 3283, 8437, 9792, 23031, 20098, 25661, 2415, 14678, 13666, 14147, 33387, 24728, 6056, 37435, 22804, 11018, 31110, 15570, 8385, 7712, 21753, 9939, 28911, 9532, 12631, 13708, 15827, 18665, 9302, 3431, 14357, 826, 33040, 10848, 24310, 30297, 24011, 18337, 37143, 11124, 35350, 21326, 33820, 1087, 2519, 18838, 16222, 34465, 26887, 10494, 10150, 37351, 19444, 40462, 21001, 20820, 13580, 10707, 18670, 9502, 11431, 6669, 20994, 20540, 2380, 13278, 39588, 27002, 15094, 30306, 24371, 32737, 39689, 31042, 12850, 22468, 38539, 26003, 16095, 29385, 28492, 33733, 38568, 27163, 21534, 1179, 6199, 2194, 5838, 28715, 1692, 26719, 3774, 28077, 17133, 29944, 9891, 26991, 14654, 12706, 16708, 12944, 26228, 25095, 20736, 10220, 40151, 8561, 14752, 16626, 9664, 17911, 20103, 25861, 10415, 6990, 33834, 1647, 24919, 13696, 15347, 40426, 19561, 4181, 3396, 12957, 26748, 4934, 33516, 29888, 7651, 19313, 35222, 16206, 33825, 1287, 10519, 11150, 36390, 21965, 18419, 40423, 19441, 40342, 16201, 33625, 34248, 18207, 31943, 7929, 30433, 29451, 31132, 16450, 2624, 23038, 20378, 36861, 40805, 34721, 37127, 10484, 9750, 21351, 34820, 126, 5040, 37756, 35644, 33086, 12688, 15988, 25105, 21136, 26220}
var WSqrtRev = [512]RINGELT{14741, 19825, 15856, 24973, 28273, 7875, 5317, 3205, 35921, 40835, 6141, 19610, 31211, 30477, 3834, 6240, 156, 4100, 20583, 17923, 38337, 24511, 9829, 11510, 10528, 33032, 9018, 22754, 6713, 7336, 24760, 619, 21520, 538, 22542, 18996, 4571, 29811, 30442, 39674, 7136, 24755, 5739, 21648, 33310, 11073, 7445, 36027, 14213, 28004, 37565, 36780, 21400, 535, 25614, 27265, 16042, 39314, 7127, 33971, 30546, 15100, 20858, 23050, 31297, 24335, 26209, 32400, 810, 30741, 20225, 15866, 14733, 28017, 24253, 28255, 26307, 13970, 31070, 11017, 23828, 12884, 37187, 14242, 39269, 12246, 35123, 38767, 34762, 39782, 19427, 13798, 2393, 7228, 12469, 11576, 24866, 14958, 2422, 18493, 28111, 9919, 1272, 8224, 16590, 10655, 25867, 13959, 1373, 27683, 38581, 20421, 19967, 34292, 29530, 31459, 22291, 30254, 27381, 20141, 19960, 499, 21517, 3610, 30811, 30467, 14074, 6496, 24739, 22123, 38442, 39874, 7141, 19635, 5611, 29837, 3818, 22624, 16950, 10664, 16651, 30113, 7921, 40135, 26604, 37530, 31659, 22296, 25134, 27253, 28330, 31429, 12050, 31022, 19208, 33249, 32576, 25391, 9851, 29943, 18157, 3526, 34905, 16233, 7574, 26814, 27295, 26283, 38546, 15300, 20863, 17930, 31169, 32524, 37678, 2990, 10315, 5378, 22663, 17975, 26050, 31372, 29457, 24289, 32352, 9001, 40162, 39917, 4070, 10342, 18691, 30164, 37619, 22445, 36402, 39823, 18404, 37325, 36774, 27544, 17073, 7595, 5310, 10373, 27908, 12986, 14661, 19823, 17904, 16832, 8613, 27864, 17081, 40364, 37874, 7091, 29874, 6891, 29869, 12011, 29997, 3822, 18528, 33232, 9023, 17634, 6585, 15525, 36229, 12170, 31025, 16136, 24980, 21105, 15888, 33166, 35646, 35708, 13181, 19786, 14831, 9587, 13552, 8531, 29910, 10988, 12563, 38203, 38844, 37836, 5042, 39039, 2000, 50, 30722, 39681, 40929, 32768, 33588, 13128, 33097, 24380, 21090, 31248, 33550, 11079, 1301, 19489, 32232, 8998, 2273, 7225, 15541, 19845, 36337, 24461, 20068, 12790, 10560, 264, 16391, 9626, 14577, 23917, 3670, 10332, 28931, 30420, 21241, 40468, 13300, 20813, 28169, 32449, 32556, 4910, 10363, 38148, 13242, 39244, 37846, 35763, 38783, 18378, 22988, 12863, 17730, 31164, 37644, 37806, 35762, 39807, 34788, 13158, 2377, 23612, 29263, 18140, 20934, 27148, 12967, 34117, 3925, 35939, 22403, 38449, 32706, 15154, 6523, 38052, 29624, 17125, 36269, 12171, 30001, 40687, 34810, 31591, 10006, 35067, 14189, 11619, 21795, 5665, 15502, 18820, 20951, 9740, 20724, 37383, 18343, 17867, 13759, 1368, 32803, 38709, 12232, 8498, 22741, 20025, 15861, 19853, 28145, 16064, 16786, 14756, 4465, 15472, 8579, 21719, 1567, 33832, 9038, 2274, 6201, 40092, 29675, 5862, 18579, 21969, 32294, 27432, 8878, 2270, 10297, 23810, 31316, 4879, 1146, 14365, 36200, 905, 15383, 17793, 7613, 27839, 1720, 43, 37890, 31668, 13080, 327, 33801, 40782, 19452, 29159, 1753, 7212, 28853, 28370, 31430, 11026, 14612, 29038, 2774, 26694, 27292, 29355, 5854, 26771, 30366, 35576, 25466, 14973, 28023, 18109, 11717, 3365, 35925, 36739, 22423, 17969, 32194, 6949, 11438, 2334, 26683, 38556, 5060, 20607, 34308, 13146, 14665, 15727, 34186, 15191, 9596, 4336, 24685, 36458, 23440, 586, 14351, 9575, 25840, 646, 34833, 8039, 1225, 15391, 9601, 40177, 24557, 3686, 34909, 12137, 23856, 25173, 28278, 2755, 5189, 11394, 6429, 11425, 15646, 35208, 33649, 32586, 15151, 9595, 5360, 134, 26628, 12954, 6468, 12450, 31032, 8968, 32993, 7993, 7368, 32953, 7992, 8392, 8402, 39123, 38867, 14284, 37222, 19363, 38373, 28608, 33484, 37702, 19375, 26085, 36493, 28561, 40651, 30713, 7936, 24775}

var MISPOWEROFTWO RINGELT = 1

/* Public Parameter a. Each a parameter rejection sampled from non-overlapping
 * segments of the digits of e.
 * Note that this is held in the FFT / CRT basis.*/
var a = [1024]RINGELT{
	0x678B, 0x4782, 0x2A5E, 0x0D10, 0x8D67, 0x0D4F, 0x80F3, 0x3D02,
	0x342A, 0x2179, 0x805F, 0x2A12, 0x5776, 0x5C7C, 0x26D7, 0x3F1A,
	0x5304, 0x6507, 0x90C9, 0x6119, 0x6EAF, 0x2F19, 0x05D3, 0x839B,
	0x7987, 0x856A, 0x834F, 0x802E, 0x1D79, 0x71DC, 0x9F6C, 0x6000,
	0x31C9, 0x02B2, 0x8519, 0x0AF0, 0x2C6E, 0x50FA, 0x0655, 0x919A,
	0x2ADD, 0x6F1F, 0x2D18, 0x2CD8, 0x358C, 0x4D7B, 0x2C71, 0x7863,
	0x5D5A, 0x2D75, 0x3A6C, 0x895D, 0x8197, 0x974E, 0x000C, 0x15BF,
	0x37F2, 0x7AA2, 0x6281, 0x51D2, 0x3A9D, 0x3079, 0x8866, 0x3009,
	0x0BD8, 0x9576, 0x34A5, 0x9976, 0x81BE, 0x6A5F, 0x1A9A, 0x0D92,
	0x4253, 0x9F7C, 0x9AAF, 0x8370, 0x57D9, 0x4CFD, 0x83BC, 0x59CC,
	0x95CF, 0x496D, 0x6A2B, 0x1C97, 0x2816, 0x51EE, 0x3C4B, 0x39DD,
	0x05A5, 0x1977, 0x0F25, 0x3EDC, 0x8B65, 0x02E3, 0x479B, 0x0077,
	0x70CB, 0x71D6, 0x56A3, 0x83C7, 0x6FB6, 0x083E, 0x2A7B, 0x1D6F,
	0x6DD7, 0x3E9A, 0x3E58, 0x9EE3, 0x7015, 0x4C4F, 0x493D, 0x1401,
	0x1D52, 0x9542, 0x8E22, 0x7812, 0x8CD1, 0x16B8, 0x9143, 0x7822,
	0x35D0, 0x3DD7, 0x7263, 0x849A, 0x6EE2, 0x473F, 0x8F19, 0x2D21,
	0x3E12, 0x67A3, 0x5E3D, 0x415A, 0x97A0, 0x5A8C, 0x195F, 0x3B80,
	0x4D56, 0x6AB0, 0x07DA, 0x1ADA, 0x4BF1, 0x19EA, 0x2584, 0x0CB8,
	0x5C06, 0x6392, 0x09FD, 0x4944, 0x83BC, 0x36F5, 0x2CE3, 0x12A3,
	0x1DFF, 0x970A, 0x6884, 0x72E9, 0x3D9A, 0x3A38, 0x2F8A, 0x1394,
	0x00C6, 0x7BFA, 0x067A, 0x8610, 0x1378, 0x3F6D, 0x526D, 0x971B,
	0x7CB8, 0x985F, 0x752C, 0x0FA7, 0x3545, 0x371A, 0x91DE, 0x82D5,
	0x428A, 0x9832, 0x20F4, 0x7A76, 0x8B58, 0x8ABB, 0x6FC8, 0x1FBE,
	0x2C7F, 0x1426, 0x5DA1, 0x4D8B, 0x5AF1, 0x9D85, 0x03C8, 0x4D45,
	0x506C, 0x50BD, 0x862F, 0x0EF0, 0x5F9D, 0x621C, 0x8650, 0x52C9,
	0x6E13, 0x5152, 0x6795, 0x9E35, 0x8D4E, 0x37C9, 0x908B, 0x2FD5,
	0x671E, 0x0CAD, 0x8776, 0x6F6E, 0x849A, 0x9EF7, 0x7873, 0x7630,
	0x82ED, 0x3B52, 0x214D, 0x78C7, 0x63A7, 0x69AC, 0x4ABF, 0x7B4F,
	0x3E83, 0x4430, 0x9ACA, 0x801A, 0x2B18, 0x2B92, 0x9D24, 0x8811,
	0x7CFD, 0x3A18, 0x577E, 0x11CA, 0x99EB, 0x6826, 0x01BB, 0x00D9,
	0x5263, 0x220C, 0x917E, 0x726D, 0x2E7B, 0x98AB, 0x1981, 0x3001,
	0x365E, 0x4689, 0x276B, 0x5FD4, 0x4DA0, 0x4F50, 0x0E91, 0x4184,
	0x5C98, 0x7997, 0x6B24, 0x0BF7, 0x384C, 0x6BB6, 0x33B1, 0x4601,
	0x9146, 0x5678, 0x9B14, 0x96AC, 0x12AF, 0x0F36, 0x9B60, 0x1A36,
	0x727C, 0x57F7, 0x65BA, 0x5F2A, 0x88FD, 0x8913, 0x3516, 0x29EF,
	0x3655, 0x2E9A, 0x5D74, 0x084F, 0x4930, 0x8B20, 0x1BD5, 0x1935,
	0x0CEE, 0x51DF, 0x740E, 0x0A1D, 0x994E, 0x6F95, 0x5616, 0x8493,
	0x44A3, 0x6D0C, 0x3D41, 0x60B0, 0x43B5, 0x9419, 0x25E4, 0x1F44,
	0x6CAF, 0x7FD7, 0x75A8, 0x53BB, 0x7D92, 0x54DF, 0x8BF3, 0x76CC,
	0x866F, 0x8A9C, 0x3F29, 0x3281, 0x6CA0, 0x1229, 0x6B05, 0x12CA,
	0x6A15, 0x8C2B, 0x7C51, 0x8087, 0x5C3E, 0x1908, 0x46B6, 0x41E5,
	0x22C4, 0x8FD6, 0x87D6, 0x1BFE, 0x0CEF, 0x9DDB, 0x84F3, 0x3510,
	0x3CD4, 0x5014, 0x59D0, 0x4792, 0x7199, 0x07D9, 0x3A56, 0x5B89,
	0x563B, 0x1DB3, 0x08B2, 0x2B0B, 0x3C21, 0x936F, 0x4FFA, 0x75D4,
	0x03C1, 0x5322, 0x6BA2, 0x31EB, 0x055F, 0x90FA, 0x1A34, 0x1297,
	0x3FCC, 0x6949, 0x5486, 0x6194, 0x7343, 0x8617, 0x6F3E, 0x0300,
	0x5C5F, 0x78DD, 0x9465, 0x72AD, 0x293D, 0x70C4, 0x23C6, 0x842C,
	0x337F, 0x29AE, 0x424B, 0x715F, 0x0364, 0x33AE, 0x42AE, 0x0C07,
	0x6FF4, 0x046E, 0x027E, 0x9AF9, 0x0A7F, 0x616B, 0x15BE, 0x09E0,
	0x0D85, 0x0708, 0x3B43, 0x7C48, 0x9D10, 0x4A0C, 0x6800, 0x8778,
	0x9636, 0x40AA, 0x4E6F, 0x3D4C, 0x305C, 0x54F4, 0x944E, 0x6BCA,
	0x87A3, 0x421D, 0x5358, 0x8C15, 0x9AB3, 0x3FC0, 0x4561, 0x1575,
	0x9FCB, 0x7303, 0x7B5F, 0x591B, 0x6868, 0x4A00, 0x3670, 0x01B4,
	0x715F, 0x47A7, 0x2F8B, 0x9F8D, 0x4F9B, 0x5EE5, 0x7482, 0x1C8C,
	0x9F39, 0x4ABC, 0x9893, 0x9365, 0x8A63, 0x4665, 0x2326, 0x0569,
	0x9B5B, 0x40FB, 0x8B83, 0x1452, 0x8FD6, 0x9570, 0x1365, 0x46DD,
	0x284C, 0x3640, 0x88A3, 0x30F4, 0x61D3, 0x9963, 0x6390, 0x4FE5,
	0x641F, 0x791D, 0x3D41, 0x73DE, 0x5333, 0x187D, 0x1526, 0x386F,
	0x48CC, 0x3A4B, 0x61DC, 0x87A3, 0x4171, 0x02FB, 0x2B27, 0x9A37,
	0x4A61, 0x54BB, 0x5825, 0x718D, 0x238A, 0x18D1, 0x1592, 0x31CC,
	0x3807, 0x4032, 0x750A, 0x88BB, 0x3489, 0x9E1F, 0x0B84, 0x35A0,
	0x6A40, 0x1748, 0x8A91, 0x9CDA, 0x3399, 0x3D92, 0x05DB, 0x1BB2,
	0x3E96, 0x76B2, 0x3114, 0x431C, 0x37FE, 0x566D, 0x2145, 0x5C79,
	0x7340, 0x22C4, 0x936D, 0x360B, 0x876A, 0x6E55, 0x80DA, 0x9C5F,
	0x6B67, 0x7A23, 0x86C1, 0x2E55, 0x7700, 0x363C, 0x1D9C, 0x8186,
	0x4D6D, 0x40F7, 0x1D65, 0x5807, 0x5C13, 0x733F, 0x4B0F, 0x0289,
	0x4C9F, 0x9891, 0x8339, 0x8F13, 0x19C4, 0x558B, 0x4B3F, 0x256D,
	0x3DD9, 0x0989, 0x61C4, 0x9836, 0x18FA, 0x7E6E, 0x3145, 0x9A4D,
	0x7ACD, 0x22FC, 0x523D, 0x46D2, 0x913B, 0x1868, 0x2FB1, 0x8EE9,
	0x81E2, 0x8FAF, 0x652A, 0x06E7, 0x0B89, 0x1AB4, 0x9072, 0x81D9,
	0x0C9A, 0x5D74, 0x29E6, 0x4B04, 0x6E08, 0x1675, 0x79B8, 0x2E98,
	0x6FFE, 0x6E1B, 0x6CCA, 0x1A7F, 0x5AB9, 0x4A36, 0x5946, 0x8A57,
	0x295E, 0x9A46, 0x5F1B, 0x6F89, 0x06F8, 0x51E2, 0x0E8D, 0x9D3F,
	0x1B3C, 0x443E, 0x676A, 0x1270, 0x9178, 0x8F2C, 0x8E9D, 0x1F0C,
	0x5F02, 0x2E44, 0x7F51, 0x27F7, 0x4A1D, 0x0D21, 0x7343, 0x14E1,
	0x8A81, 0x18F6, 0x2640, 0x206B, 0x1751, 0x7C84, 0x4F44, 0x846F,
	0x47FA, 0x472E, 0x9C6E, 0x1F49, 0x52AE, 0x5D05, 0x5281, 0x7A1B,
	0x48FA, 0x3CD2, 0x75DD, 0x894C, 0x3DC1, 0x009E, 0x5555, 0x4670,
	0x4EDB, 0x348C, 0x39E1, 0x1800, 0x0D44, 0x48D3, 0x6B0C, 0x9026,
	0x5A22, 0x19B1, 0x054D, 0x018E, 0x773A, 0x4D3A, 0x5C52, 0x4CB1,
	0x3877, 0x2B68, 0x1867, 0x5825, 0x6B07, 0x26A4, 0x84A0, 0x0A45,
	0x1C05, 0x94E4, 0x44D1, 0x018F, 0x0BEE, 0x742D, 0x50B1, 0x25BC,
	0x0D3E, 0x3927, 0x7527, 0x21D6, 0x1F15, 0x56F0, 0x5F87, 0x8E43,
	0x0C54, 0x9801, 0x5EED, 0x6892, 0x3457, 0x6F75, 0x63A7, 0x818C,
	0x47AF, 0x1FA0, 0x615F, 0x7A22, 0x8BCF, 0x94DD, 0x9DA2, 0x93CB,
	0x0FF2, 0x67C6, 0x9803, 0x969F, 0x8FB2, 0x01BC, 0x09A4, 0x88D0,
	0x7093, 0x2B2B, 0x5F9F, 0x4DF3, 0x066E, 0x562A, 0x5136, 0x3BD0,
	0x1ED7, 0x6EF6, 0x7F9D, 0x7A2A, 0x21C1, 0x369A, 0x9744, 0x6927,
	0x2761, 0x2C10, 0x3734, 0x9952, 0x6775, 0x00B1, 0x838B, 0x33C7,
	0x1F28, 0x1721, 0x73C8, 0x96C7, 0x0A53, 0x9EC0, 0x3440, 0x1A0B,
	0x6D3B, 0x0F11, 0x652B, 0x2C02, 0x0DD7, 0x36ED, 0x8962, 0x047C,
	0x3A44, 0x27D6, 0x667E, 0x6392, 0x873A, 0x4A54, 0x075B, 0x8B80,
	0x51DF, 0x6A7E, 0x64BE, 0x2B65, 0x3770, 0x9AD8, 0x7E8B, 0x6E20,
	0x1F4B, 0x2387, 0x742F, 0x28AC, 0x3143, 0x15E3, 0x1CEA, 0x9CA7,
	0x008A, 0x69B4, 0x4EDD, 0x3E6A, 0x46F9, 0x9C0D, 0x5C45, 0x9ABB,
	0x6EF5, 0x48FC, 0x0363, 0x1CCD, 0x00CA, 0x2F65, 0x7615, 0x3BC3,
	0x1E30, 0x1480, 0x5F16, 0x7F12, 0x5258, 0x33E5, 0x00AB, 0x0D41,
	0x2E47, 0x4FC7, 0x441B, 0x1190, 0x4D5C, 0x9401, 0x7BA5, 0x53B0,
	0x5058, 0x51F5, 0x1694, 0x55EA, 0x6316, 0x39F1, 0x03C2, 0x5518,
	0x3961, 0x2970, 0x07F9, 0x4C0D, 0x205F, 0x8AC5, 0x2F3A, 0x2AAF,
	0x8895, 0x3D01, 0x74BF, 0x1EC8, 0x83C3, 0x239F, 0x33BA, 0x0F63,
	0x222B, 0x969D, 0x275B, 0x6FEA, 0x39FD, 0x087F, 0x9870, 0x12BB,
	0x2108, 0x2517, 0x2261, 0x653E, 0x4168, 0x95CC, 0x6CF3, 0x7788,
	0x3C76, 0x3652, 0x13D5, 0x3F49, 0x0C42, 0x9267, 0x7F6C, 0x9C25,
	0x2D48, 0x34EB, 0x9AE7, 0x7CCB, 0x7DD6, 0x777C, 0x363F, 0x5971,
	0x0BDA, 0x7AAA, 0x754A, 0x9D22, 0x8C89, 0x3213, 0x42E5, 0x0F9F,
	0x3A47, 0x037F, 0x7F89, 0x661C, 0x9410, 0x0C4A, 0x23E2, 0x8B86,
	0x0B3C, 0x9DC0, 0x3997, 0x02F0, 0x3DBD, 0x964F, 0x17F3, 0x6B53,
	0x148E, 0x57E9, 0x1F10, 0x57A1, 0x2A30, 0x5F05, 0x8223, 0x46EE,
	0x5725, 0x4576, 0x8F5C, 0x0FEF, 0x20C9, 0x30AE, 0x8379, 0x7D62,
	0x0599, 0x2765, 0x5B5B, 0x14D3, 0x5004, 0x67F6, 0x6C8F, 0x34EA,
	0x3B7E, 0x8688, 0x0B5F, 0x723E, 0x7079, 0x5232, 0x376F, 0x9415,
	0x1905, 0x58BE, 0x96C0, 0x03D2, 0x7B1A, 0x111D, 0x7E79, 0x6137,
	0x3C2B, 0x1C3B, 0x9490, 0x0BBB, 0x0BCA, 0x6A44, 0x5DE8, 0x5989,
	0x9160, 0x88BC, 0x9D29, 0x8545, 0x4C00, 0x90CE, 0x17F7, 0x6A1E,
	0x808F, 0x02DC, 0x1407, 0x5CE1, 0x337C, 0x5199, 0x4E07, 0x03E3,
	0x1755, 0x1982, 0x1D89, 0x08DE, 0x664A, 0x0DD0, 0x6D51, 0x0DC4,
	0x393C, 0x100B, 0x40CA, 0x72EB, 0x7E64, 0x46A0, 0x3D0F, 0x8837,
	0x36C6, 0x3A8A, 0x9832, 0x5647, 0x5945, 0x58A1, 0x314A, 0x78EF,
	0x99E8, 0x6DEC, 0x7055, 0x8F4C, 0x9A37, 0x07E3, 0x5567, 0x42D4,
	0x5782, 0x6D57, 0x8033, 0x37CA, 0x2290, 0x9146, 0x5CA7, 0x6A2A,
	0x097A, 0x76DC, 0x4ED1, 0x8131, 0x8B22, 0x01DF, 0x730A, 0x9E1D,
	0x315E, 0x06DB, 0x3538, 0x731B, 0x6E46, 0x9D46, 0x1AE5, 0x9B9A,
	0x7E97, 0x4A3D, 0x0416, 0x13AF, 0x61A4, 0x49EA, 0x1747, 0x96C4,
	0x85A3, 0x245C, 0x375B, 0x652E, 0x9494, 0x539F, 0x3AF3, 0x0712,
	0x370D, 0x38F1, 0x67B5, 0x591E, 0x2E1F, 0x813B, 0x618D, 0x0CEE,
	0x69B0, 0x4D15, 0x2143, 0x425C, 0x9A71, 0x0EB8, 0x0DFE, 0x5F53,
	0x48B5, 0x8FDB, 0x8FE7, 0x88D7, 0x63CB, 0x0B3A, 0x83FC, 0x7CC4,
}

var rlweTable = [52][3]uint64{
	{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0x1FFFFFFFFFFFFFFF},
	{0xE0C81DA0D6A8BD22, 0x161ABD186DA13542, 0x5CEF2C248806C827},
	{0x8D026C4E14BC7408, 0x4344C125B3533F22, 0x9186506BCC065F20},
	{0x10AC7CEC7D7E2A3B, 0x5D62CE65E6217813, 0xBAAB5F82BCDB43B3},
	{0x709C92996E94D801, 0x1411F551608E4D22, 0xD7D9769FAD23BCB1},
	{0x6287D827008404B7, 0x7E1526D618902F20, 0xEA9BE2F4D6DDB5ED},
	{0x34CBDC118C15F40E, 0xE7D2A13787E94674, 0xF58A99474919B8C9},
	{0xD521F7EBBBE8C3A2, 0xE8A773D9A1EA0AAB, 0xFB5117812753B7B8},
	{0xC3D9E58131089A6A, 0x148CB49FF716491B, 0xFE151BD0928596D3},
	{0x2E060C4A842A27F6, 0x07E44D009ADB0049, 0xFF487508BA9F7208},
	{0xFCEDEFCFAA887582, 0x1A5409BF5D4B039E, 0xFFC16686270CFC82},
	{0x4FE22E5DF9FAAC20, 0xFDC99BFE0F991958, 0xFFEC8AC3C159431B},
	{0xA36605F81B14FEDF, 0xA6FCD4C13F4AFCE0, 0xFFFA7DF4B6E92C28},
	{0x9D1FDCFF97BBC957, 0x4B869C6286ED0BB5, 0xFFFE94BB4554B5AC},
	{0x6B3EEBA74AAD104B, 0xEC72329E974D63C7, 0xFFFFAADE1B1CAA95},
	{0x48C8DA4009C10760, 0x337F6316C1FF0A59, 0xFFFFEDDC1C6436DC},
	{0x84480A71312F35E7, 0xD95E7B2CD6933C97, 0xFFFFFC7C9DC2569A},
	{0x23C01DAC1513FA0F, 0x8E0B132AE72F729F, 0xFFFFFF61BC337FED},
	{0x90C89D6570165907, 0x05B9D725AAEA5CAD, 0xFFFFFFE6B3CF05F7},
	{0x692E2A94C500EC7D, 0x99E8F72C370F27A6, 0xFFFFFFFC53EA610E},
	{0x28C2998CEAE37CC8, 0xC6E2F0D7CAFA9AB8, 0xFFFFFFFF841943DE},
	{0xC515CF4CB0130256, 0x4745913CB4F9E4DD, 0xFFFFFFFFF12D07EC},
	{0x39F0ECEA047D6E3A, 0xEE62D42142AC6544, 0xFFFFFFFFFE63E348},
	{0xDF11BB25B50462D6, 0x064A0C6CC136E943, 0xFFFFFFFFFFD762C7},
	{0xCDBA0DD69FD2EA0F, 0xC672F3A74DB0F175, 0xFFFFFFFFFFFC5E37},
	{0xFDB966A75F3604D9, 0x6ABEF8B144723D83, 0xFFFFFFFFFFFFB48F},
	{0x3C4FECBB600740D1, 0x697598CEADD71A15, 0xFFFFFFFFFFFFFA72},
	{0x1574CC916D60E673, 0x12F5A30DD99D7051, 0xFFFFFFFFFFFFFFA1},
	{0xDD3DCD1B9CB7321D, 0x4016ED3E05883572, 0xFFFFFFFFFFFFFFFA},
	{0xB4A4E8CF3DF79A7A, 0xAF22D9AFAD5A73CF, 0xFFFFFFFFFFFFFFFF},
	{0x91056A8196F74466, 0xFBF88681905332BA, 0xFFFFFFFFFFFFFFFF},
	{0x965B9ED9BD366C04, 0xFFD16385AF29A51F, 0xFFFFFFFFFFFFFFFF},
	{0xF05F75D38F2D28A3, 0xFFFE16FF8EA2B60C, 0xFFFFFFFFFFFFFFFF},
	{0x77E35C8980421EE8, 0xFFFFEDD3C9DDC7E8, 0xFFFFFFFFFFFFFFFF},
	{0x92783617956F140A, 0xFFFFFF63392B6E8F, 0xFFFFFFFFFFFFFFFF},
	{0xA536DC994639AD78, 0xFFFFFFFB3592B3D1, 0xFFFFFFFFFFFFFFFF},
	{0x8F3A871874DD9FD5, 0xFFFFFFFFDE04A5BB, 0xFFFFFFFFFFFFFFFF},
	{0x310DE3650170B717, 0xFFFFFFFFFF257152, 0xFFFFFFFFFFFFFFFF},
	{0x1F21A853A422F8CC, 0xFFFFFFFFFFFB057B, 0xFFFFFFFFFFFFFFFF},
	{0x3CA9D5C6DB4EE2BA, 0xFFFFFFFFFFFFE5AD, 0xFFFFFFFFFFFFFFFF},
	{0xCFD9CE958E59869C, 0xFFFFFFFFFFFFFF81, 0xFFFFFFFFFFFFFFFF},
	{0xDB8E1F91D955C452, 0xFFFFFFFFFFFFFFFD, 0xFFFFFFFFFFFFFFFF},
	{0xF78EE3A8E99E08C3, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFE1D7858BABDA25, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFF9E52E32CAB4A, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFEE13217574F, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFD04888041, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFFF8CD8A56, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFFFFF04111, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFFFFFFE0C5, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFFFFFFFFC7, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
}

const m RINGELT = 1024
const M = m
const muwords RINGELT = 16 /* key (mu) is m bits */
const Muwords RINGELT = 16

//var  q RINGELT = 40961
const qmod4 RINGELT = 1

const B RINGELT = 5
const BB RINGELT = 11
const LOG2B RINGELT = 4
const BMASK RINGELT = 0xf

var smallCoeffTable = [11]RINGELT{40956, 40957, 40958, 40959, 40960, 0, 1, 2, 3, 4, 5}

const q_1_4, q_2_4, q_3_4 RINGELT = 10240, 20480, 30721
const r0_l, r0_u, r1_l, r1_u RINGELT = 15360, 35841, 5119, 25601
