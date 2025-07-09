package simulation

import (
	"log"
	"math"
)

const (
	AbschreibewertMaschineProzent        = 0.1
	AbschreibewertProduktionsraumProzent = 0.05
	AbschreibewertLagerraumProzent       = 0.05
	SachaufwandMaschine                  = 900.0
	SachaufwandProduktionsraum           = 6_200.0
	SachaufwandVerbrauchsmaterial        = 0.04

	StückMaterialaufwand = 25.0

	LohnAuszubildende      = 12_000.0
	AnzahlGeschäftsleitung = 6.0

	ZinssatzLangfristigeKredite  = 0.04
	ZinssatzÜberbrückungsKredite = 0.08

	GezeichnetesKapital = 7_000_000
	Gewinnrücklagen     = 3_931_000
)

type SimInput struct {
	Jahr                        int
	Produktionsmenge            float64
	Materialverbrauch           float64
	MaterialstufeÖkoUndQualität int

	VerkaufspreisInland      float64
	VerkaufspreisAusland     float64
	ProduktionsräumeInBesitz Produktionsräume
	MaschinenInBesitz        Maschinen
	LagerräumeInBesitz       Lagerräume

	AnzahlProduktionsmitarbeiter      float64
	LohnProduktionsmitarbeiter        float64
	AnzahlMarketingMitarbeiterInland  float64
	AnzahlMarketingMitarbeiterAusland float64
	GehaltMarketingMitarbeiterInland  float64
	GehaltMarketingMitarbeiterAusland float64
	AnzahlAuszubildende               float64
	GehaltGeschäftsleitung            float64

	StückverbrauchEnergieEinheiten float64

	Lageranfangsbestand float64

	WerbungAllgemeinInland   float64
	WerbungOnlineshopInland  float64
	WerbungAllgemeinAusland  float64
	WerbungOnlineshopAusland float64

	AbsatzmengeInlandDetailhandelMittlereQualität float64
	AbsatzmengeInlandDetailhandelHoheQualität     float64

	AbsatzmengeInlandOnlineshopMittlereQualität float64
	AbsatzmengeInlandOnlineshopHoheQualität     float64

	AbsatzmengeAuslandDetailhandelHoheQualität     float64
	AbsatzmengeAuslandDetailhandelMittlereQualität float64

	AbsatzmengeAuslandOnlineshopMittlereQualität float64
	AbsatzmengeAuslandOnlineshopHoheQualität     float64

	AufwandForschungEntwicklung float64

	StückWeiterbildungProduktionsmitarbeiter float64
	StückWeiterbildungMarketingpersonal      float64

	HöheLangfristigeBankkredite float64
	HöheÜberbrückungskredite    float64

	ÜbrigerBetrieblicherAufwand float64
}

type SimResults struct {
	Inflationsrate                        float64
	SachaufwandProduktion                 float64
	MaterialAufwand                       float64
	LohnzusatzkostenInlandProzent         float64
	LohnzusatzkostenAuslandProzent        float64
	PersonalaufwandProduktionsmitarbeiter float64
	PersonalaufwandMarketingmitarbeiter   float64
	AnzahlVerwaltungsmitarbeiter          float64
	LohnVerwaltungsmitarbeiter            float64
	PersonalAufwandÜbrigesPersonal        float64
	EnergieEinheitKosten                  float64
	SachaufwandEnergie                    float64
	StückEntsorgungskosten                float64
	SachaufwandEntsorgung                 float64
	Werbeaufwand                          float64
	GehaltAufwandMarketing                float64
	AbsatzmengeInlandGesamt               float64
	AbsatzmengeAuslandGesamt              float64
	AbsatzmengeGesamt                     float64
	StücktransportkostenInland            float64
	StücktransportkostenAusland           float64
	TransportkostenGesamt                 float64
	LaufendeKostenOnlineshop              float64
	KostenDetailhandel                    float64
	Distributionsaufwand                  float64
	AufwandLagerkosten                    float64
	WeiterbildungsAufwand                 float64
	Umsatz                                float64
	Verwaltungskosten                     float64
	Zinsaufwand                           float64
	AbschreibungenLagerräumeUndÜbriges    float64
	AbschreibungenProduktionsanlagen      float64
	Produktionsaufwand                    float64
	Herstellkosten                        float64
	SelbstkostenOhneSteuern               float64
	Selbstkosten                          float64
	Steueraufwand                         float64
	EBIT                                  float64
	EAT                                   float64
}

