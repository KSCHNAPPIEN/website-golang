## Omschrijving
Deze Go-website heeft een database-verbinding waarmee het gegevens kan opslaan en ophalen. Er zijn drie routes beschikbaar: /, /Locatie en /Login. De /-route laat de gebruiker de beschikbare routes zien, de /Locatie-route laat de gebruiker de beschikbare landen zien en de /Login-route laat de gebruiker inloggen met zijn of haar credentials.

## Uitvoeren
Om deze Go-website uit te voeren, heb je Go 1.13 of hoger nodig. Voer de volgende stappen uit:

Clone deze repository naar je lokale machine:
Copy code
git clone https://github.com/user/repo.git
Ga naar de map waar je de repository hebt gekloond:
Copy code
cd repo
Installeer de benodigde packages:
Copy code
go get -d -v ./...
Voer het programma uit:
Copy code
go run main.go
Open een webbrowser en ga naar http://localhost:8080 om de website te bekijken.
## Extra informatie
De database-configuratie-instellingen worden uit het bestand Db.json gelezen. Zorg ervoor dat dit bestand aanwezig is en dat de instellingen correct zijn voordat je het programma uitvoert.