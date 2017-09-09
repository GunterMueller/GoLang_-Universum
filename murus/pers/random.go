package pers

// (c) Christian Maurer   v. 140217 - license see murus.go

import (
  "murus/rand"
  "murus/text"
)
const
  N = 200
var (
  NN, VN [N]text.Text
  nn int
)


func (x *person) Rand() {
  x.surname.Copy (NN[rand.Natural (N)])
  n:= rand.Natural (N)
  x.firstName.Copy (VN[n])
  x.TruthValue.Set (n < N / 2)
  x.Calendarday.Randomize()
}


func ng (S string) {
  NN[nn].Defined (S)
  nn++
}


func vg (S string) {
  VN[nn].Defined (S)
  nn++
}


func init1() {
  for n:= 0; n < N; n++ {
    NN[n] = text.New (lenName)
    VN[n] = text.New (lenFirstName)
  }
  nn = 0
  ng ("Aal");       ng ("Adler");     ng ("Affe");      ng ("Albatros");  ng ("Ameise")
  ng ("Ameisenbär");ng ("Amsel");     ng ("Antilope");  ng ("Assel");     ng ("Barsch")
  ng ("Bär");       ng ("Biber");     ng ("Biene");     ng ("Bilch");     ng ("Blindschleiche")
  ng ("Büffel");    ng ("Bonobo");    ng ("Braunbär");  ng ("Bulle");     ng ("Bulldogge")
  ng ("Chinchilla");ng ("Chamäleon"); ng ("Dackel");    ng ("Delphin");   ng ("Dogge")
  ng ("Dromedar");  ng ("Drossel");   ng ("Echse");     ng ("Ente");      ng ("Eidechse")
  ng ("Eisbär");    ng ("Elch");      ng ("Elefant");   ng ("Emu");       ng ("Erdmännchen")
  ng ("Erpel");     ng ("Esel");      ng ("Eule");      ng ("Faultier");  ng ("Fink")
  ng ("Fisch");     ng ("Flamingo");  ng ("Fledermaus");ng ("Fliege");    ng ("Flußpferd")
  ng ("Frosch");    ng ("Fuchs");     ng ("Gans");      ng ("Gazelle");   ng ("Gecko")
  ng ("Gemse");     ng ("Geier");     ng ("Gibbon");    ng ("Giraffe");   ng ("Gorilla")
  ng ("Gnu");       ng ("Gürteltier");ng ("Hai");       ng ("Hamster");   ng ("Hase")
  ng ("Hecht");     ng ("Hering");    ng ("Hornisse");  ng ("Huhn");      ng ("Hummel")
  ng ("Hund");      ng ("Hirsch");    ng ("Hirschkäfer");ng("Hutaffe");   ng ("Hyäne")
  ng ("Ibis");      ng ("Igel");      ng ("Iltis");     ng ("Jaguar");    ng ("Kapuziner")
  ng ("Käfer");     ng ("Känguruh");  ng ("Kamel");     ng ("Kater");     ng ("Katze")
  ng ("Koala");     ng ("Kojote");    ng ("Kormoran");  ng ("Kranich");   ng ("Krebs")
  ng ("Kröte");     ng ("Krokodil");  ng ("Kuh");       ng ("Krabbe");    ng ("Krähe")
  ng ("Kreuzotter");ng ("Lachs");     ng ("Languste");  ng ("Laubfrosch");ng ("Leopard")
  ng ("Lerche");    ng ("Lippenbär"); ng ("Löwe");      ng ("Lori");      ng ("Luchs")
  ng ("Lurch");     ng ("Maikäfer");  ng ("Maki");      ng ("Makrele");   ng ("Malaienbär")
  ng ("Manguste");  ng ("Marder");    ng ("Maultier");  ng ("Maulwurf");  ng ("Maus")
  ng ("Meise");     ng ("Meerkatze"); ng ("Meerschwein");ng("Möwe");      ng ("Molch")
  ng ("Mücke");     ng ("Mungo");     ng ("Muschel");   ng ("Nachtaffe"); ng ("Nasenaffe")
  ng ("Nashorn");   ng ("Natter");    ng ("Nilpferd");  ng ("Ochse");     ng ("Olm")
  ng ("OrangUtan"); ng ("Ozelot");    ng ("Otter");     ng ("Panther");   ng ("Panda")
  ng ("Pelikan");   ng ("Pfau");      ng ("Pferd");     ng ("Pinguin");   ng ("Pinscher")
  ng ("Pirol");     ng ("Pottwal");   ng ("Pudel");     ng ("Puma");      ng ("Qualle")
  ng ("Rabe");      ng ("Ratte");     ng ("Reh");       ng ("Reiher");    ng ("Rentier")
  ng ("Rochen");    ng ("Rotfuchs");  ng ("Salamander");ng ("Schaf");     ng ("Schäferhund")
  ng ("Schakal");   ng ("Schellfisch");ng("Schimpanse");ng ("Schlange");  ng ("Schleiche")
  ng ("Schildkröte");ng("Schnecke");  ng ("Schnepfe");  ng ("Schwalbe");  ng ("Schwan")
  ng ("Schwein");   ng ("Seekuh");    ng ("Seelöwe");   ng ("Seepferdchen"); ng ("Seezunge")
  ng ("Spatz");     ng ("Specht");    ng ("Spinne");    ng ("Stachelschwein"); ng ("Star")
  ng ("Steinbock"); ng ("Stier");     ng ("Stinktier"); ng ("Storch");    ng ("Strauß")
  ng ("Tapir");     ng ("Taube");     ng ("Terrier");   ng ("Tiger");     ng ("Tukan")
  ng ("Uhu");       ng ("Viper");     ng ("Vogel");     ng ("Wal");       ng ("Walroß")
  ng ("Wapiti");    ng ("Waran");     ng ("Wasserbock");ng ("Wiesel");    ng ("Wolf")
  ng ("Wildkatze"); ng ("Wildpferd"); ng ("Wisent");    ng ("Wühlmaus");  ng ("Wurm")
  ng ("Yak");       ng ("Zebra");     ng ("Zaunkönig"); ng ("Ziege");     ng ("Zwergohreule")

  nn = 0
  vg ("Adelheid");  vg ("Angelika");  vg ("Anita");     vg ("Anke");      vg ("Anna")
  vg ("Annemarie"); vg ("Anneliese"); vg ("Annette");   vg ("Antje");     vg ("Antonia")
  vg ("Barbara");   vg ("Beate");     vg ("Beatrix");   vg ("Berta");     vg ("Bettina")
  vg ("Bianca");    vg ("Birgit");    vg ("Brigitte");  vg ("Carola");    vg ("Christa")
  vg ("Christina"); vg ("Claudia");   vg ("Corinna");   vg ("Cornelia");  vg ("Dagmar")
  vg ("Dora");      vg ("Doris");     vg ("Edith");     vg ("Elfriede");  vg ("Elisabeth")
  vg ("Elke");      vg ("Ellen");     vg ("Elsa");      vg ("Elvira");    vg ("Emma")
  vg ("Erika");     vg ("Erna");      vg ("Eva");       vg ("Gabriele");  vg ("Gerda")
  vg ("Gisela");    vg ("Gudrun");    vg ("Hanna");     vg ("Hannelore"); vg ("Heide")
  vg ("Heike");     vg ("Helene");    vg ("Helga");     vg ("Herta");     vg ("Hilde")
  vg ("Ilse");      vg ("Ines");      vg ("Ingeborg");  vg ("Ingrid");    vg ("Irene")
  vg ("Iris");      vg ("Irmgard");   vg ("Jeanette");  vg ("Karin");     vg ("Karla")
  vg ("Katja");     vg ("Kathi");     vg ("Katrin");    vg ("Kerstin");   vg ("Lara")
  vg ("Luise");     vg ("Margarete"); vg ("Margot");    vg ("Maria");     vg ("Marion")
  vg ("Marianna");  vg ("Martha");    vg ("Martina");   vg ("Michaela");  vg ("Monika")
  vg ("Nadja");     vg ("Nicole");    vg ("Paula");     vg ("Petra");     vg ("Regina")
  vg ("Renate");    vg ("Rita");      vg ("Rosemarie"); vg ("Ruth");      vg ("Sabine")
  vg ("Sara");      vg ("Sigrid");    vg ("Simone");    vg ("Sonja");     vg ("Stefanie")
  vg ("Susanne");   vg ("Sylvia");    vg ("Tamara");    vg ("Tanja");     vg ("Tina")
  vg ("Ulrike");    vg ("Ursula");    vg ("Ute");       vg ("Verena");    vg ("Waltraud")

  vg ("Achim");     vg ("Adalbert");  vg ("Albert");    vg ("Alexander"); vg ("Alfred")
  vg ("Andreas");   vg ("Axel");      vg ("Armin");     vg ("Bernd");     vg ("Bernhard")
  vg ("Bodo");      vg ("Christian"); vg ("Christoph"); vg ("Cornelius"); vg ("Daniel")
  vg ("David");     vg ("Detlef");    vg ("Dieter");    vg ("Dietmar");   vg ("Dirk")
  vg ("Eckhard");   vg ("Egon");      vg ("Erich");     vg ("Ernst");     vg ("Erwin")
  vg ("Ewald");     vg ("Felix");     vg ("Florian");   vg ("Frank");     vg ("Fritz")
  vg ("Gerd");      vg ("Gerhard");   vg ("Georg");     vg ("Gottlieb");  vg ("Gregor")
  vg ("Günter");    vg ("Hans");      vg ("Harald");    vg ("Hartwig");   vg ("Helmut")
  vg ("Heinrich");  vg ("Heinz");     vg ("Henry");     vg ("Herbert");   vg ("Holger")
  vg ("Horst");     vg ("Ingo");      vg ("Joachim");   vg ("Johannes");  vg ("Jörg")
  vg ("Joseph");    vg ("Jürgen");    vg ("Karl");      vg ("Karsten");   vg ("Klaus")
  vg ("Knuth");     vg ("Konrad");    vg ("Kurt");      vg ("Lars");      vg ("Lothar")
  vg ("Ludwig");    vg ("Manfred");   vg ("Marco");     vg ("Markus");    vg ("Martin")
  vg ("Matthias");  vg ("Max");       vg ("Michael");   vg ("Nils");      vg ("Norbert")
  vg ("Olaf");      vg ("Oliver");    vg ("Otto");      vg ("Paul");      vg ("Peter")
  vg ("Philipp");   vg ("Raimund");   vg ("Rainer");    vg ("Ralf");      vg ("Richard")
  vg ("Robert");    vg ("Roland");    vg ("Rolf");      vg ("Ronald");    vg ("Rudolf")
  vg ("Siegfried"); vg ("Sascha");    vg ("Stefan");    vg ("Sven");      vg ("Tim")
  vg ("Thomas");    vg ("Torsten");   vg ("Udo");       vg ("Ulrich");    vg ("Uwe")
  vg ("Walter");    vg ("Werner");    vg ("Wilhelm");   vg ("Winfried");  vg ("Wolfgang")
}