type SimData struct {
	Input   SimInput
	Results SimResults
}

func (s *SimData) GetResults() SimResults {
	// TODO: current state is really error prone
	// make the fields NaN at start to get notified if a value is used before it is set
	// TODO: Check in functions if the used value is NaN
	s.Results.Inflationsrate = math.NaN()
	s.Results.SachaufwandProduktion = math.NaN()
	s.Results.MaterialAufwand = math.NaN()
	s.Results.LohnzusatzkostenInlandProzent = math.NaN()
	s.Results.LohnzusatzkostenAuslandProzent = math.NaN()
	s.Results.PersonalaufwandProduktionsmitarbeiter = math.NaN()
	s.Results.PersonalaufwandMarketingmitarbeiter = math.NaN()
	s.Results.AnzahlVerwaltungsmitarbeiter = math.NaN()
	s.Results.LohnVerwaltungsmitarbeiter = math.NaN()
	s.Results.PersonalAufwandÜbrigesPersonal = math.NaN()
	s.Results.EnergieEinheitKosten = math.NaN()
	s.Results.SachaufwandEnergie = math.NaN()
	s.Results.StückEntsorgungskosten = math.NaN()
	s.Results.SachaufwandEntsorgung = math.NaN()
	s.Results.Werbeaufwand = math.NaN()
	s.Results.GehaltAufwandMarketing = math.NaN()
	s.Results.AbsatzmengeInlandGesamt = math.NaN()
	s.Results.AbsatzmengeAuslandGesamt = math.NaN()
	s.Results.AbsatzmengeGesamt = math.NaN()
	s.Results.StücktransportkostenInland = math.NaN()
	s.Results.StücktransportkostenAusland = math.NaN()
	s.Results.TransportkostenGesamt = math.NaN()
	s.Results.LaufendeKostenOnlineshop = math.NaN()
	s.Results.KostenDetailhandel = math.NaN()
	s.Results.Distributionsaufwand = math.NaN()
	s.Results.AufwandLagerkosten = math.NaN()
	s.Results.WeiterbildungsAufwand = math.NaN()
	s.Results.Umsatz = math.NaN()
	s.Results.Verwaltungskosten = math.NaN()
	s.Results.Zinsaufwand = math.NaN()
	s.Results.AbschreibungenLagerräumeUndÜbriges = math.NaN()
	s.Results.AbschreibungenProduktionsanlagen = math.NaN()
	s.Results.Produktionsaufwand = math.NaN()
	s.Results.Herstellkosten = math.NaN()
	s.Results.SelbstkostenOhneSteuern = math.NaN()
	s.Results.Selbstkosten = math.NaN()
	s.Results.Steueraufwand = math.NaN()
	s.Results.EBIT = math.NaN()
	s.Results.EAT = math.NaN()

	s.getInflationsrate()
	s.getSachaufwandProduktion()
	s.getMaterialaufwand()
	s.getLohnzusatzkostenInlandProzent()
	s.getLohnzusatzkostenAuslandProzent()
	s.getPersonalaufwandProduktionsmitarbeiter()
	s.getPersonalaufwandMarketingmitarbeiter()
	s.getAnzahlVerwaltungsmitarbeiter()
	s.getLohnVerwaltungsmitarbeiter()
	s.getPersonalAufwandÜbrigesPersonal()
	s.getEnergieEinheitKosten()
	s.getSachaufwandEnergie()
	s.getStückEntsorgungskosten()
	s.getSachaufwandEntsorgung()
	s.getWerbeaufwand()
	s.getGehaltAufwandMarketing()
	s.getAbsatzmengeInlandGesamt()
	s.getAbsatzmengeAuslandGesamt()
	s.getAbsatzmengeGesamt()
	s.getStücktransportkostenInland()
	s.getStücktransportkostenAusland()
	s.getTransportkostenGesamt()
	s.getLaufendeKostenOnlineshop()
	s.getKostenDetailhandel()
	s.getDistributionsaufwand()
	s.getAufwandLagerkosten()
	s.getWeiterbildungsAufwand()
	s.getUmsatz()
	s.getVerwaltungskosten()
	s.getZinsaufwand()
	s.getAbschreibungenLagerräumeUndÜbriges()
	s.getAbschreibungenProduktionsanlagen()
	s.getProduktionsaufwand()
	s.getHerstellkosten()
	s.getSelbstkostenOhneSteuern()
	s.getEBIT()
	s.getSteueraufwand()
	s.getSelbstkosten()
	s.getEAT()

	return s.Results
}

