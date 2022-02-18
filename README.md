## ErroneusDilettante
Effortless effort cmd app for adding simple string records from .csv files into a local database. **Randomized** data retrievals from stored collections are its raison d'etre.

## Development
For now the app can be started via cli interface, either by using the provided sample database (*sample/AceArtists.db*) or by creating and filling the database prior to starting the app. For now database constists of key/value independently stored names, surnames and 'reviews'. The app is really a webapp which is meant to provide user interface for all database operations. Still, there is a convenient cli interface which is used to connect to database, create it if it does not exist, and fill in the names, surnames and review stores within the database, from .csv files.
### With the sample database
```cli
go mod tidy 
go build
ErroneusDilettante -basepath ./sample/AceArtists.db -webapp
```
### Without the sample database
```cli
go mod tidy 
go build
ErroneusDilettante -basepath [database_name].db -names [path_to].csv -surnames [path_to].csv
ErroneusDilettante -basepath [database_name].db -webapp
```

*[gopherize.me](https://gopherize.me/)* is such a sweet resource!
![goperize.me](gopheravatar.png)    
