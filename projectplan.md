# MyAssets
-	En applikation för ett samlande sparande, vare sig det rör sig om aktier, bilar, klockor eller tavlor, etc. 

## Målsättning
Det är svårt att hålla sig uppdaterad på alla ens tillgångar (assets) utan att lägga ner tid och energi på det, det här verktyget skulle vara till för dels att ge en överblick, dels att göra det enklare att spekulera i annat än immateriella egendomar.
Den primära målsättningen är att ge överblick över ens tillgångar, dess ungefärliga värde som användaren själv får undersöka och ange.
Vidare målsättningar kan vara att hämta dagsaktuellt värde på t.ex. klockor, liknande bilar, tavlor, etc (kommer dock sannolikt att vara för avancerat).

## Scope/omfattning av tänkbara features
### Börsspekulation
Min id´e var från början att bygga en börshandelplatform för att underlätta aktie- och optionshandel, eftersom det är min huvudsakliga inkomst. Ur den aspekten vill jag bygga en applikation som kan spåra flertalet portföljer, göra det möjligt att slå dem samman vare sig kontot ligger på ett konto i Sverige, Schweiz eller Usa, se statistik, etc (vilket är bara ett par fåtal möjligheter jag kan tänka mig). Främst vill jag inkludera funktionen Simple Moving Averages till att börja med och sedan andra mer kortsiktiga indikationer. Sådana applikationer kostar vanligtvis 10k+/år.

### Övrig spekulation
En utveckling av originalkonceptet, bara att i stället för att spåra aktier kan man spåra aktuella värden av andra saker och immateriella egendomar. Sådana kan vara bilar, klockor, tavlor, limiterade upplagor, fastigheter och innehavda skuldebrev. 

### Översikt (huvudsaklig feature)
Den huvudsakliga poängen är att få bättre översikt över sina samlade tillgångar, i vilken form de nu kan tänkas vara.
Bifogat här är ett exempel på hur en person skulle få översikt över sina tillgångar. Vad jag märkte var att det blev svårt att se de mindre posterna när det finns väldigt stora poster.
Säg till exempel att man har 100 miljoner i aktiekapital så blir det svårt att få översikt över till exempel ens klocksamling för en bråkdel av detsamma.

![alt text](https://i.gyazo.com/12c4071c86695ddb49373cb9be0f1099.png "graph1")

## Tekniskt underbyggande
Jag vill bygga plattformen med hjälp av möjligtvis Python, Ruby och PHP/MySQL. Till exempel har jag hittat libraries för python med algoritmer för börshandel.

![alt text](https://i.gyazo.com/d36071bb987bb17ab58e38f4750c432e.png "graph2")

Med reservation för att det kan bli relativt mycket simplare än illustrerat, inom ramen för vad som faller inom kursen.

## Milstolpar
### Vecka 1
Sammanställning av alla resurser som kan tänkas behövas och färdig skiss på hur det kommer att fungera.
### Vecka 2 
En fungerande databas för inmatning av användarens tillgångar.
### Vecka 3 
Implementation av ett grafiskt index.
### Vecka 4 
Färdigställande.