func (s *SimData) getInflationsrate() {
	Inflationsraten := map[int]float64{
		10: 0.01,
		11: 0.01,
		12: 0.005,
		13: 0.01,
		14: 0.015,
	}

	inflationsrate, ok := Inflationsraten[s.Input.Jahr]
	if !ok {
		log.Fatal("Inflationsrate für Jahr wurde nicht gefunden: ", s.Input.Jahr)
	}
	s.Results.Inflationsrate = inflationsrate
}

func (s *SimData) getSachaufwandProduktion() {
	inflationsrate := s.Results.Inflationsrate
	AnzahlMaschinen := float64(s.Input.MaschinenInBesitz.GetAnzahlMaschinen())
	AnzahlProduktionsräume := float64(s.Input.ProduktionsräumeInBesitz.GetAnzahlProduktionsräume())
	SachaufwandProduktion := (AnzahlProduktionsräume*SachaufwandProduktionsraum + AnzahlMaschinen*SachaufwandMaschine + s.Input.Produktionsmenge*SachaufwandVerbrauchsmaterial) * (inflationsrate + 1.0)

	s.Results.SachaufwandProduktion = SachaufwandProduktion
}

func (s *SimData) getMaterialaufwand() {
	// Stufe = map[Jahr]Wert – nur für Öko- und Qualitätsstufen 1 & 2
	StückMaterialaufwandStufenÖkoUndQualität := map[int]map[int]float64{
		1: { // Ökologiestufe 1, Qualitätsstufe 1
			10: 39.00,
			11: 39.00,
			12: 42.00,
			13: 44.00,
			14: 45.00,
		},
		2: { // Ökologiestufe 2, Qualitätsstufe 2
			10: 25.00,
			11: 25.00,
			12: 26.00,
			13: 27.00,
			14: 29.00,
		},
	}

	StückMaterialaufwand, ok := StückMaterialaufwandStufenÖkoUndQualität[s.Input.MaterialstufeÖkoUndQualität][s.Input.Jahr]
	if !ok {
		log.Fatal("StückMaterialaufwand für Jahr wurde nicht gefunden: ", s.Input.Jahr)
	}

	result := (StückMaterialaufwand * s.Input.Materialverbrauch) * s.Input.Produktionsmenge
	s.Results.MaterialAufwand = result
}

func (s *SimData) getLohnzusatzkostenInlandProzent() {
	LohnzusatzkostenInlandProzent := map[int]float64{
		10: 0.2,
		11: 0.22,
		12: 0.24,
		13: 0.24,
		14: 0.25,
	}
	JahrLohnzusatzkostenInlandProzent, ok := LohnzusatzkostenInlandProzent[s.Input.Jahr]
	if !ok {
		log.Fatal("LohnzusatzkostenInlandProzent für Jahr nicht gefunden: ", s.Input.Jahr)
	}
	s.Results.LohnzusatzkostenInlandProzent = JahrLohnzusatzkostenInlandProzent
}

func (s *SimData) getLohnzusatzkostenAuslandProzent() {
	LohnzusatzkostenAuslandProzent := map[int]float64{
		10: 0.25,
		11: 0.27,
		12: 0.28,
		13: 0.28,
		14: 0.30,
	}
	JahrLohnzusatzkostenAuslandProzent, ok := LohnzusatzkostenAuslandProzent[s.Input.Jahr]
	if !ok {
		log.Fatal("LohnzusatzkostenAuslandProzent für Jahr nicht gefunden: ", s.Input.Jahr)
	}
	s.Results.LohnzusatzkostenAuslandProzent = JahrLohnzusatzkostenAuslandProzent
}

func (s *SimData) getPersonalaufwandProduktionsmitarbeiter() {
	result := s.Input.AnzahlProduktionsmitarbeiter * s.Input.LohnProduktionsmitarbeiter * (s.Results.LohnzusatzkostenInlandProzent + 1.0)
	s.Results.PersonalaufwandProduktionsmitarbeiter = result
}

