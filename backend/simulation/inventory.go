package simulation

import (
	"log"
	"math"
)

// TODO: Add booth into one struct to remove duplicate logic: DRY

type Maschinen struct {
	Maschinen []*Maschine
}

type Maschine struct {
	Anschaffungsjahr int
	// map[AktuellesJahr]Wert
	BilanzwertJahr map[int]float64
	sell           bool
	buy            bool
}

func InitMaschine(Anschaffungsjahr int, sell bool, buy bool) *Maschine {

	Bilanzwerte := map[int]map[int]float64{
		9:  {9: 104075.0, 10: 93_700.0, 11: 84_300.0, 12: 75_900.0, 13: 68_300.0, 14: 61_500.0},
		10: {9: 200_000.0, 10: 180_000.0, 11: 162_000.0, 12: 145_800.0, 13: 131_200.0, 14: 118_100.0},
		11: {10: 210_000.0, 11: 189_000.0, 12: 170_100.0, 13: 153_100.0, 14: 137_800.0},
		12: {11: 350_000.0, 12: 315_000.0, 13: 283_500.0, 14: 255_200.0},
		13: {12: 400_000.0, 13: 360_000.0, 14: 324_000.0},
		14: {13: 500_000.0, 14: 450_000.0},
	}

	BilanzwerteMaschine, ok := Bilanzwerte[Anschaffungsjahr]
	if !ok {
		log.Fatal("Bilanzwert für maschine nicht gefunden für anschaffungsjahr: ", Anschaffungsjahr)
	}

	return &Maschine{
		Anschaffungsjahr: Anschaffungsjahr,
		BilanzwertJahr:   BilanzwerteMaschine,
		sell:             sell,
		buy:              buy,
	}
}

func (m *Maschinen) GetAbschreibehöhe(Jahr int) float64 {
	gesamteAbschreibehöhe := 0.0

	for _, maschine := range m.Maschinen {
		if maschine.sell {
			continue
		}

		gesamteAbschreibehöhe += maschine.BilanzwertJahr[Jahr-1] - maschine.BilanzwertJahr[Jahr]
	}

	return gesamteAbschreibehöhe
}

func (m *Maschinen) GetAnzahlMaschinen() int {
	anzahlMaschinen := 0
	for _, maschine := range m.Maschinen {
		if maschine.sell {
			continue
		}
		anzahlMaschinen++
	}
	return anzahlMaschinen
}

type Produktionsräume struct {
	Produktionsräume []*Produktionsraum
}

type Produktionsraum struct {
	Anschaffungsjahr  int
	AnfangsBilanzwert float64
}

func NewProduktionsraum(Anschaffungsjahr int) *Produktionsraum {
	Anschaffungskosten := map[int]float64{
		9:  1_550_877,
		10: 4_000_000,
		11: 4_200_000,
		12: 4_200_000,
		13: 4_200_000,
		14: 4_200_000,
	}
	kosten, ok := Anschaffungskosten[Anschaffungsjahr]
	if !ok {
		log.Fatal("Anschaffungsjahr für Produktionsraum in Map nicht enthalten!")
	}

	return &Produktionsraum{
		Anschaffungsjahr:  Anschaffungsjahr,
		AnfangsBilanzwert: kosten,
	}
}

func (p *Produktionsräume) GetAbschreibehöhe(Jahr int) float64 {
	gesamteAbschreibehöhe := 0.0
	for _, produktionsraum := range p.Produktionsräume {
		// TODO: add check to exclude Produktionsräume being sold from calculation
		gesamteAbschreibehöhe += produktionsraum.AnfangsBilanzwert * math.Pow(1-AbschreibewertProduktionsraumProzent, float64(Jahr-produktionsraum.Anschaffungsjahr)) * AbschreibewertProduktionsraumProzent
	}
	return gesamteAbschreibehöhe
}

func (p *Produktionsräume) GetAnzahlProduktionsräume() int {
	return len(p.Produktionsräume)
}

type Lagerräume struct {
	Lagerräume []*Lagerraum
}

type Lagerraum struct {
	Anschaffungsjahr  int
	AnfangsBilanzwert float64
}

func NewLagerraum(Anschaffungsjahr int) *Lagerraum {
	Anschaffungskosten := map[int]float64{
		9:  1_031_578,
		10: 2_000_000,
		11: 2_100_000,
		12: 2_100_000,
		13: 2_100_000,
		14: 2_100_000,
	}
	kosten, ok := Anschaffungskosten[Anschaffungsjahr]
	if !ok {
		log.Fatal("Anschaffungsjahr für Lagerraum in Map nicht enthalten!")
	}

	return &Lagerraum{
		Anschaffungsjahr:  Anschaffungsjahr,
		AnfangsBilanzwert: kosten,
	}
}

func (l *Lagerräume) GetAbschreibehöhe(Jahr int) float64 {
	gesamteAbschreibehöhe := 0.0
	for _, lagerraum := range l.Lagerräume {
		// TODO: add check to exclude Lagerräume being sold from calculation
		gesamteAbschreibehöhe += lagerraum.AnfangsBilanzwert * math.Pow(1-AbschreibewertLagerraumProzent, float64(Jahr-lagerraum.Anschaffungsjahr)) * AbschreibewertLagerraumProzent
	}
	return gesamteAbschreibehöhe
}

func (l *Lagerräume) GetAnzahlLagerräume() int {
	return len(l.Lagerräume)
}
