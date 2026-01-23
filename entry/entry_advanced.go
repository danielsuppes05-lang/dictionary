package entry

import "strings"

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {

	// manueller Ansatz
	// komma := -1
	// kommaz := 0
	//
	// for i, c := range pair {
	// 	if c == ',' {
	// 		komma = i
	// 		kommaz++
	// 		if kommaz > 1 {
	// 			return Empty()
	// 		}
	// 	}
	// }

	// Mögliche Vorab-Prüfung auf Anzahl der Kommas.
	// (teuer)
	// if strings.Count(pair, ",") != 1 {
	// 	return Empty()
	// }

	parts := strings.Split(pair, ",")

	// oder das?
	if len(parts) != 2 {
		return Empty()
	}

	return New(parts[0], parts[1])
}