func (s *SimData) getPersonalaufwandMarketingmitarbeiter() {
	PersonalaufwandMarketingmitarbeiterInland := s.Input.AnzahlMarketingMitarbeiterInland * s.Input.GehaltMarketingMitarbeiterInland * (s.Results.LohnzusatzkostenInlandProzent + 1.0)
	PersonalaufwandMarketingmitarbeiterAusland := s.Input.AnzahlMarketingMitarbeiterAusland * s.Input.GehaltMarketingMitarbeiterAusland * (s.Results.LohnzusatzkostenAuslandProzent + 1.0)

	result := PersonalaufwandMarketingmitarbeiterInland + PersonalaufwandMarketingmitarbeiterAusland
	s.Results.PersonalaufwandMarketingmitarbeiter = result
}

func (s *SimData) getAnzahlVerwaltungsmitarbeiter() {
	minVerwaltungsmitarbeiter := 29.0
	numVerwaltungsmitarbeiter := (s.Input.AnzahlProduktionsmitarbeiter+s.Input.AnzahlMarketingMitarbeiterInland+s.Input.AnzahlMarketingMitarbeiterAusland)*0.38 + s.Input.AnzahlAuszubildende*0.1 - AnzahlGeschäftsleitung

	if numVerwaltungsmitarbeiter < minVerwaltungsmitarbeiter {
		numVerwaltungsmitarbeiter = minVerwaltungsmitarbeiter
		return
	}
	s.Results.AnzahlVerwaltungsmitarbeiter = numVerwaltungsmitarbeiter
}

func (s *SimData) getLohnVerwaltungsmitarbeiter() {
	result := s.Input.LohnProduktionsmitarbeiter * 1.45
	s.Results.LohnVerwaltungsmitarbeiter = result
}

func (s *SimData) getPersonalAufwandÜbrigesPersonal() {
	result := (AnzahlGeschäftsleitung*s.Input.GehaltGeschäftsleitung + s.Input.AnzahlAuszubildende*LohnAuszubildende + s.Results.AnzahlVerwaltungsmitarbeiter*s.Results.LohnVerwaltungsmitarbeiter) * (s.Results.LohnzusatzkostenInlandProzent + 1.0)
	s.Results.PersonalAufwandÜbrigesPersonal = result
}

func (s *SimData) getEnergieEinheitKosten() {
	EnergieEinheitenKosten := map[int]float64{
		10: 2.5,
		11: 2.5,
		12: 2.7,
		13: 2.7,
		14: 2.7,
	}
	EnergieEinheitKosten, ok := EnergieEinheitenKosten[s.Input.Jahr]
	if !ok {
		log.Fatal("EnergieeinheitKosten nicht gefunden für Jahr: ", s.Input.Jahr)
	}
	s.Results.EnergieEinheitKosten = EnergieEinheitKosten
}

func (s *SimData) getSachaufwandEnergie() {
	s.Results.SachaufwandEnergie = s.Results.EnergieEinheitKosten * s.Input.StückverbrauchEnergieEinheiten * s.Input.Produktionsmenge
}

func (s *SimData) getStückEntsorgungskosten() {
	StückEntsorgungskostenJahre := map[int]float64{
		10: 0.5,
		11: 0.5,
		12: 1.0,
		13: 1.0,
		14: 1.2,
	}
	StückEntsorgungskosten, ok := StückEntsorgungskostenJahre[s.Input.Jahr]
	if !ok {
		log.Fatal("StückEntsorgungskosten nicht gefunden für Jahr: ", s.Input.Jahr)
	}
	s.Results.StückEntsorgungskosten = StückEntsorgungskosten
}

func (s *SimData) getSachaufwandEntsorgung() {
	s.Results.SachaufwandEntsorgung = s.Results.StückEntsorgungskosten * s.Input.Produktionsmenge
}

func (s *SimData) getWerbeaufwand() {
	s.Results.Werbeaufwand = s.Input.WerbungAllgemeinInland + s.Input.WerbungOnlineshopInland + s.Input.WerbungAllgemeinAusland + s.Input.WerbungOnlineshopAusland
}

func (s *SimData) getGehaltAufwandMarketing() {
	s.Results.GehaltAufwandMarketing = s.Input.AnzahlMarketingMitarbeiterInland*s.Input.GehaltMarketingMitarbeiterInland*(s.Results.LohnzusatzkostenInlandProzent+1.0) + s.Input.AnzahlMarketingMitarbeiterAusland*s.Input.GehaltMarketingMitarbeiterAusland*(s.Results.LohnzusatzkostenAuslandProzent+1.0)
}

