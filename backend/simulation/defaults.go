package simulation

func InitYear14() SimInput {
	AnzahlMaschinenJahr13 := 24
	AnzahlMaschinenJahr14 := 23
	AnzahlProduktionsräume := 3
	AnzahlLagerräume := 1

	m := Maschinen{}
	for i := 0; i < AnzahlMaschinenJahr14; i++ {
		m.Maschinen = append(m.Maschinen, InitMaschine(14, false, true))
	}
	for i := 0; i < AnzahlMaschinenJahr13; i++ {
		m.Maschinen = append(m.Maschinen, InitMaschine(13, true, false))
	}

	p := Produktionsräume{}
	for i := 0; i < AnzahlProduktionsräume; i++ {
		p.Produktionsräume = append(p.Produktionsräume, NewProduktionsraum(9))
	}

	l := Lagerräume{}
	for i := 0; i < AnzahlLagerräume; i++ {
		l.Lagerräume = append(l.Lagerräume, NewLagerraum(9))
	}

	simInput := SimInput{
		Jahr:                                           14.0,
		Produktionsmenge:                               690_000.0,
		Materialverbrauch:                              0.888, //  0.888 // 1.2
		MaterialstufeÖkoUndQualität:                    1.0,
		VerkaufspreisInland:                            130.0,
		VerkaufspreisAusland:                           140.80, // Changed to euros from usd
		ProduktionsräumeInBesitz:                       p,
		MaschinenInBesitz:                              m,
		LagerräumeInBesitz:                             l,
		AnzahlProduktionsmitarbeiter:                   92.4,
		LohnProduktionsmitarbeiter:                     63_500.0,
		AnzahlMarketingMitarbeiterInland:               10,
		AnzahlMarketingMitarbeiterAusland:              3,
		GehaltMarketingMitarbeiterInland:               80_500.0,
		GehaltMarketingMitarbeiterAusland:              89_000.0,
		AnzahlAuszubildende:                            1.0,
		GehaltGeschäftsleitung:                         6.0,
		StückverbrauchEnergieEinheiten:                 1.1,
		Lageranfangsbestand:                            29_051,
		WerbungAllgemeinInland:                         2_000_000.0,
		WerbungOnlineshopInland:                        1_000_000.0,
		WerbungAllgemeinAusland:                        1_000_000.0,
		WerbungOnlineshopAusland:                       500_000.0,
		AbsatzmengeInlandDetailhandelMittlereQualität:  0.0,
		AbsatzmengeInlandDetailhandelHoheQualität:      193_200.0, // 205_677
		AbsatzmengeInlandOnlineshopMittlereQualität:    0.0,
		AbsatzmengeInlandOnlineshopHoheQualität:        289_800.0, // 263_107
		AbsatzmengeAuslandDetailhandelHoheQualität:     82_800,    //  87_299
		AbsatzmengeAuslandDetailhandelMittlereQualität: 0.0,
		AbsatzmengeAuslandOnlineshopMittlereQualität:   0.0,
		AbsatzmengeAuslandOnlineshopHoheQualität:       135_872, // 124_200
		AufwandForschungEntwicklung:                    5_000_000.0,
		StückWeiterbildungProduktionsmitarbeiter:       10_000.0,
		StückWeiterbildungMarketingpersonal:            13_000.0,
		HöheLangfristigeBankkredite:                    8_050_000.0,
		HöheÜberbrückungskredite:                       0.0,
		ÜbrigerBetrieblicherAufwand:                    189_000.0,
	}
	return simInput
}

type ResultsYear struct {
	Materialaufwand             float64
	Produktionsaufwand          float64
	ÜbrigerPersonalaufwand      float64
	Werbeaufwand                float64
	Distributionsaufwand        float64
	Lageraufwand                float64
	ForschungUndEntwicklung     float64
	Weiterbildungsaufwand       float64
	Verwaltungsaufwand          float64
	ÜbrigerBetrieblicherAufwand float64
	Abschreibungen              float64
	Zinsaufwand                 float64
	Steuern                     float64
	Total                       float64
}

func ExpectedValuesYear14() ResultsYear {
	return ResultsYear{
		Materialaufwand:             39.39,
		Produktionsaufwand:          19.04,
		ÜbrigerPersonalaufwand:      6.52,
		Werbeaufwand:                6.72,
		Distributionsaufwand:        19.26,
		Lageraufwand:                0.32,
		ForschungUndEntwicklung:     7.23,
		Weiterbildungsaufwand:       2.49,
		Verwaltungsaufwand:          5.68,
		ÜbrigerBetrieblicherAufwand: 0.27,
		Abschreibungen:              0.26,
		Zinsaufwand:                 0.46,
		Steuern:                     6.36,
		Total:                       113.99,
	}
}

