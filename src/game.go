package power4

import "net/http"

func GamePower4(w http.ResponseWriter, r *http.Request, NbTurn int, infosGameVar *InfosGame) {
	if NbTurn%2 != 0 {
		infosGameVar.CurrentPlayer = 1
		infosGameVar.WaitingPlayer = 2
	} else {
		infosGameVar.CurrentPlayer = 2
		infosGameVar.WaitingPlayer = 1
	}

	y = 0

	switch infosGameVar.ColPlayed {
	case "Col0":
		y = 0
	case "Col1":
		y = 1
	case "Col2":
		y = 2
	case "Col3":
		y = 3
	case "Col4":
		y = 4
	case "Col5":
		y = 5
	case "Col6":
		y = 6
	}

	var x int
	for x = xInit; x >= 0; x-- {
		if infosGameVar.Board[x][y] == 0 {
			infosGameVar.Board[x][y] = infosGameVar.CurrentPlayer
			break
		}
	}
	checkWin(w, r, infosGameVar)

}

func checkWin(w http.ResponseWriter, r *http.Request, infoGameVar *InfosGame) {
	var VerifVal int
	var VerifDraw bool

	for i := 0; i <= xInit; i++ {
		for j := 0; j <= yInit; j++ {
			VerifVal = infoGameVar.Board[i][j]

			if j+3 <= yInit && //Verif horizontal
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i][j+1] &&
				VerifVal == infoGameVar.Board[i][j+2] &&
				VerifVal == infoGameVar.Board[i][j+3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/end", http.StatusSeeOther)
			}

			if i+3 <= xInit && //Verif vertical
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i+1][j] &&
				VerifVal == infoGameVar.Board[i+2][j] &&
				VerifVal == infoGameVar.Board[i+3][j] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/end", http.StatusSeeOther)
			}

			if i-3 >= 0 && j-3 >= 0 && //Verif diago bas droit --> haut gauche
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i-1][j-1] &&
				VerifVal == infoGameVar.Board[i-2][j-2] &&
				VerifVal == infoGameVar.Board[i-3][j-3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/end", http.StatusSeeOther)
			}

			if i-3 >= 0 && j+3 <= yInit && //Verif diago bas gauche --> haut droite
				infoGameVar.Board[i][j] != 0 &&
				VerifVal == infoGameVar.Board[i-1][j+1] &&
				VerifVal == infoGameVar.Board[i-2][j+2] &&
				VerifVal == infoGameVar.Board[i-3][j+3] {
				if infoGameVar.Board[i][j] == 1 {
					infoGameVar.Winner = infoGameVar.Player1
				} else {
					infoGameVar.Winner = infoGameVar.Player2
				}
				http.Redirect(w, r, "/end", http.StatusSeeOther)
			}
		}
	}
	for CptDraw := 0; CptDraw <= yInit; CptDraw++ {
		if infoGameVar.Board[0][CptDraw] != 0 {
			VerifDraw = true
		} else {
			VerifDraw = false
			break
		}
	}
	if VerifDraw {
		infoGameVar.Winner = "Draw"
		http.Redirect(w, r, "/end", http.StatusSeeOther)
	}
}