func (s *SimData) getAbsatzmengeInlandGesamt() {
	s.Results.AbsatzmengeInlandGesamt = s.Input.AbsatzmengeInlandDetailhandelMittlereQualität + s.Input.AbsatzmengeInlandDetailhandelHoheQualität + s.Input.AbsatzmengeInlandOnlineshopMittlereQualität + s.Input.AbsatzmengeInlandOnlineshopHoheQualität
}

func (s *SimData) getAbsatzmengeAuslandGesamt() {
	s.Results.AbsatzmengeAuslandGesamt = s.Input.AbsatzmengeAuslandDetailhandelMittlereQualität + s.Input.AbsatzmengeAuslandDetailhandelHoheQualität + s.Input.AbsatzmengeAuslandOnlineshopMittlereQualität + s.Input.AbsatzmengeAuslandOnlineshopHoheQualität
}

func (s *SimData) getAbsatzmengeGesamt() {
	s.Results.AbsatzmengeGesamt = s.Results.AbsatzmengeInlandGesamt + s.Results.AbsatzmengeAuslandGesamt
}

func (s *SimData) getStücktransportkostenInland() {
	StücktransportkostenInlandJahr := map[int]float64{
		10: 2.5,
		11: 2.5,
		12: 3.2,
		13: 2.9,
		14: 2.9,
	}
	StücktransportkostenInland, ok := StücktransportkostenInlandJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StücktransportkostenInland nicht gefunden für Jahr: ", s.Input.Jahr)
	}
	s.Results.StücktransportkostenInland = StücktransportkostenInland
}

func (s *SimData) getStücktransportkostenAusland() {
	StücktransportkostenAuslandJahr := map[int]float64{
		10: 5.0,
		11: 5.0,
		12: 6.4,
		13: 5.8,
		14: 5.8,
	}
	StücktransportkostenAusland, ok := StücktransportkostenAuslandJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StücktransportkostenAusland nicht gefunden für Jahr: ", s.Input.Jahr)
	}
	s.Results.StücktransportkostenAusland = StücktransportkostenAusland
}

func (s *SimData) getTransportkostenGesamt() {
	s.Results.TransportkostenGesamt = s.Results.AbsatzmengeInlandGesamt*s.Results.StücktransportkostenInland + s.Results.AbsatzmengeAuslandGesamt*s.Results.StücktransportkostenAusland
}

func (s *SimData) getLaufendeKostenOnlineshop() {
	StückLaufendeKostenOnlineShopJahr := map[int]float64{
		10: 10.0,
		11: 10.5,
		12: 11.0,
		13: 11.0,
		14: 11.0,
	}
	StückLaufendeKostenOnlineShop, ok := StückLaufendeKostenOnlineShopJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückLaufendeKostenOnlineShop nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	s.Results.LaufendeKostenOnlineshop = StückLaufendeKostenOnlineShop * (s.Input.AbsatzmengeInlandOnlineshopHoheQualität + s.Input.AbsatzmengeInlandOnlineshopMittlereQualität + s.Input.AbsatzmengeAuslandOnlineshopHoheQualität + s.Input.AbsatzmengeAuslandOnlineshopMittlereQualität)
}

