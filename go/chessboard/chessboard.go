package chessboard

// File stores whether a square is occupied by a piece
type File []bool

// Chessboard contains a mapping from file names, labeled "A" to "H", to the corresponding File
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for k, v := range cb {
		if k != file {
			continue
		}
		for _, occupied := range v {
			if occupied {
				count += 1
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 0 || rank > 8 {
		return 0
	}
	count := 0
	for _, v := range cb {
		for i, occupied := range v {
			if i != rank-1 {
				continue
			}
			if occupied {
				count += 1
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for range v {
			count += 1
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for _, occupied := range v {
			if occupied {
				count += 1
			}
		}
	}
	return count
}
