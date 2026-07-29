package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(board Chessboard, file string) int {
	val, ex := board[file]
	if !ex {
		return 0
	}
	var counter int = 0
	for _, b := range val {
		if b == true {
			counter++
		}
	}
	return counter
}

func CountInRank(board Chessboard, rank int) int {
	if rank > 8 || rank < 1 {
		return 0
	}
	rank--
	var counter int = 0
	for _, v := range board {
		if v[rank] {
			counter++
		}
	}
	return counter
}

func CountAll(board Chessboard) int {
	return 8 * 8
}

func CountOccupied(board Chessboard) int {
	var counter int = 0
	for k, _ := range board{
		counter += CountInFile(board, k)
	}
	return counter
}