func (s *SimData) getKostenDetailhandel() {
	StückKostenDetailhandelMittlereQualitätInnlandProzentJahr := map[int]float64{
		10: 0.09,
		11: 0.095,
		12: 0.10,
		13: 0.10,
		14: 0.10,
	}
	StückKostenDetailhandelMittlereQualitätInnlandProzent, ok := StückKostenDetailhandelMittlereQualitätInnlandProzentJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückKostenDetailhandelMittlereQualitätInnlandProzent nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	StückKostenDetailhandelHoheQualitätInnlandProzentJahr := map[int]float64{
		10: 0.10,
		11: 0.105,
		12: 0.11,
		13: 0.11,
		14: 0.11,
	}
	StückKostenDetailhandelHoheQualitätInnlandProzent, ok := StückKostenDetailhandelHoheQualitätInnlandProzentJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückKostenDetailhandelHoheQualitätInnlandProzent nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	StückKostenDetailhandelMittlereQualitätAuslandProzentJahr := map[int]float64{
		10: 0.110,
		11: 0.115,
		12: 0.115,
		13: 0.115,
		14: 0.115,
	}
	StückKostenDetailhandelMittlereQualitätAuslandProzent, ok := StückKostenDetailhandelMittlereQualitätAuslandProzentJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückKostenDetailhandelMittlereQualitätAuslandProzent nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	StückKostenDetailhandelHoheQualitätAuslandProzentJahr := map[int]float64{
		10: 0.120,
		11: 0.125,
		12: 0.130,
		13: 0.130,
		14: 0.135,
	}
	StückKostenDetailhandelHoheQualitätAuslandProzent, ok := StückKostenDetailhandelHoheQualitätAuslandProzentJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückKostenDetailhandelHoheQualitätAuslandProzent nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	KostenDetailhandelInnland := s.Input.VerkaufspreisInland*StückKostenDetailhandelMittlereQualitätInnlandProzent*s.Input.AbsatzmengeInlandDetailhandelMittlereQualität + s.Input.VerkaufspreisInland*StückKostenDetailhandelHoheQualitätInnlandProzent*s.Input.AbsatzmengeInlandDetailhandelHoheQualität
	KostenDetailhandelAusland := s.Input.VerkaufspreisAusland*StückKostenDetailhandelMittlereQualitätAuslandProzent*s.Input.AbsatzmengeAuslandDetailhandelMittlereQualität + s.Input.VerkaufspreisAusland*StückKostenDetailhandelHoheQualitätAuslandProzent*s.Input.AbsatzmengeAuslandDetailhandelHoheQualität

	s.Results.KostenDetailhandel = KostenDetailhandelInnland + KostenDetailhandelAusland
}

func (s *SimData) getDistributionsaufwand() {
	s.Results.Distributionsaufwand = s.Results.GehaltAufwandMarketing + s.Results.TransportkostenGesamt + s.Results.LaufendeKostenOnlineshop + s.Results.KostenDetailhandel
}