func initYear11() SimInput {
	AnzahlMaschinenJahr9 := 24
	AnzahlMaschinenJahr11 := 3
	AnzahlProduktionsräume := 3
	AnzahlLagerräume := 1

	m := Maschinen{}
	for i := 0; i < AnzahlMaschinenJahr9; i++ {
		m.Maschinen = append(m.Maschinen, InitMaschine(9, false, false))
	}
	for i := 0; i < AnzahlMaschinenJahr11; i++ {
		m.Maschinen = append(m.Maschinen, InitMaschine(11, false, true))
	}

	p := Produktionsräume{}
	for i := 0; i < AnzahlProduktionsräume; i++ {
		p.Produktionsräume = append(p.Produktionsräume, NewProduktionsraum(9))
	}

	l := Lagerräume{}
	for i := 0; i < AnzahlLagerräume; i++ {
		l.Lagerräume = append(l.Lagerräume, NewLagerraum(9))
	}

	simInput := SimInput{
		Jahr:                                           11,
		Produktionsmenge:                               395_912.0, //405_000.0
		Materialverbrauch:                              1.0,
		MaterialstufeÖkoUndQualität:                    2,
		VerkaufspreisInland:                            96.50,
		VerkaufspreisAusland:                           0.0,
		ProduktionsräumeInBesitz:                       p,
		MaschinenInBesitz:                              m,
		LagerräumeInBesitz:                             l,
		AnzahlProduktionsmitarbeiter:                   129.7,
		LohnProduktionsmitarbeiter:                     62_000.0,
		AnzahlMarketingMitarbeiterInland:               6.0,
		AnzahlMarketingMitarbeiterAusland:              0.0,
		GehaltMarketingMitarbeiterInland:               80_500.0,
		GehaltMarketingMitarbeiterAusland:              0.0,
		AnzahlAuszubildende:                            7.0,
		GehaltGeschäftsleitung:                         100_000.0,
		StückverbrauchEnergieEinheiten:                 1.0,
		Lageranfangsbestand:                            13_258.0,
		WerbungAllgemeinInland:                         900_000.0,
		WerbungOnlineshopInland:                        140_000.0,
		WerbungAllgemeinAusland:                        0.0,
		WerbungOnlineshopAusland:                       0.0,
		AbsatzmengeInlandDetailhandelMittlereQualität:  182_421.0, //209_129.0
		AbsatzmengeInlandDetailhandelHoheQualität:      0.0,
		AbsatzmengeInlandOnlineshopMittlereQualität:    178_100.0, //209_129.0
		AbsatzmengeInlandOnlineshopHoheQualität:        0.0,
		AbsatzmengeAuslandDetailhandelHoheQualität:     0.0,
		AbsatzmengeAuslandDetailhandelMittlereQualität: 0.0,
		AbsatzmengeAuslandOnlineshopMittlereQualität:   0.0,
		AbsatzmengeAuslandOnlineshopHoheQualität:       0.0,
		AufwandForschungEntwicklung:                    900_000.0,
		StückWeiterbildungProduktionsmitarbeiter:       1_000.0,
		StückWeiterbildungMarketingpersonal:            5_000.0,
		HöheLangfristigeBankkredite:                    3_300_000.0,
		HöheÜberbrückungskredite:                       0.0,
		ÜbrigerBetrieblicherAufwand:                    31_000.0,
	}

	return simInput
}

func initYear10() SimInput {
	Jahr := 10
	AnzahlMaschinen := 24
	AnzahlProduktionsräume := 3
	AnzahlLagerräume := 1

	m := Maschinen{}
	for i := 0; i < AnzahlMaschinen; i++ {
		m.Maschinen = append(m.Maschinen, InitMaschine(9, false, false))
	}

	p := Produktionsräume{}
	for i := 0; i < AnzahlProduktionsräume; i++ {
		p.Produktionsräume = append(p.Produktionsräume, NewProduktionsraum(9))
	}

	l := Lagerräume{}
	for i := 0; i < AnzahlLagerräume; i++ {
		l.Lagerräume = append(l.Lagerräume, NewLagerraum(9))
	}

	simInput := SimInput{
		Jahr:                                           Jahr,
		Produktionsmenge:                               350_000.0,
		Materialverbrauch:                              1.0,
		MaterialstufeÖkoUndQualität:                    2,
		VerkaufspreisInland:                            100.0,
		VerkaufspreisAusland:                           0.0,
		ProduktionsräumeInBesitz:                       p,
		MaschinenInBesitz:                              m,
		LagerräumeInBesitz:                             l,
		AnzahlProduktionsmitarbeiter:                   118.2,
		LohnProduktionsmitarbeiter:                     60_000,
		AnzahlMarketingMitarbeiterInland:               6,
		AnzahlMarketingMitarbeiterAusland:              0.0,
		GehaltMarketingMitarbeiterInland:               80_000.0,
		GehaltMarketingMitarbeiterAusland:              0.0,
		AnzahlAuszubildende:                            9.0,
		GehaltGeschäftsleitung:                         100_000.0,
		StückverbrauchEnergieEinheiten:                 1.0,
		Lageranfangsbestand:                            13_258.0,
		WerbungAllgemeinInland:                         1_000_000.0,
		WerbungOnlineshopInland:                        150_000.0,
		WerbungAllgemeinAusland:                        0.0,
		WerbungOnlineshopAusland:                       0.0,
		AbsatzmengeInlandDetailhandelMittlereQualität:  177_577.0,
		AbsatzmengeInlandDetailhandelHoheQualität:      0.0,
		AbsatzmengeInlandOnlineshopMittlereQualität:    172_423.0,
		AbsatzmengeInlandOnlineshopHoheQualität:        0.0,
		AbsatzmengeAuslandDetailhandelHoheQualität:     0.0,
		AbsatzmengeAuslandDetailhandelMittlereQualität: 0.0,
		AbsatzmengeAuslandOnlineshopMittlereQualität:   0.0,
		AbsatzmengeAuslandOnlineshopHoheQualität:       0.0,
		AufwandForschungEntwicklung:                    900_000.0,
		StückWeiterbildungProduktionsmitarbeiter:       1_000.0,
		StückWeiterbildungMarketingpersonal:            5_000.0,
		HöheLangfristigeBankkredite:                    3_300_000,
		HöheÜberbrückungskredite:                       0.0,
		ÜbrigerBetrieblicherAufwand:                    33_000.0,
	}
	return simInput
}
