CREATE TABLE IF NOT EXISTS
  "ingredients" (
    "id" INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "name" VARCHAR NOT NULL,
    "created" timestamp(0) with time zone NOT NULL DEFAULT NOW()
  );

INSERT INTO 
  ingredients (name)
VALUES
  ('Amarantus'), ('Ananas'), ('Anchois'), ('Awokado'), ('Bakłażan grillowany'), ('Bakłażan'), ('Banany liofilizowane'), ('Banany'), ('Bazylia'), ('Biała fasola'), ('Boczek surowy'), ('Boczek wędzony'), ('Boczek'), ('Borówki suszone'), ('Borówki'), ('Botwinka'), ('Brokuły'), ('Brukselka'), ('Brzoskwinie'), ('Buraki'), ('Bułka tarta'), ('Bób suszony'), ('Bób świeży'), ('Camembert'), ('Cebula suszona'), ('Cebula'), ('Chia (nasiona)'), ('Chili w proszku'), ('Chili w płatkach'), ('Chleb'), ('Chorizo'), ('Chrzan'), ('Ciecierzyca konserwowa'), ('Ciecierzyca'), ('Cielęcina'), ('Cukier biały'), ('Cukier brązowy'), ('Cukier kokosowy'), ('Cukier palmowy'), ('Cukier puder'), ('Cukier trzcinowy'), ('Cukier waniliowy'), ('Cukinia'), ('Cynamon'), ('Cytryny'), ('Czarna fasola'), ('Czekolada biała'), ('Czekolada deserowa'), ('Czekolada gorzka'), ('Czekolada mleczna'), ('Czekolada rubinowa'), ('Czerwona fasola'), ('Czosnek granulowany'), ('Czosnek'), ('Daktyle'), ('Drożdże'), ('Dynia piżmowa'), ('Dynia'), ('Ekstrakt cytrynowy'), ('Ekstrakt migdałowy'), ('Ekstrakt waniliowy'), ('Erytrol'), ('Estragon'), ('Fasola adzuki'), ('Fasola biała'), ('Fasola czerwona'), ('Fasola mung'), ('Figi suszone'), ('Gałka muszkatołowa'), ('Gorgonzola'), ('Goździki'), ('Grejpfruty'), ('Groszek cukrowy'), ('Groszek mrożony'), ('Groszek ptysiowy'), ('Groszek zielony'), ('Gruszki suszone'), ('Gruszki'), ('Grzyby leśne'), ('Gęsina'), ('Halibut'), ('Herbata biała'), ('Herbata czarna'), ('Herbata zielona'), ('Herbata'), ('Imbir suszony'), ('Imbir'), ('Indyk'), ('Jabłka suszone'), ('Jabłka'), ('Jagnięcina'), ('Jagody goji'), ('Jagody'), ('Jaja ekologiczne'), ('Jaja od kur z wolnego wybiegu'), ('Jajka kacze'), ('Jajka przepiórcze'), ('Jajka'), ('Jarmuż'), ('Jogurt grecki odtłuszczony'), ('Jogurt grecki'), ('Jogurt kokosowy'), ('Jogurt naturalny'), ('Jogurt sojowy'), ('Kabanosy'), ('Kaczka'), ('Kakao surowe'), ('Kakao'), ('Kalafior'), ('Kalarepa'), ('Kalmary'), ('Kapary'), ('Kapusta biała'), ('Kapusta czerwona'), ('Kapusta kiszona'), ('Kapusta pekińska'), ('Kapusta włoska'), ('Karczochy'), ('Kardamon'), ('Karp'), ('Kasza bulgur'), ('Kasza gryczana'), ('Kasza jaglana'), ('Kasza jęczmienna'), ('Kasza kuskus'), ('Kasza manna'), ('Kasza perłowa'), ('Kasza pęczak'), ('Kasza quinoa'), ('Kasztany jadalne'), ('Kasztany pieczone'), ('Kawa rozpuszczalna'), ('Kawa ziarnista'), ('Kawa'), ('Kefir kokosowy'), ('Kefir naturalny'), ('Kefir wodny'), ('Kefir'), ('Kiełbasa'), ('Kiełki brokuła'), ('Kiełki lucerny'), ('Kiełki rzodkiewki'), ('Kiwi'), ('Kminek'), ('Kokos wiórki'), ('Kokos świeży'), ('Kolendra'), ('Komosa ryżowa'), ('Koper'), ('Krab królewski'), ('Krab'), ('Krewetki koktajlowe'), ('Krewetki tygrysie'), ('Krewetki'), ('Ksylitol'), ('Kukurydza'), ('Kurczak'), ('Kurkuma'), ('Likier kokosowy'), ('Liść laurowy'), ('Majeranek'), ('Mak'), ('Makaron'), ('Makrela'), ('Maliny'), ('Marchew'), ('Margaryna'), ('Masa kakaowa'), ('Mascarpone'), ('Masło kakaowe'), ('Masło klarowane (ghee)'), ('Masło orzechowe'), ('Masło'), ('Matcha'), ('Małże świętego Jakuba'), ('Małże'), ('Melasa'), ('Mielona wieprzowina'), ('Mielona wołowina'), ('Migdały blanszowane'), ('Migdały'), ('Miód akacjowy'), ('Miód gryczany'), ('Miód lipowy'), ('Miód rzepakowy'), ('Miód wielokwiatowy'), ('Miód'), ('Mleko kokosowe'), ('Mleko kozie'), ('Mleko migdałowe'), ('Mleko owcze'), ('Mleko owsiane'), ('Mleko ryżowe'), ('Mleko skondensowane'), ('Mleko sojowe'), ('Mleko w proszku'), ('Mleko'), ('Morele suszone'), ('Musztarda Dijon'), ('Musztarda francuska'), ('Musztarda sarepska'), ('Mąka bananowa'), ('Mąka dyniowa'), ('Mąka gryczana'), ('Mąka jaglana'), ('Mąka kasztanowa'), ('Mąka kokosowa'), ('Mąka kukurydziana'), ('Mąka migdałowa'), ('Mąka orkiszowa'), ('Mąka pełnoziarnista'), ('Mąka pszenna'), ('Mąka razowa'), ('Mąka ryżowa'), ('Mąka sojowa'), ('Mąka z ciecierzycy'), ('Mąka z soczewicy'), ('Mąka ziemniaczana'), ('Napoje roślinne smakowe'), ('Napój kombucha'), ('Nasiona chia'), ('Nasiona dyni'), ('Nasiona konopi'), ('Nasiona słonecznika'), ('Natka pietruszki'), ('Nerkowce'), ('Ocet balsamiczny'), ('Ocet gruszkowy'), ('Ocet jabłkowy'), ('Ocet kokosowy'), ('Ocet malinowy'), ('Ocet ryżowy'), ('Ocet winny'), ('Ogórki kiszone'), ('Ogórki konserwowe'), ('Ogórki'), ('Olej kokosowy'), ('Olej lniany'), ('Olej rzepakowy'), ('Olej sezamowy'), ('Olej słonecznikowy'), ('Oliwa z oliwek'), ('Oliwki czarne'), ('Oliwki zielone'), ('Oregano'), ('Orkisz'), ('Orzechy brazylijskie'), ('Orzechy laskowe prażone'), ('Orzechy laskowe'), ('Orzechy macadamia'), ('Orzechy pekan solone'), ('Orzechy pekan'), ('Orzechy piniowe'), ('Orzechy tygrysie'), ('Orzechy włoskie prażone'), ('Orzechy włoskie'), ('Orzechy ziemne'), ('Ostrygi'), ('Ośmiornica'), ('Papryczki chili'), ('Papryczki jalapeño'), ('Papryka czerwona'), ('Papryka ostra'), ('Papryka słodka'), ('Papryka zielona'), ('Parmezan'), ('Pasta curry czerwona'), ('Pasta curry zielona'), ('Pasta curry żółta'), ('Pasta miso'), ('Pasta tahini'), ('Pasternak'), ('Pasztet'),('Pestki dyni'), ('Pestki granatu'), ('Pestki słonecznika'), ('Pieczarki'), ('Pieprz biały'), ('Pieprz cayenne'), ('Pieprz czarny'), ('Pieprz różowy'), ('Pieprz syczuański'), ('Pieprz zielony'), ('Pistacje'), ('Piwo bezalkoholowe'), ('Pomarańcze'), ('Pomidory'), ('Por'), ('Proszek do pieczenia'), ('Pstrąg wędzony'), ('Pstrąg'), ('Płatki jaglane'), ('Płatki kukurydziane'), ('Płatki owsiane'), ('Płatki ryżowe'), ('Ricotta'), ('Rodzynki'), ('Rozmaryn'), ('Rukola'), ('Ryba dorsz'), ('Ryż'), ('Rzepa'), ('Rzodkiewka'), ('Salami'), ('Sandacz'), ('Sardynki'), ('Sałata'), ('Seitan'), ('Ser biały'), ('Ser burrata'), ('Ser cheddar'), ('Ser edamski'), ('Ser feta'), ('Ser halloumi'), ('Ser kozi'), ('Ser mozzarella'), ('Ser pecorino'), ('Ser pleśniowy'), ('Ser provolone'), ('Ser roquefort'), ('Ser topiony'), ('Ser żółty'), ('Serek homogenizowany'), ('Sezam'), ('Siemię lniane'), ('Skorupiaki'), ('Skrobia kukurydziana'), ('Skrobia tapiokowa'), ('Skrobia ziemniaczana'), ('Skórka cytrynowa'), ('Skórka limonki'), ('Skórka pomarańczowa'), ('Soczewica czerwona'), ('Soczewica zielona'), ('Soczewica'), ('Soda oczyszczona'), ('Soja fermentowana'), ('Soja'), ('Sok z aloesu'), ('Sok z brzozy'), ('Sos Worcestershire'), ('Sos barbecue'), ('Sos hoisin'), ('Sos ostrygowy'), ('Sos rybny'), ('Sos sojowy'), ('Sos teriyaki'), ('Stewia'), ('Suszone pomidory'), ('Syrop buraczany'), ('Syrop daktylowy'), ('Syrop klonowy'), ('Syrop miętowy'), ('Syrop z agawy'), ('Szałwia'), ('Szparagi'), ('Szpinak'), ('Szynka dojrzewająca'), ('Szynka parmeńska'), ('Sól czarna (Kala Namak)'), ('Sól himalajska'), ('Sól kłodawska'), ('Sól morska'), ('Sól wędzona'), ('Sól'), ('Słodzik erytrytol'), ('Słodzik sacharyna'), ('Słonecznik łuskany'), ('Tempeh naturalny'), ('Tempeh wędzony'), ('Tempeh'), ('Tofu jedwabiste'), ('Tofu marynowane w sosie sojowym'), ('Tofu marynowane'), ('Tofu wędzone'), ('Tofu'), ('Topinambur'), ('Truskawki liofilizowane'), ('Truskawki'), ('Tuńczyk w oleju'), ('Tuńczyk'), ('Twarożek wiejski'), ('Twaróg'), ('Tymianek'), ('Wanilia'), ('Wieprzowina'), ('Wino bezalkoholowe'), ('Winogrona'), ('Woda kokosowa'), ('Wołowina'), ('Yerba mate'), ('Ziele angielskie'), ('Ziemniaki'), ('Łosoś wędzony'), ('Łosoś'), ('Śliwki suszone'), ('Śliwki wędzone'), ('Śliwki'), ('Śmietana 18%'), ('Śmietana 30%'), ('Śmietana kokosowa'), ('Śmietanka kokosowa w proszku'), ('Śmietanka kokosowa'), ('Śmietanka ryżowa'), ('Śmietanka sojowa'), ('Żurawina suszona');