func (s *SimData) getAufwandLagerkosten() {
	StückEigenlagerkostenJahr := map[int]float64{
		10: 5.8,
		11: 5.8,
		12: 5.9,
		13: 5.8,
		14: 5.6,
	}
	StückEigenlagerkosten, ok := StückEigenlagerkostenJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückEigenlagerkosten nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	StückFremdlagerkostenJahr := map[int]float64{
		10: 12.0,
		11: 12.0,
		12: 12.0,
		13: 11.8,
		14: 12.5,
	}
	StückFremdlagerkosten, ok := StückFremdlagerkostenJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("StückFremdlagerkosten nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	FixUnterhaltLagergebäude := +5_000.0 * (s.Results.Inflationsrate + 1.0)
	maxEigenlagerbestand := float64(s.Input.LagerräumeInBesitz.GetAnzahlLagerräume()) * 20_000.0
	if maxEigenlagerbestand < s.Input.Lageranfangsbestand {
		s.Results.AufwandLagerkosten = maxEigenlagerbestand*StückEigenlagerkosten + (s.Input.Lageranfangsbestand-maxEigenlagerbestand)*StückFremdlagerkosten + FixUnterhaltLagergebäude
		return
	}
	s.Results.AufwandLagerkosten = s.Input.Lageranfangsbestand*StückEigenlagerkosten + FixUnterhaltLagergebäude
}

func (s *SimData) getWeiterbildungsAufwand() {
	WeiterbildungProduktionsmitarbeiter := s.Input.AnzahlProduktionsmitarbeiter * s.Input.StückWeiterbildungProduktionsmitarbeiter
	WeiterbildungMarketingpersonal := (s.Input.AnzahlMarketingMitarbeiterInland + s.Input.AnzahlMarketingMitarbeiterAusland) * s.Input.StückWeiterbildungMarketingpersonal
	WeiterbildungÜbrigesPersonal := s.Input.StückWeiterbildungProduktionsmitarbeiter * 1.53 * (s.Results.AnzahlVerwaltungsmitarbeiter + AnzahlGeschäftsleitung + s.Input.AnzahlAuszubildende)

	s.Results.WeiterbildungsAufwand = WeiterbildungProduktionsmitarbeiter + WeiterbildungMarketingpersonal + WeiterbildungÜbrigesPersonal
}

func (s *SimData) getUmsatz() {
	s.Results.Umsatz = s.Results.AbsatzmengeInlandGesamt*s.Input.VerkaufspreisInland + s.Results.AbsatzmengeAuslandGesamt*s.Input.VerkaufspreisAusland
}

func (s *SimData) getVerwaltungskosten() {
	VerwaltungskostenProzentJahr := map[int]float64{
		10: 0.045,
		11: 0.045,
		12: 0.046,
		13: 0.046,
		14: 0.047,
	}
	VerwaltungskostenProzent, ok := VerwaltungskostenProzentJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("VerwaltungskostenProzent nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	s.Results.Verwaltungskosten = s.Results.Umsatz * VerwaltungskostenProzent
}

func (s *SimData) getZinsaufwand() {
	s.Results.Zinsaufwand = s.Input.HöheLangfristigeBankkredite*ZinssatzLangfristigeKredite + s.Input.HöheÜberbrückungskredite*ZinssatzÜberbrückungsKredite
}

func (s *SimData) getAbschreibungenLagerräumeUndÜbriges() {
	// übrige Anlagen crazy rechnung
	s.Results.AbschreibungenLagerräumeUndÜbriges = s.Input.LagerräumeInBesitz.GetAbschreibehöhe(s.Input.Jahr) + 180_000*math.Pow(0.95, float64(s.Input.Jahr-9))
}

func (s *SimData) getAbschreibungenProduktionsanlagen() {
	// TODO: add real cost of Abschreibungen der Ökologischen anlagen
	s.Results.AbschreibungenProduktionsanlagen = s.Input.MaschinenInBesitz.GetAbschreibehöhe(s.Input.Jahr) + s.Input.ProduktionsräumeInBesitz.GetAbschreibehöhe(s.Input.Jahr) + 100_000.0
}

func (s *SimData) getProduktionsaufwand() {
	s.Results.Produktionsaufwand = s.Results.PersonalaufwandProduktionsmitarbeiter + s.Results.SachaufwandProduktion + s.Results.SachaufwandEnergie + s.Results.SachaufwandEntsorgung + s.Results.AbschreibungenProduktionsanlagen
}

func (s *SimData) getHerstellkosten() {
	s.Results.Herstellkosten = s.Results.MaterialAufwand + s.Results.Produktionsaufwand
}

func (s *SimData) getSelbstkostenOhneSteuern() {
	// need to be calculated separately because of mismatch between Produktionsmenge and Absatzmenge
	Herstellkosten := ((s.Results.MaterialAufwand + s.Results.Produktionsaufwand) / s.Input.Produktionsmenge) * (s.Results.AbsatzmengeInlandGesamt + s.Results.AbsatzmengeAuslandGesamt)

	s.Results.SelbstkostenOhneSteuern = Herstellkosten + s.Results.PersonalAufwandÜbrigesPersonal + s.Results.Werbeaufwand + s.Results.Distributionsaufwand + s.Results.AufwandLagerkosten + s.Input.AufwandForschungEntwicklung + s.Results.WeiterbildungsAufwand + s.Results.Verwaltungskosten + s.Results.AbschreibungenLagerräumeUndÜbriges + s.Results.Zinsaufwand + s.Input.ÜbrigerBetrieblicherAufwand
}

func (s *SimData) getEBIT() {
	s.Results.EBIT = s.Results.Umsatz - s.Results.SelbstkostenOhneSteuern
}

func (s *SimData) getSteueraufwand() {
	ErtragssteuersatzJahr := map[int]float64{
		10: 0.20,
		11: 0.20,
		12: 0.22,
		13: 0.22,
		14: 0.22,
	}
	Ertragssteuersatz, ok := ErtragssteuersatzJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("Ertragssteuersatz nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	KapitalsteuersatzJahr := map[int]float64{
		10: 0.03,
		11: 0.03,
		12: 0.03,
		13: 0.04,
		14: 0.04,
	}
	Kapitalsteuersatz, ok := KapitalsteuersatzJahr[s.Input.Jahr]
	if !ok {
		log.Fatal("Kapitalsteuersatz nicht gefunden für Jahr: ", s.Input.Jahr)
	}

	s.Results.Steueraufwand = s.Results.EBIT*Ertragssteuersatz + (GezeichnetesKapital+Gewinnrücklagen)*Kapitalsteuersatz
}

func (s *SimData) getSelbstkosten() {
	s.Results.Selbstkosten = s.Results.SelbstkostenOhneSteuern + s.Results.Steueraufwand
}

func (s *SimData) getEAT() {
	s.Results.EAT = s.Results.EBIT - s.Results.Steueraufwand
}
