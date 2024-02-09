package ram

// Load Immediate F and Display 15
var Program = [16]uint8{
	0b00111001, // [00] 0000  1111	LDI
	0b11100000, // [04] 1110  0000	OUT
	0b11110000, // [05] 1111  0000	HLT
	0b00000000, // [06] 0000  0000
}

// 0x10 + 0x14 + 0x18 - 0x20 = 0x1B   0b 0001 1100
//   16 +   20 +   24 -   32 =   28
// var Program = [16]uint8{
// 	0b00001001, // [00] 0000  1001	LDA 9H  [0001 0000]
// 	0b00011010, // [01] 0001  1010	ADD AH  [0010 0100]
// 	0b00011011, // [02] 0001  1011	ADD BH  [0011 1100]
// 	0b00101100, // [03] 0010  1100	SUB CH  [0001 1100]
// 	0b11100000, // [04] 1110  0000	OUT
// 	0b11110000, // [05] 1111  0000	HLT
// 	0b00000000, // [06] 0000  0000
// 	0b00000000, // [07] 0000  0000
// 	0b00000000, // [08] 0000  0000
// 	0b00010000, // [09] 0001  0000	0x10 (016) <- Address (9H)
// 	0b00010100, // [10] 0001  0100	0x14 (020) <- Address (AH)
// 	0b00011000, // [11] 0001  1000	0x18 (024) <- Address (BH)
// 	0b00100000, // [12] 0010  0000	0x20 (032) <- Address (CH)
// 	0b00000000, // [13] 0000  0000
// 	0b00000000, // [14] 0000  0000
// 	0b00000000, // [15] 0000  0000
// }

// Display 016
// var Program = [16]uint8{
// 	0b00001001, // [00] 0000  1001	LDA 9H  [0001 0000]
// 	0b11100000, // [04] 1110  0000	OUT
// 	0b11110000, // [05] 1111  0000	HLT
// 	0b00000000, // [06] 0000  0000
// 	0b00000000, // [07] 0000  0000
// 	0b00000000, // [08] 0000  0000
// 	0b00000000, // [06] 0000  0000
// 	0b00000000, // [07] 0000  0000
// 	0b00000000, // [08] 0000  0000
// 	0b00010000, // [09] 0001  0000	0x10 (016) <- Address (9H)
// 	0b00010100, // [10] 0001  0100	0x14 (020) <- Address (AH)
// 	0b00011000, // [11] 0001  1000	0x18 (024) <- Address (BH)
// 	0b00100000, // [12] 0010  0000	0x20 (032) <- Address (CH)
// 	0b00000000, // [13] 0000  0000
// 	0b00000000, // [14] 0000  0000
// 	0b00000000, // [15] 0000  0000
// }