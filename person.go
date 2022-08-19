package faker

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker/v4/pkg/options"
)

// Dowser provides interfaces to generate random logical Names with their initials
type Dowser interface {
	TitleMale(v reflect.Value) (interface{}, error)
	TitleFeMale(v reflect.Value) (interface{}, error)
	FirstName(v reflect.Value) (interface{}, error)
	FirstNameMale(v reflect.Value) (interface{}, error)
	FirstNameFemale(v reflect.Value) (interface{}, error)
	LastName(v reflect.Value) (interface{}, error)
	Name(v reflect.Value) (interface{}, error)
	Gender(v reflect.Value) (interface{}, error)
	ChineseFirstName(v reflect.Value) (interface{}, error)
	ChineseLastName(v reflect.Value) (interface{}, error)
	ChineseName(v reflect.Value) (interface{}, error)
}

var titlesMale = []string{
	"Mr.", "Dr.", "Prof.", "Lord", "King", "Prince",
}
var titlesFemale = []string{
	"Mrs.", "Ms.", "Miss", "Dr.", "Prof.", "Lady", "Queen", "Princess",
}
var firstNamesMale = []string{
	"Aaron", "Abdiel", "Abdul", "Abdullah", "Abe", "Abel", "Abelardo", "Abner", "Abraham", "Adalberto", "Adam", "Adan", "Adelbert", "Adolfo", "Adolph", "Adolphus", "Adonis", "Adrain", "Adrian", "Adriel", "Adrien", "Afton", "Agustin", "Ahmad", "Ahmed", "Aidan", "Aiden", "Akeem", "Al", "Alan", "Albert", "Alberto", "Albin", "Alden", "Alec", "Alejandrin", "Alek", "Alessandro", "Alex", "Alexander", "Alexandre", "Alexandro", "Alexie", "Alexis", "Alexys", "Alexzander", "Alf", "Alfonso", "Alfonzo", "Alford", "Alfred", "Alfredo", "Ali", "Allan", "Allen", "Alphonso", "Alvah", "Alvis", "Amani", "Amari", "Ambrose", "Americo", "Amir", "Amos", "Amparo", "Anastacio", "Anderson", "Andre", "Andres", "Andrew", "Andy", "Angel", "Angelo", "Angus", "Anibal", "Ansel", "Ansley", "Anthony", "Antone", "Antonio", "Antwan", "Antwon", "Arch", "Archibald", "Arden", "Arely", "Ari", "Aric", "Ariel", "Arjun", "Arlo", "Armand", "Armando", "Armani", "Arnaldo", "Arne", "Arno", "Arnold", "Arnoldo", "Arnulfo", "Aron", "Art", "Arthur", "Arturo", "Arvel", "Arvid", "Ashton", "August", "Augustus", "Aurelio", "Austen", "Austin", "Austyn", "Avery", "Axel", "Ayden",
	"Bailey", "Barney", "Baron", "Barrett", "Barry", "Bart", "Bartholome", "Barton", "Baylee", "Beau", "Bell", "Ben", "Benedict", "Benjamin", "Bennett", "Bennie", "Benny", "Benton", "Bernard", "Bernardo", "Bernhard", "Bernie", "Berry", "Berta", "Bertha", "Bertram", "Bertrand", "Bill", "Billy", "Blair", "Blaise", "Blake", "Blaze", "Bo", "Bobbie", "Bobby", "Boris", "Boyd", "Brad", "Braden", "Bradford", "Bradley", "Bradly", "Brady", "Braeden", "Brain", "Brando", "Brandon", "Brandt", "Brannon", "Branson", "Brant", "Braulio", "Braxton", "Brayan", "Brendan", "Brenden", "Brendon", "Brennan", "Brennon", "Brent", "Bret", "Brett", "Brian", "Brice", "Brock", "Broderick", "Brody", "Brook", "Brooks", "Brown", "Bruce", "Bryce", "Brycen", "Bryon", "Buck", "Bud", "Buddy", "Buford", "Burley", "Buster",
	"Cade", "Caden", "Caesar", "Cale", "Caleb", "Camden", "Cameron", "Camren", "Camron", "Camryn", "Candelario", "Candido", "Carey", "Carleton", "Carlo", "Carlos", "Carmel", "Carmelo", "Carmine", "Carol", "Carroll", "Carson", "Carter", "Cary", "Casey", "Casimer", "Casimir", "Casper", "Ceasar", "Cecil", "Cedrick", "Celestino", "Cesar", "Chad", "Chadd", "Chadrick", "Chaim", "Chance", "Chandler", "Charles", "Charley", "Charlie", "Chase", "Chauncey", "Chaz", "Chelsey", "Chesley", "Chester", "Chet", "Chris", "Christ", "Christian", "Christop", "Christophe", "Christopher", "Cicero", "Cielo", "Clair", "Clark", "Claud", "Claude", "Clay", "Clemens", "Clement", "Cleo", "Cletus", "Cleve", "Cleveland", "Clifford", "Clifton", "Clint", "Clinton", "Clovis", "Cloyd", "Clyde", "Coby", "Cody", "Colby", "Cole", "Coleman", "Colin", "Collin", "Colt", "Colten", "Colton", "Columbus", "Conner", "Connor", "Conor", "Conrad", "Constantin", "Consuelo", "Cooper", "Corbin", "Cordelia", "Cordell", "Cornelius", "Cornell", "Cortez", "Cory", "Coty", "Coy", "Craig", "Crawford", "Cristian", "Cristina", "Cristobal", "Cristopher", "Cruz", "Cullen", "Curt", "Curtis", "Cyril", "Cyrus",
	"Dagmar", "Dale", "Dallas", "Dallin", "Dalton", "Dameon", "Damian", "Damien", "Damion", "Damon", "Dan", "Dane", "D'angelo", "Dangelo", "Danial", "Danny", "Dante", "Daren", "Darian", "Darien", "Dario", "Darion", "Darius", "Daron", "Darrel", "Darrell", "Darren", "Darrick", "Darrin", "Darrion", "Darron", "Darryl", "Darwin", "Daryl", "Dashawn", "Dave", "David", "Davin", "Davion", "Davon", "Davonte", "Dawson", "Dax", "Dayne", "Dayton", "Dean", "Deangelo", "Declan", "Dedric", "Dedrick", "Dee", "Deion", "Dejon", "Dejuan", "Delaney", "Delbert", "Dell", "Delmer", "Demarco", "Demarcus", "Demario", "Demetrius", "Demond", "Denis", "Dennis", "Deon", "Deondre", "Deontae", "Deonte", "Dereck", "Derek", "Derick", "Deron", "Derrick", "Deshaun", "Deshawn", "Desmond", "Destin", "Devan", "Devante", "Deven", "Devin", "Devon", "Devonte", "Devyn", "Dewayne", "Dewitt", "Dexter", "Diamond", "Diego", "Dillan", "Dillon", "Dimitri", "Dino", "Dion", "Dock", "Domenic", "Domenick", "Domenico", "Domingo", "Dominic", "Don", "Donald", "Donato", "Donavon", "Donnell", "Donnie", "Donny", "Dorcas", "Dorian", "Doris", "Dorthy", "Doug", "Douglas", "Doyle", "Drake", "Dudley", "Duncan", "Durward", "Dustin", "Dusty", "Dwight", "Dylan",
	"Earl", "Earnest", "Easter", "Easton", "Ed", "Edd", "Eddie", "Edgar", "Edgardo", "Edison", "Edmond", "Edmund", "Eduardo", "Edward", "Edwardo", "Edwin", "Efrain", "Efren", "Einar", "Eino", "Eladio", "Elbert", "Eldon", "Eldred", "Eleazar", "Eli", "Elian", "Elias", "Eliezer", "Elijah", "Eliseo", "Elliot", "Elliott", "Ellis", "Ellsworth", "Elmer", "Elmo", "Elmore", "Eloy", "Elroy", "Elton", "Elvis", "Elwin", "Elwyn", "Emanuel", "Emerald", "Emerson", "Emery", "Emil", "Emile", "Emiliano", "Emilio", "Emmanuel", "Emmet", "Emmett", "Emmitt", "Emory", "Enid", "Enoch", "Enos", "Enrico", "Enrique", "Ephraim", "Eriberto", "Eric", "Erich", "Erick", "Erik", "Erin", "Erling", "Ernest", "Ernesto", "Ernie", "Ervin", "Erwin", "Esteban", "Estevan", "Ethan", "Ethel", "Eugene", "Eusebio", "Evan", "Evans", "Everardo", "Everett", "Evert", "Ewald", "Ewell", "Ezekiel", "Ezequiel", "Ezra",
	"Fabian", "Faustino", "Fausto", "Favian", "Federico", "Felipe", "Felix", "Felton", "Fermin", "Fern", "Fernando", "Ferne", "Fidel", "Filiberto", "Finn", "Flavio", "Fletcher", "Florencio", "Florian", "Floy", "Floyd", "Ford", "Forest", "Forrest", "Foster", "Francesco", "Francis", "Francisco", "Franco", "Frank", "Frankie", "Franz", "Fred", "Freddie", "Freddy", "Frederic", "Frederick", "Frederik", "Fredrick", "Fredy", "Freeman", "Friedrich", "Fritz", "Furman",
	"Gabe", "Gabriel", "Gaetano", "Gage", "Gardner", "Garett", "Garfield", "Garland", "Garnet", "Garnett", "Garret", "Garrett", "Garrick", "Garrison", "Garry", "Garth", "Gaston", "Gavin", "Gay", "Gayle", "Gaylord", "Gene", "General", "Gennaro", "Geo", "Geoffrey", "George", "Geovanni", "Geovanny", "Geovany", "Gerald", "Gerard", "Gerardo", "Gerhard", "German", "Gerson", "Gianni", "Gideon", "Gilbert", "Gilberto", "Giles", "Gillian", "Gino", "Giovani", "Giovanni", "Giovanny", "Giuseppe", "Glen", "Glennie", "Godfrey", "Golden", "Gonzalo", "Gordon", "Grady", "Graham", "Grant", "Granville", "Grayce", "Grayson", "Green", "Greg", "Gregg", "Gregorio", "Gregory", "Greyson", "Griffin", "Grover", "Guido", "Guillermo", "Guiseppe", "Gunnar", "Gunner", "Gus", "Gussie", "Gust", "Gustave", "Guy",
	"Hadley", "Hailey", "Hal", "Haleigh", "Haley", "Halle", "Hank", "Hans", "Hardy", "Harley", "Harmon", "Harold", "Harrison", "Harry", "Harvey", "Haskell", "Hassan", "Hayden", "Hayley", "Hazel", "Hazle", "Heber", "Hector", "Helmer", "Henderson", "Henri", "Henry", "Herbert", "Herman", "Hermann", "Herminio", "Hershel", "Hester", "Hilario", "Hilbert", "Hillard", "Hilton", "Hipolito", "Hiram", "Hobart", "Holden", "Hollis", "Horace", "Horacio", "Houston", "Howard", "Howell", "Hoyt", "Hubert", "Hudson", "Hugh", "Humberto", "Hunter", "Hyman",
	"Ian", "Ibrahim", "Ignacio", "Ignatius", "Ike", "Imani", "Immanuel", "Irving", "Irwin", "Isaac", "Isac", "Isadore", "Isai", "Isaiah", "Isaias", "Isidro", "Ismael", "Isom", "Israel", "Issac", "Izaiah",
	"Jabari", "Jace", "Jacey", "Jacinto", "Jack", "Jackson", "Jacques", "Jaden", "Jadon", "Jaeden", "Jaiden", "Jaime", "Jairo", "Jake", "Jakob", "Jaleel", "Jalen", "Jalon", "Jamaal", "Jamal", "Jamar", "Jamarcus", "Jamel", "Jameson", "Jamey", "Jamie", "Jamil", "Jamir", "Jamison", "Jan", "Janick", "Jaquan", "Jared", "Jaren", "Jarod", "Jaron", "Jarred", "Jarrell", "Jarret", "Jarrett", "Jarrod", "Jarvis", "Jasen", "Jasmin", "Jason", "Jasper", "Javier", "Javon", "Javonte", "Jay", "Jayce", "Jaycee", "Jayde", "Jayden", "Jaydon", "Jaylan", "Jaylen", "Jaylin", "Jaylon", "Jayme", "Jayson", "Jean", "Jed", "Jedediah", "Jedidiah", "Jeff", "Jefferey", "Jeffery", "Jeffrey", "Jeffry", "Jennings", "Jensen", "Jerad", "Jerald", "Jeramie", "Jeramy", "Jerel", "Jeremie", "Jeremy", "Jermain", "Jermey", "Jerod", "Jerome", "Jeromy", "Jerrell", "Jerrod", "Jerrold", "Jerry", "Jess", "Jesse", "Jessie", "Jessy", "Jesus", "Jett", "Jettie", "Jevon", "Jillian", "Jimmie", "Jimmy", "Jo", "Joan", "Joany", "Joaquin", "Jocelyn", "Joe", "Joel", "Joesph", "Joey", "Johan", "Johann", "Johathan", "John", "Johnathan", "Johnathon", "Johnnie", "Johnny", "Johnpaul", "Johnson", "Jon", "Jonas", "Jonatan", "Jonathan", "Jonathon", "Jordan", "Jordi", "Jordon", "Jordy", "Jordyn", "Jorge", "Jose", "Joseph", "Josh", "Joshua", "Joshuah", "Josiah", "Josue", "Jovan", "Jovani", "Jovanny", "Jovany", "Judah", "Judd", "Judge", "Judson", "Jules", "Julian", "Julien", "Julio", "Julius", "Junior", "Junius", "Justen", "Justice", "Juston", "Justus", "Justyn", "Juvenal", "Juwan",
	"Kacey", "Kade", "Kaden", "Kadin", "Kale", "Kaleb", "Kaleigh", "Kaley", "Kameron", "Kamren", "Kamron", "Kamryn", "Kane", "Kareem", "Karl", "Karley", "Karson", "Kay", "Kayden", "Kayleigh", "Kayley", "Keagan", "Keanu", "Keaton", "Keegan", "Keeley", "Keenan", "Keith", "Kellen", "Kelley", "Kelton", "Kelvin", "Ken", "Kendall", "Kendrick", "Kennedi", "Kennedy", "Kenneth", "Kennith", "Kenny", "Kenton", "Kenyon", "Keon", "Keshaun", "Keshawn", "Keven", "Kevin", "Kevon", "Keyon", "Keyshawn", "Khalid", "Khalil", "Kian", "Kiel", "Kieran", "Kiley", "Kim", "King", "Kip", "Kirk", "Kobe", "Koby", "Kody", "Kolby", "Kole", "Korbin", "Korey", "Kory", "Kraig", "Kris", "Kristian", "Kristofer", "Kristoffer", "Kristopher", "Kurt", "Kurtis", "Kyle", "Kyleigh", "Kyler",
	"Ladarius", "Lafayette", "Lamar", "Lambert", "Lamont", "Lance", "Landen", "Lane", "Laron", "Larry", "Larue", "Laurel", "Lavern", "Laverna", "Laverne", "Lavon", "Lawrence", "Lawson", "Layne", "Lazaro", "Lee", "Leif", "Leland", "Lemuel", "Lennie", "Lenny", "Leo", "Leon", "Leonard", "Leonardo", "Leone", "Leonel", "Leopold", "Leopoldo", "Lesley", "Lester", "Levi", "Lew", "Lewis", "Lexus", "Liam", "Lincoln", "Lindsey", "Linwood", "Lionel", "Lisandro", "Llewellyn", "Lloyd", "Logan", "Lon", "London", "Lonnie", "Lonny", "Lonzo", "Lorenz", "Lorenza", "Lorenzo", "Louie", "Louisa", "Lourdes", "Louvenia", "Lowell", "Loy", "Loyal", "Lucas", "Luciano", "Lucio", "Lucious", "Lucius", "Ludwig", "Luigi", "Luis", "Lukas", "Lula", "Luther", "Lyric",
	"Mac", "Macey", "Mack", "Mackenzie", "Madisen", "Madison", "Madyson", "Magnus", "Major", "Makenna", "Malachi", "Malcolm", "Mallory", "Manley", "Manuel", "Manuela", "Marc", "Marcel", "Marcelino", "Marcellus", "Marcelo", "Marco", "Marcos", "Marcus", "Mariano", "Mario", "Mark", "Markus", "Marley", "Marlin", "Marlon", "Marques", "Marquis", "Marshall", "Martin", "Marty", "Marvin", "Mason", "Mateo", "Mathew", "Mathias", "Matt", "Matteo", "Maurice", "Mauricio", "Maverick", "Mavis", "Max", "Maxime", "Maximilian", "Maximillian", "Maximo", "Maximus", "Maxine", "Maxwell", "Maynard", "Mckenna", "Mckenzie", "Mekhi", "Melany", "Melvin", "Melvina", "Merl", "Merle", "Merlin", "Merritt", "Mervin", "Micah", "Michael", "Michale", "Micheal", "Michel", "Miguel", "Mike", "Mikel", "Milan", "Miles", "Milford", "Miller", "Milo", "Milton", "Misael", "Mitchel", "Mitchell", "Modesto", "Mohamed", "Mohammad", "Mohammed", "Moises", "Monroe", "Monserrat", "Monserrate", "Montana", "Monte", "Monty", "Morgan", "Moriah", "Morris", "Mortimer", "Morton", "Mose", "Moses", "Moshe", "Muhammad", "Murl", "Murphy", "Murray", "Mustafa", "Myles", "Myrl", "Myron",
	"Napoleon", "Narciso", "Nash", "Nasir", "Nat", "Nathan", "Nathanael", "Nathanial", "Nathaniel", "Nathen", "Neal", "Ned", "Neil", "Nels", "Nelson", "Nestor", "Newell", "Newton", "Nicholas", "Nicholaus", "Nick", "Nicklaus", "Nickolas", "Nico", "Nicola", "Nicolas", "Nigel", "Nikko", "Niko", "Nikolas", "Nils", "Noah", "Noble", "Noe", "Noel", "Nolan", "Norbert", "Norberto", "Norris", "Norval", "Norwood",
	"Obie", "Oda", "Odell", "Okey", "Ola", "Olaf", "Ole", "Olen", "Olin", "Oliver", "Omari", "Omer", "Oral", "Oran", "Oren", "Orin", "Orion", "Orland", "Orlando", "Orlo", "Orrin", "Orval", "Orville", "Osbaldo", "Osborne", "Oscar", "Osvaldo", "Oswald", "Oswaldo", "Otho", "Otis", "Ottis", "Otto", "Owen",
	"Pablo", "Paolo", "Paris", "Parker", "Patrick", "Paul", "Paxton", "Payton", "Pedro", "Percival", "Percy", "Perry", "Pete", "Peter", "Peyton", "Philip", "Pierce", "Pierre", "Pietro", "Porter", "Presley", "Preston", "Price", "Prince",
	"Quentin", "Quincy", "Quinn", "Quinten", "Quinton",
	"Rafael", "Raheem", "Rahul", "Raleigh", "Ralph", "Ramiro", "Ramon", "Randal", "Randall", "Randi", "Randy", "Ransom", "Raoul", "Raphael", "Rashad", "Rashawn", "Rasheed", "Raul", "Raven", "Ray", "Raymond", "Raymundo", "Reagan", "Reece", "Reed", "Reese", "Regan", "Reggie", "Reginald", "Reid", "Reilly", "Reinhold", "Remington", "Rene", "Reuben", "Rex", "Rey", "Reyes", "Reymundo", "Reynold", "Rhett", "Rhiannon", "Ricardo", "Richard", "Richie", "Richmond", "Rick", "Rickey", "Rickie", "Ricky", "Rico", "Rigoberto", "Riley", "Robb", "Robbie", "Robert", "Roberto", "Robin", "Rocio", "Rocky", "Rod", "Roderick", "Rodger", "Rodolfo", "Rodrick", "Rodrigo", "Roel", "Rogelio", "Roger", "Rogers", "Rolando", "Rollin", "Roman", "Ron", "Ronaldo", "Ronny", "Roosevelt", "Rory", "Rosario", "Roscoe", "Rosendo", "Ross", "Rowan", "Rowland", "Roy", "Royal", "Royce", "Ruben", "Rudolph", "Rudy", "Rupert", "Russ", "Russel", "Russell", "Rusty", "Ryan", "Ryann", "Ryder", "Rylan", "Ryleigh", "Ryley",
	"Sage", "Saige", "Salvador", "Salvatore", "Sam", "Samir", "Sammie", "Sammy", "Samson", "Sanford", "Santa", "Santiago", "Santino", "Santos", "Saul", "Savion", "Schuyler", "Scot", "Scottie", "Scotty", "Seamus", "Sean", "Sebastian", "Sedrick", "Selmer", "Seth", "Shad", "Shane", "Shaun", "Shawn", "Shayne", "Sheldon", "Sheridan", "Sherman", "Sherwood", "Sid", "Sidney", "Sigmund", "Sigrid", "Sigurd", "Silas", "Sim", "Simeon", "Skye", "Skylar", "Sofia", "Soledad", "Solon", "Sonny", "Spencer", "Stan", "Stanford", "Stanley", "Stanton", "Stefan", "Stephan", "Stephen", "Stephon", "Sterling", "Steve", "Stevie", "Stewart", "Stone", "Stuart", "Sven", "Sydney", "Sylvan", "Sylvester",
	"Tad", "Talon", "Tanner", "Tate", "Tatum", "Taurean", "Tavares", "Taylor", "Ted", "Terence", "Terrance", "Terrell", "Terrence", "Terrill", "Terry", "Tevin", "Thad", "Thaddeus", "Theo", "Theodore", "Theron", "Thomas", "Thurman", "Tillman", "Timmothy", "Timmy", "Timothy", "Tito", "Titus", "Tobin", "Toby", "Tod", "Tom", "Tomas", "Tommie", "Toney", "Toni", "Tony", "Torey", "Torrance", "Torrey", "Toy", "Trace", "Tracey", "Travis", "Travon", "Tre", "Tremaine", "Tremayne", "Trent", "Trenton", "Trever", "Trevion", "Trevor", "Trey", "Tristian", "Tristin", "Triston", "Troy", "Trystan", "Turner", "Tyler", "Tyree", "Tyreek", "Tyrel", "Tyrell", "Tyrese", "Tyrique", "Tyshawn", "Tyson",
	"Ubaldo", "Ulices", "Ulises", "Unique", "Urban", "Uriah", "Uriel",
	"Valentin", "Van", "Vance", "Vaughn", "Vern", "Verner", "Vernon", "Vicente", "Victor", "Vidal", "Vince", "Vincent", "Vincenzo", "Vinnie", "Virgil", "Vito", "Vladimir",
	"Wade", "Waino", "Waldo", "Walker", "Wallace", "Walter", "Walton", "Ward", "Warren", "Watson", "Waylon", "Wayne", "Webster", "Weldon", "Wellington", "Wendell", "Werner", "Westley", "Weston", "Wilber", "Wilbert", "Wilburn", "Wiley", "Wilford", "Wilfred", "Wilfredo", "Wilfrid", "Wilhelm", "Will", "Willard", "William", "Willis", "Willy", "Wilmer", "Wilson", "Wilton", "Winfield", "Winston", "Woodrow", "Wyatt", "Wyman",
	"Xavier", "Xzavier", "Xander",
	"Yadav", "Yogesh", "Yaatiesh", "Yaamir",
	"Zachariah", "Zachary", "Zachery", "Zack", "Zackary", "Zackery", "Zakary", "Zander", "Zane", "Zechariah", "Zion",
}
var firstNamesFemale = []string{
	"Aaliyah", "Abagail", "Abbey", "Abbie", "Abbigail", "Abby", "Abigail", "Abigale", "Abigayle", "Ada", "Adah", "Adaline", "Addie", "Addison", "Adela", "Adele", "Adelia", "Adeline", "Adell", "Adella", "Adelle", "Aditya", "Adriana", "Adrianna", "Adrienne", "Aglae", "Agnes", "Agustina", "Aida", "Aileen", "Aimee", "Aisha", "Aiyana", "Alaina", "Alana", "Alanis", "Alanna", "Alayna", "Alba", "Alberta", "Albertha", "Albina", "Alda", "Aleen", "Alejandra", "Alena", "Alene", "Alessandra", "Alessia", "Aletha", "Alexa", "Alexandra", "Alexandrea", "Alexandria", "Alexandrine", "Alexane", "Alexanne", "Alfreda", "Alia", "Alice", "Alicia", "Alisa", "Alisha", "Alison", "Alivia", "Aliya", "Aliyah", "Aliza", "Alize", "Allene", "Allie", "Allison", "Ally", "Alta", "Althea", "Alva", "Alvena", "Alvera", "Alverta", "Alvina", "Alyce", "Alycia", "Alysa", "Alysha", "Alyson", "Alysson", "Amalia", "Amanda", "Amara", "Amaya", "Amber", "Amelia", "Amelie", "Amely", "America", "Amie", "Amina", "Amira", "Amiya", "Amy", "Amya", "Ana", "Anabel", "Anabelle", "Anahi", "Anais", "Anastasia", "Andreane", "Andreanne", "Angela", "Angelica", "Angelina", "Angeline", "Angelita", "Angie", "Anika", "Anissa", "Anita", "Aniya", "Aniyah", "Anjali", "Anna", "Annabel", "Annabell", "Annabelle", "Annalise", "Annamae", "Annamarie", "Anne", "Annetta", "Annette", "Annie", "Antoinette", "Antonetta", "Antonette", "Antonia", "Antonietta", "Antonina", "Anya", "April", "Ara", "Araceli", "Aracely", "Ardella", "Ardith", "Ariane", "Arianna", "Arielle", "Arlene", "Arlie", "Arvilla", "Aryanna", "Asa", "Asha", "Ashlee", "Ashleigh", "Ashley", "Ashly", "Ashlynn", "Ashtyn", "Asia", "Assunta", "Astrid", "Athena", "Aubree", "Aubrey", "Audie", "Audra", "Audreanne", "Audrey", "Augusta", "Augustine", "Aurelia", "Aurelie", "Aurore", "Autumn", "Ava", "Avis", "Ayana", "Ayla", "Aylin",
	"Baby", "Bailee", "Barbara", "Beatrice", "Beaulah", "Bella", "Belle", "Berenice", "Bernadette", "Bernadine", "Berneice", "Bernice", "Berniece", "Bernita", "Bert", "Beryl", "Bessie", "Beth", "Bethany", "Bethel", "Betsy", "Bette", "Bettie", "Betty", "Bettye", "Beulah", "Beverly", "Bianka", "Billie", "Birdie", "Blanca", "Blanche", "Bonita", "Bonnie", "Brandi", "Brandy", "Brandyn", "Breana", "Breanna", "Breanne", "Brenda", "Brenna", "Bria", "Briana", "Brianne", "Bridget", "Bridgette", "Bridie", "Brielle", "Brigitte", "Brionna", "Brisa", "Britney", "Brittany", "Brooke", "Brooklyn", "Bryana", "Bulah", "Burdette", "Burnice",
	"Caitlyn", "Caleigh", "Cali", "Calista", "Callie", "Camila", "Camilla", "Camille", "Camylle", "Candace", "Candice", "Candida", "Cara", "Carissa", "Carlee", "Carley", "Carli", "Carlie", "Carlotta", "Carmela", "Carmella", "Carmen", "Carolanne", "Carole", "Carolina", "Caroline", "Carolyn", "Carolyne", "Carrie", "Casandra", "Cassandra", "Cassandre", "Cassidy", "Cassie", "Catalina", "Caterina", "Catharine", "Catherine", "Cathrine", "Cathryn", "Cathy", "Cayla", "Cecelia", "Cecile", "Cecilia", "Celestine", "Celia", "Celine", "Chanel", "Chanelle", "Charity", "Charlene", "Charlotte", "Chasity", "Chaya", "Chelsea", "Chelsie", "Cheyanne", "Cheyenne", "Chloe", "Christa", "Christelle", "Christiana", "Christina", "Christine", "Christy", "Chyna", "Ciara", "Cierra", "Cindy", "Citlalli", "Claire", "Clara", "Clarabelle", "Clare", "Clarissa", "Claudia", "Claudie", "Claudine", "Clementina", "Clementine", "Clemmie", "Cleora", "Cleta", "Clotilde", "Colleen", "Concepcion", "Connie", "Constance", "Cora", "Coralie", "Cordia", "Cordie", "Corene", "Corine", "Corrine", "Cortney", "Courtney", "Creola", "Cristal", "Crystal", "Crystel", "Cydney", "Cynthia",
	"Dahlia", "Daija", "Daisha", "Daisy", "Dakota", "Damaris", "Dana", "Dandre", "Daniela", "Daniella", "Danielle", "Danika", "Dannie", "Danyka", "Daphne", "Daphnee", "Daphney", "Darby", "Dariana", "Darlene", "Dasia", "Dawn", "Dayana", "Dayna", "Deanna", "Deborah", "Deja", "Dejah", "Delfina", "Delia", "Delilah", "Della", "Delores", "Delpha", "Delphia", "Delphine", "Delta", "Demetris", "Dena", "Desiree", "Dessie", "Destany", "Destinee", "Destiney", "Destini", "Destiny", "Diana", "Dianna", "Dina", "Dixie", "Dolly", "Dolores", "Domenica", "Dominique", "Donna", "Dora", "Dorothea", "Dorothy", "Dorris", "Dortha", "Dovie", "Drew", "Duane", "Dulce",
	"Earlene", "Earline", "Earnestine", "Ebba", "Ebony", "Eda", "Eden", "Edna", "Edwina", "Edyth", "Edythe", "Effie", "Eileen", "Elaina", "Elda", "Eldora", "Eldridge", "Eleanora", "Eleanore", "Electa", "Elena", "Elenor", "Elenora", "Eleonore", "Elfrieda", "Eliane", "Elinor", "Elinore", "Elisa", "Elisabeth", "Elise", "Elisha", "Elissa", "Eliza", "Elizabeth", "Ella", "Ellen", "Ellie", "Elmira", "Elna", "Elnora", "Elody", "Eloisa", "Eloise", "Elouise", "Elsa", "Else", "Elsie", "Elta", "Elva", "Elvera", "Elvie", "Elyse", "Elyssa", "Elza", "Emelia", "Emelie", "Emely", "Emie", "Emilia", "Emilie", "Emily", "Emma", "Emmalee", "Emmanuelle", "Emmie", "Emmy", "Ena", "Enola", "Era", "Erica", "Ericka", "Erika", "Erna", "Ernestina", "Ernestine", "Eryn", "Esmeralda", "Esperanza", "Esta", "Estefania", "Estel", "Estell", "Estella", "Estelle", "Esther", "Estrella", "Etha", "Ethelyn", "Ethyl", "Ettie", "Eudora", "Eugenia", "Eula", "Eulah", "Eulalia", "Euna", "Eunice", "Eva", "Evalyn", "Evangeline", "Eve", "Eveline", "Evelyn", "Everette", "Evie",
	"Fabiola", "Fae", "Fannie", "Fanny", "Fatima", "Fay", "Faye", "Felicia", "Felicita", "Felicity", "Felipa", "Filomena", "Fiona", "Flavie", "Fleta", "Flo", "Florence", "Florida", "Florine", "Flossie", "Frances", "Francesca", "Francisca", "Freda", "Frederique", "Freeda", "Freida", "Frida", "Frieda",
	"Gabriella", "Gabrielle", "Gail", "Genesis", "Genevieve", "Genoveva", "Georgette", "Georgiana", "Georgianna", "Geraldine", "Gerda", "Germaine", "Gerry", "Gertrude", "Gia", "Gilda", "Gina", "Giovanna", "Gisselle", "Gladyce", "Gladys", "Glenda", "Glenna", "Gloria", "Golda", "Grace", "Gracie", "Graciela", "Gregoria", "Greta", "Gretchen", "Guadalupe", "Gudrun", "Gwen", "Gwendolyn",
	"Hailee", "Hailie", "Halie", "Hallie", "Hanna", "Hannah", "Harmony", "Hassie", "Hattie", "Haven", "Haylee", "Haylie", "Heath", "Heather", "Heaven", "Heidi", "Helen", "Helena", "Helene", "Helga", "Hellen", "Heloise", "Henriette", "Hermina", "Herminia", "Herta", "Hertha", "Hettie", "Hilda", "Hildegard", "Hillary", "Hilma", "Hollie", "Holly", "Hope", "Hortense", "Hosea", "Hulda",
	"Icie", "Ida", "Idell", "Idella", "Ila", "Ilene", "Iliana", "Ima", "Imelda", "Imogene", "Ines", "Irma", "Isabel", "Isabell", "Isabella", "Isabelle", "Isobel", "Itzel", "Iva", "Ivah", "Ivory", "Ivy", "Izabella",
	"Jacinthe", "Jackeline", "Jackie", "Jacklyn", "Jacky", "Jaclyn", "Jacquelyn", "Jacynthe", "Jada", "Jade", "Jadyn", "Jaida", "Jailyn", "Jakayla", "Jalyn", "Jammie", "Jana", "Janae", "Jane", "Janelle", "Janessa", "Janet", "Janice", "Janie", "Janis", "Janiya", "Jannie", "Jany", "Jaquelin", "Jaqueline", "Jaunita", "Jayda", "Jayne", "Jazlyn", "Jazmin", "Jazmyn", "Jazmyne", "Jeanette", "Jeanie", "Jeanne", "Jena", "Jenifer", "Jennie", "Jennifer", "Jennyfer", "Jermaine", "Jessica", "Jessika", "Jessyca", "Jewel", "Jewell", "Joana", "Joanie", "Joanne", "Joannie", "Joanny", "Jodie", "Jody", "Joelle", "Johanna", "Jolie", "Jordane", "Josefa", "Josefina", "Josephine", "Josiane", "Josianne", "Josie", "Joy", "Joyce", "Juana", "Juanita", "Jude", "Judy", "Julia", "Juliana", "Julianne", "Julie", "Juliet", "June", "Justina", "Justine",
	"Kaci", "Kacie", "Kaela", "Kaelyn", "Kaia", "Kailee", "Kailey", "Kailyn", "Kaitlin", "Kaitlyn", "Kali", "Kallie", "Kamille", "Kara", "Karelle", "Karen", "Kari", "Kariane", "Karianne", "Karina", "Karine", "Karlee", "Karli", "Karlie", "Karolann", "Kasandra", "Kasey", "Kassandra", "Katarina", "Katelin", "Katelyn", "Katelynn", "Katharina", "Katherine", "Katheryn", "Kathleen", "Kathlyn", "Kathryn", "Kathryne", "Katlyn", "Katlynn", "Katrina", "Katrine", "Kattie", "Kavon", "Kaya", "Kaycee", "Kayla", "Kaylah", "Kaylee", "Kayli", "Kaylie", "Kaylin", "Keara", "Keely", "Keira", "Kelli", "Kellie", "Kelly", "Kelsi", "Kelsie", "Kendra", "Kenna", "Kenya", "Kenyatta", "Kiana", "Kianna", "Kiara", "Kiarra", "Kiera", "Kimberly", "Kira", "Kirsten", "Kirstin", "Kitty", "Krista", "Kristin", "Kristina", "Kristy", "Krystal", "Krystel", "Krystina", "Kyla", "Kylee", "Kylie", "Kyra",
	"Lacey", "Lacy", "Laila", "Laisha", "Laney", "Larissa", "Laura", "Lauren", "Laurence", "Lauretta", "Lauriane", "Laurianne", "Laurie", "Laurine", "Laury", "Lauryn", "Lavada", "Lavina", "Lavinia", "Lavonne", "Layla", "Lea", "Leann", "Leanna", "Leanne", "Leatha", "Leda", "Leila", "Leilani", "Lela", "Lelah", "Lelia", "Lempi", "Lenna", "Lenora", "Lenore", "Leola", "Leonie", "Leonor", "Leonora", "Leora", "Lera", "Leslie", "Lesly", "Lessie", "Leta", "Letha", "Letitia", "Lexi", "Lexie", "Lia", "Liana", "Libbie", "Libby", "Lila", "Lilian", "Liliana", "Liliane", "Lilla", "Lillian", "Lilliana", "Lillie", "Lilly", "Lily", "Lilyan", "Lina", "Linda", "Lindsay", "Linnea", "Linnie", "Lisa", "Lisette", "Litzy", "Liza", "Lizeth", "Lizzie", "Lois", "Lola", "Lolita", "Loma", "Lonie", "Lora", "Loraine", "Loren", "Lorena", "Lori", "Lorine", "Lorna", "Lottie", "Lou", "Loyce", "Lucie", "Lucienne", "Lucile", "Lucinda", "Lucy", "Ludie", "Lue", "Luella", "Luisa", "Lulu", "Luna", "Lupe", "Lura", "Lurline", "Luz", "Lyda", "Lydia", "Lyla", "Lynn", "Lysanne",
	"Mabel", "Mabelle", "Mable", "Maci", "Macie", "Macy", "Madaline", "Madalyn", "Maddison", "Madeline", "Madelyn", "Madelynn", "Madge", "Madie", "Madilyn", "Madisyn", "Madonna", "Mae", "Maegan", "Maeve", "Mafalda", "Magali", "Magdalen", "Magdalena", "Maggie", "Magnolia", "Maia", "Maida", "Maiya", "Makayla", "Makenzie", "Malika", "Malinda", "Mallie", "Malvina", "Mandy", "Mara", "Marcelina", "Marcella", "Marcelle", "Marcia", "Margaret", "Margarete", "Margarett", "Margaretta", "Margarette", "Margarita", "Marge", "Margie", "Margot", "Margret", "Marguerite", "Maria", "Mariah", "Mariam", "Marian", "Mariana", "Mariane", "Marianna", "Marianne", "Maribel", "Marie", "Mariela", "Marielle", "Marietta", "Marilie", "Marilou", "Marilyne", "Marina", "Marion", "Marisa", "Marisol", "Maritza", "Marjolaine", "Marjorie", "Marjory", "Marlee", "Marlen", "Marlene", "Marquise", "Marta", "Martina", "Martine", "Mary", "Maryam", "Maryjane", "Maryse", "Mathilde", "Matilda", "Matilde", "Mattie", "Maud", "Maude", "Maudie", "Maureen", "Maurine", "Maxie", "Maximillia", "May", "Maya", "Maybell", "Maybelle", "Maye", "Maymie", "Mayra", "Mazie", "Mckayla", "Meagan", "Meaghan", "Meda", "Megane", "Meggie", "Meghan", "Melba", "Melisa", "Melissa", "Mellie", "Melody", "Melyna", "Melyssa", "Mercedes", "Meredith", "Mertie", "Meta", "Mia", "Micaela", "Michaela", "Michele", "Michelle", "Mikayla", "Millie", "Mina", "Minerva", "Minnie", "Miracle", "Mireille", "Mireya", "Missouri", "Misty", "Mittie", "Modesta", "Mollie", "Molly", "Mona", "Monica", "Monique", "Mossie", "Mozell", "Mozelle", "Muriel", "Mya", "Myah", "Mylene", "Myra", "Myriam", "Myrna", "Myrtice", "Myrtie", "Myrtis", "Myrtle",
	"Nadia", "Nakia", "Name", "Nannie", "Naomi", "Naomie", "Natalia", "Natalie", "Natasha", "Nayeli", "Nedra", "Neha", "Nelda", "Nella", "Nelle", "Nellie", "Neoma", "Nettie", "Neva", "Nia", "Nichole", "Nicole", "Nicolette", "Nikita", "Nikki", "Nina", "Noelia", "Noemi", "Noemie", "Noemy", "Nola", "Nona", "Nora", "Norene", "Norma", "Nova", "Novella", "Nya", "Nyah", "Nyasia",
	"Oceane", "Ocie", "Octavia", "Odessa", "Odie", "Ofelia", "Oleta", "Olga", "Ollie", "Oma", "Ona", "Onie", "Opal", "Ophelia", "Ora", "Orie", "Orpha", "Otha", "Otilia", "Ottilie", "Ova", "Ozella",
	"Paige", "Palma", "Pamela", "Pansy", "Pascale", "Pasquale", "Pat", "Patience", "Patricia", "Patsy", "Pattie", "Paula", "Pauline", "Pearl", "Pearlie", "Pearline", "Peggie", "Penelope", "Petra", "Phoebe", "Phyllis", "Pink", "Pinkie", "Piper", "Polly", "Precious", "Princess", "Priscilla", "Providenci", "Prudence",
	"Queen", "Queenie",
	"Rachael", "Rachel", "Rachelle", "Rae", "Raegan", "Rafaela", "Rahsaan", "Raina", "Ramona", "Raphaelle", "Raquel", "Reanna", "Reba", "Rebeca", "Rebecca", "Rebeka", "Rebekah", "Reina", "Renee", "Ressie", "Reta", "Retha", "Retta", "Reva", "Reyna", "Rhea", "Rhianna", "Rhoda", "Rita", "River", "Roberta", "Robyn", "Roma", "Romaine", "Rosa", "Rosalee", "Rosalia", "Rosalind", "Rosalinda", "Rosalyn", "Rosamond", "Rosanna", "Rose", "Rosella", "Roselyn", "Rosemarie", "Rosemary", "Rosetta", "Rosie", "Rosina", "Roslyn", "Rossie", "Rowena", "Roxane", "Roxanne", "Rozella", "Rubie", "Ruby", "Rubye", "Ruth", "Ruthe", "Ruthie", "Rylee",
	"Sabina", "Sabrina", "Sabryna", "Sadie", "Sadye", "Sallie", "Sally", "Salma", "Samanta", "Samantha", "Samara", "Sandra", "Sandrine", "Sandy", "Santina", "Sarah", "Sarai", "Sarina", "Sasha", "Savanah", "Savanna", "Savannah", "Scarlett", "Selena", "Selina", "Serena", "Serenity", "Shaina", "Shakira", "Shana", "Shanel", "Shanelle", "Shania", "Shanie", "Shaniya", "Shanna", "Shannon", "Shanny", "Shanon", "Shany", "Sharon", "Shawna", "Shaylee", "Shayna", "Shea", "Sheila", "Shemar", "Shirley", "Shyann", "Shyanne", "Sibyl", "Sienna", "Sierra", "Simone", "Sincere", "Sister", "Skyla", "Sonia", "Sonya", "Sophia", "Sophie", "Stacey", "Stacy", "Stefanie", "Stella", "Stephania", "Stephanie", "Stephany", "Summer", "Sunny", "Susan", "Susana", "Susanna", "Susie", "Suzanne", "Syble", "Sydnee", "Sydni", "Sydnie", "Sylvia",
	"Tabitha", "Talia", "Tamara", "Tamia", "Tania", "Tanya", "Tara", "Taryn", "Tatyana", "Taya", "Teagan", "Telly", "Teresa", "Tess", "Tessie", "Thalia", "Thea", "Thelma", "Theodora", "Theresa", "Therese", "Theresia", "Thora", "Tia", "Tiana", "Tianna", "Tiara", "Tierra", "Tiffany", "Tina", "Tomasa", "Tracy", "Tressa", "Tressie", "Treva", "Trinity", "Trisha", "Trudie", "Trycia", "Twila", "Tyra",
	"Una", "Ursula",
	"Vada", "Valentina", "Valentine", "Valerie", "Vallie", "Vanessa", "Veda", "Velda", "Vella", "Velma", "Velva", "Vena", "Verda", "Verdie", "Vergie", "Verla", "Verlie", "Verna", "Vernice", "Vernie", "Verona", "Veronica", "Vesta", "Vicenta", "Vickie", "Vicky", "Victoria", "Vida", "Vilma", "Vincenza", "Viola", "Violet", "Violette", "Virgie", "Virginia", "Virginie", "Vita", "Viva", "Vivian", "Viviane", "Vivianne", "Vivien", "Vivienne",
	"Wanda", "Wava", "Wendy", "Whitney", "Wilhelmine", "Willa", "Willie", "Willow", "Wilma", "Winifred", "Winnifred", "Winona",
	"Xiu-Mei", "Xio",
	"Yadira", "Yasmeen", "Yasmin", "Yasmine", "Yazmin", "Yesenia", "Yessenia", "Yolanda", "Yoshiko", "Yvette", "Yvonne",
	"Zaria", "Zelda", "Zella", "Zelma", "Zena", "Zetta", "Zita", "Zoe", "Zoey", "Zoie", "Zoila", "Zola", "Zora", "Zula",
}

var firstNames = append(firstNamesMale, firstNamesFemale...)

var lastNames = []string{
	"Abbott", "Abernathy", "Abshire", "Adams", "Altenwerth", "Anderson", "Ankunding", "Armstrong", "Auer", "Aufderhar",
	"Bahringer", "Bailey", "Balistreri", "Barrows", "Bartell", "Bartoletti", "Barton", "Bashirian", "Batz", "Bauch", "Baumbach", "Bayer", "Beahan", "Beatty", "Bechtelar", "Becker", "Bednar", "Beer", "Beier", "Berge", "Bergnaum", "Bergstrom", "Bernhard", "Bernier", "Bins", "Blanda", "Blick", "Block", "Bode", "Boehm", "Bogan", "Bogisich", "Borer", "Bosco", "Botsford", "Boyer", "Boyle", "Bradtke", "Brakus", "Braun", "Breitenberg", "Brekke", "Brown", "Bruen", "Buckridge",
	"Carroll", "Carter", "Cartwright", "Casper", "Cassin", "Champlin", "Christiansen", "Cole", "Collier", "Collins", "Conn", "Connelly", "Conroy", "Considine", "Corkery", "Cormier", "Corwin", "Cremin", "Crist", "Crona", "Cronin", "Crooks", "Cruickshank", "Cummerata", "Cummings",
	"Dach", "D'Amore", "Daniel", "Dare", "Daugherty", "Davis", "Deckow", "Denesik", "Dibbert", "Dickens", "Dicki", "Dickinson", "Dietrich", "Donnelly", "Dooley", "Douglas", "Doyle", "DuBuque", "Durgan",
	"Ebert", "Effertz", "Eichmann", "Emard", "Emmerich", "Erdman", "Ernser", "Fadel",
	"Fahey", "Farrell", "Fay", "Feeney", "Feest", "Feil", "Ferry", "Fisher", "Flatley", "Frami", "Franecki", "Friesen", "Fritsch", "Funk",
	"Gaylord", "Gerhold", "Gerlach", "Gibson", "Gislason", "Gleason", "Gleichner", "Glover", "Goldner", "Goodwin", "Gorczany", "Gottlieb", "Goyette", "Grady", "Graham", "Grant", "Green", "Greenfelder", "Greenholt", "Grimes", "Gulgowski", "Gusikowski", "Gutkowski", "Gutmann",
	"Haag", "Hackett", "Hagenes", "Hahn", "Haley", "Halvorson", "Hamill", "Hammes", "Hand", "Hane", "Hansen", "Harber", "Harris", "Hartmann", "Harvey", "Hauck", "Hayes", "Heaney", "Heathcote", "Hegmann", "Heidenreich", "Heller", "Herman", "Hermann", "Hermiston", "Herzog", "Hessel", "Hettinger", "Hickle", "Hilll", "Hills", "Hilpert", "Hintz", "Hirthe", "Hodkiewicz", "Hoeger", "Homenick", "Hoppe", "Howe", "Howell", "Hudson", "Huel", "Huels", "Hyatt",
	"Jacobi", "Jacobs", "Jacobson", "Jakubowski", "Jaskolski", "Jast", "Jenkins", "Jerde", "Johns", "Johnson", "Johnston", "Jones",
	"Kassulke", "Kautzer", "Keebler", "Keeling", "Kemmer", "Kerluke", "Kertzmann", "Kessler", "Kiehn", "Kihn", "Kilback", "King", "Kirlin", "Klein", "Kling", "Klocko", "Koch", "Koelpin", "Koepp", "Kohler", "Konopelski", "Koss", "Kovacek", "Kozey", "Krajcik", "Kreiger", "Kris", "Kshlerin", "Kub", "Kuhic", "Kuhlman", "Kuhn", "Kulas", "Kunde", "Kunze", "Kuphal", "Kutch", "Kuvalis",
	"Labadie", "Lakin", "Lang", "Langosh", "Langworth", "Larkin", "Larson", "Leannon", "Lebsack", "Ledner", "Leffler", "Legros", "Lehner", "Lemke", "Lesch", "Leuschke", "Lind", "Lindgren", "Littel", "Little", "Lockman", "Lowe", "Lubowitz", "Lueilwitz", "Luettgen", "Lynch",
	"Macejkovic", "Maggio", "Mann", "Mante", "Marks", "Marquardt", "Marvin", "Mayer", "Mayert", "McClure", "McCullough", "McDermott", "McGlynn", "McKenzie", "McLaughlin", "Medhurst", "Mertz", "Metz", "Miller", "Mills", "Mitchell", "Moen", "Mohr", "Monahan", "Moore", "Morar", "Morissette", "Mosciski", "Mraz", "Mueller", "Muller", "Murazik", "Murphy", "Murray",
	"Nader", "Nicolas", "Nienow", "Nikolaus", "Nitzsche", "Nolan",
	"Oberbrunner", "O'Connell", "O'Conner", "O'Hara", "O'Keefe", "O'Kon", "Okuneva", "Olson", "Ondricka", "O'Reilly", "Orn", "Ortiz", "Osinski",
	"Pacocha", "Padberg", "Pagac", "Parisian", "Parker", "Paucek", "Pfannerstill", "Pfeffer", "Pollich", "Pouros", "Powlowski", "Predovic", "Price", "Prohaska", "Prosacco", "Purdy",
	"Quigley", "Quitzon",
	"Rath", "Ratke", "Rau", "Raynor", "Reichel", "Reichert", "Reilly", "Reinger", "Rempel", "Renner", "Reynolds", "Rice", "Rippin", "Ritchie", "Robel", "Roberts", "Rodriguez", "Rogahn", "Rohan", "Rolfson", "Romaguera", "Roob", "Rosenbaum", "Rowe", "Ruecker", "Runolfsdottir", "Runolfsson", "Runte", "Russel", "Rutherford", "Ryan", "Sanford", "Satterfield", "Sauer", "Sawayn",
	"Schaden", "Schaefer", "Schamberger", "Schiller", "Schimmel", "Schinner", "Schmeler", "Schmidt", "Schmitt", "Schneider", "Schoen", "Schowalter", "Schroeder", "Schulist", "Schultz", "Schumm", "Schuppe", "Schuster", "Senger", "Shanahan", "Shields", "Simonis", "Sipes", "Skiles", "Smith", "Smitham", "Spencer", "Spinka", "Sporer", "Stamm", "Stanton", "Stark", "Stehr", "Steuber", "Stiedemann", "Stokes", "Stoltenberg", "Stracke", "Streich", "Stroman", "Strosin", "Swaniawski", "Swift",
	"Terry", "Thiel", "Thompson", "Tillman", "Torp", "Torphy", "Towne", "Toy", "Trantow", "Tremblay", "Treutel", "Tromp", "Turcotte", "Turner",
	"Ullrich", "Upton", "Vandervort", "Veum", "Volkman", "Von", "VonRueden", "Waelchi", "Walker", "Walsh", "Walter", "Ward", "Waters", "Watsica", "Weber", "Wehner", "Weimann", "Weissnat", "Welch", "West", "White", "Wiegand", "Wilderman", "Wilkinson", "Will", "Williamson", "Willms", "Windler", "Wintheiser", "Wisoky", "Wisozk", "Witting", "Wiza", "Wolf", "Wolff", "Wuckert", "Wunsch", "Wyman",
	"Yost", "Yundt", "Zboncak", "Zemlak", "Ziemann", "Zieme", "Zulauf",
}

var randNameFlag int

var genders = []string{"Male", "Female", "Prefer to skip"}

var chineseFirstNames = []string{
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许",
	"何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章",
	"云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳",
	"酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "邬", "安", "常",
	"乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹",
	"姚", "邵", "湛", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成", "戴", "谈", "宋", "茅", "庞",
	"熊", "纪", "舒", "屈", "项", "祝", "董", "梁", "杜", "阮", "蓝", "闵", "席", "季", "麻", "强", "贾", "路", "娄", "危",
	"江", "童", "颜", "郭", "梅", "盛", "林", "刁", "钟", "徐", "邱", "骆", "高", "夏", "蔡", "田", "樊", "胡", "凌", "霍",
	"虞", "万", "支", "柯", "昝", "管", "卢", "莫", "经", "房", "裘", "缪", "干", "解", "应", "宗", "丁", "宣", "贲", "邓",
	"郁", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "钮", "龚", "程", "嵇", "邢", "滑", "裴", "陆", "荣", "翁",
	"荀", "羊", "於", "惠", "甄", "麴", "家", "封", "芮", "羿", "储", "靳", "汲", "邴", "糜", "松", "井", "段", "富", "巫",
	"乌", "焦", "巴", "弓", "牧", "隗", "山", "谷", "车", "侯", "宓", "蓬", "全", "郗", "班", "仰", "秋", "仲", "伊", "宫",
	"宁", "仇", "栾", "暴", "甘", "钭", "厉", "戎", "祖", "武", "符", "刘", "景", "詹", "束", "龙", "叶", "幸", "司", "韶",
	"郜", "黎", "蓟", "薄", "印", "宿", "白", "怀", "蒲", "邰", "从", "鄂", "索", "咸", "籍", "赖", "卓", "蔺", "屠", "蒙",
	"池", "乔", "阴", "郁", "胥", "能", "苍", "双", "闻", "莘", "党", "翟", "谭", "贡", "劳", "逄", "姬", "申", "扶", "堵",
	"冉", "宰", "郦", "雍", "舄", "璩", "桑", "桂", "濮", "牛", "寿", "通", "边", "扈", "燕", "冀", "郏", "浦", "尚", "农",
	"温", "别", "庄", "晏", "柴", "瞿", "阎", "充", "慕", "连", "茹", "习", "宦", "艾", "鱼", "容", "向", "古", "易", "慎",
	"戈", "廖", "庾", "终", "暨", "居", "衡", "步", "都", "耿", "满", "弘", "匡", "国", "文", "寇", "广", "禄", "阙", "东",
	"殴", "殳", "沃", "利", "蔚", "越", "夔", "隆", "师", "巩", "厍", "聂", "晁", "勾", "敖", "融", "冷", "訾", "辛", "阚",
	"那", "简", "饶", "空", "曾", "毋", "沙", "乜", "养", "鞠", "须", "丰", "巢", "关", "蒯", "相", "查", "後", "荆", "红",
	"游", "竺", "权", "逯", "盖", "益", "桓", "公", "司马", "上官", "欧阳", "夏侯", "诸葛",
}

var chineseLastNames = []string{
	"澄邈", "德泽", "海超", "海阳", "海荣", "海逸", "海昌", "瀚钰", "瀚文", "涵亮", "涵煦", "明宇",
	"涵衍", "浩皛", "浩波", "浩博", "浩初", "浩宕", "浩歌", "浩广", "浩邈", "浩气", "浩思", "浩言",
	"鸿宝", "鸿波", "鸿博", "鸿才", "鸿畅", "鸿畴", "鸿达", "鸿德", "鸿飞", "鸿风", "鸿福", "鸿光",
	"鸿晖", "鸿朗", "鸿文", "鸿轩", "鸿煊", "鸿骞", "鸿远", "鸿云", "鸿哲", "鸿祯", "鸿志", "鸿卓",
	"嘉澍", "光济", "澎湃", "彭泽", "鹏池", "鹏海", "浦和", "浦泽", "瑞渊", "越泽", "博耘", "德运",
	"辰宇", "辰皓", "辰钊", "辰铭", "辰锟", "辰阳", "辰韦", "辰良", "辰沛", "晨轩", "晨涛", "晨濡",
	"晨潍", "鸿振", "吉星", "铭晨", "起运", "运凡", "运凯", "运鹏", "运浩", "运诚", "运良", "运鸿",
	"运锋", "运盛", "运升", "运杰", "运珧", "运骏", "运凯", "运乾", "维运", "运晟", "运莱", "运华",
	"耘豪", "星爵", "星腾", "星睿", "星泽", "星鹏", "星然", "震轩", "震博", "康震", "震博", "振强",
	"振博", "振华", "振锐", "振凯", "振海", "振国", "振平", "昂然", "昂雄", "昂杰", "昂熙", "昌勋",
	"昌盛", "昌淼", "昌茂", "昌黎", "昌燎", "昌翰", "晨朗", "德明", "德昌", "德曜", "范明", "飞昂",
	"高旻", "晗日", "昊然", "昊天", "昊苍", "昊英", "昊宇", "昊嘉", "昊明", "昊伟", "昊硕", "昊磊",
	"昊东", "鸿晖", "鸿朗", "华晖", "金鹏", "晋鹏", "敬曦", "景明", "景天", "景浩", "俊晖", "君昊",
	"昆锐", "昆卉", "昆峰", "昆颉", "昆谊", "昆皓", "昆鹏", "昆明", "昆杰", "昆雄", "昆纶", "鹏涛",
	"依秋", "依波", "香巧", "紫萱", "涵易", "忆之", "幻巧", "美倩", "安寒", "白亦", "惜玉", "碧春",
	"怜雪", "听南", "念蕾", "紫夏", "凌旋", "芷梦", "凌寒", "梦竹", "千凡", "丹蓉", "慧贞", "思菱",
	"平卉", "笑柳", "雪卉", "南蓉", "谷梦", "巧兰", "绿蝶", "飞荷", "佳蕊", "芷荷", "怀瑶", "慕易",
	"若芹", "紫安", "曼冬", "寻巧", "雅昕", "尔槐", "以旋", "初夏", "依丝", "怜南", "傲菡", "谷蕊",
	"笑槐", "飞兰", "笑卉", "迎荷", "佳音", "梦君", "妙绿", "觅雪", "寒安", "沛凝", "白容", "乐蓉",
	"映安", "依云", "映冬", "凡雁", "梦秋", "梦凡", "秋巧", "若云", "元容", "怀蕾", "灵寒", "天薇",
	"翠安", "乐琴", "宛南", "怀蕊", "白风", "访波", "亦凝", "易绿", "夜南", "曼凡", "亦巧", "青易",
	"冰真", "白萱", "友安", "海之", "小蕊", "又琴", "天风", "若松", "盼菡", "秋荷", "香彤", "语梦",
	"惜蕊", "迎彤", "沛白", "雁彬", "易蓉", "雪晴", "诗珊", "春冬", "晴钰", "冰绿", "半梅", "笑容",
	"沛凝", "映秋", "盼烟", "晓凡", "涵雁", "问凝", "冬萱", "晓山", "雁蓉", "梦蕊", "山菡", "南莲",
	"飞双", "凝丝", "思萱", "怀梦", "雨梅", "冷霜", "向松", "迎丝", "迎梅", "雅彤", "香薇", "以山",
	"碧萱", "寒云", "向南", "书雁", "怀薇", "思菱", "忆文", "翠巧", "书文", "若山", "向秋", "凡白",
	"绮烟", "从蕾", "天曼", "又亦", "从语", "绮彤", "之玉", "凡梅", "依琴", "沛槐", "又槐", "元绿",
}

// GetPerson returns a new Dowser interface of Person struct
func GetPerson() Dowser {
	person := &Person{}
	return person
}

// Person struct
type Person struct {
}

func (p Person) titlemale() string {
	return randomElementFromSliceString(titlesMale)
}

// TitleMale generates random titles for males
func (p Person) TitleMale(v reflect.Value) (interface{}, error) {
	return p.titlemale(), nil
}

// TitleMale get a title male randomly in string ("Mr.", "Dr.", "Prof.", "Lord", "King", "Prince")
func TitleMale(opts ...options.OptionFunc) string {
	return singleFakeData(TitleMaleTag, func() interface{} {
		p := Person{}
		return p.titlemale()
	}, opts...).(string)
}

func (p Person) titleFemale() string {
	return randomElementFromSliceString(titlesFemale)
}

// TitleFeMale generates random titles for females
func (p Person) TitleFeMale(v reflect.Value) (interface{}, error) {
	return p.titleFemale(), nil
}

// TitleFemale get a title female randomly in string ("Mrs.", "Ms.", "Miss", "Dr.", "Prof.", "Lady", "Queen", "Princess")
func TitleFemale(opts ...options.OptionFunc) string {
	return singleFakeData(TitleFemaleTag, func() interface{} {
		p := Person{}
		return p.titleFemale()
	}, opts...).(string)
}

func (p Person) firstname() string {
	return randomElementFromSliceString(firstNames)
}

// FirstName returns first names
func (p Person) FirstName(v reflect.Value) (interface{}, error) {
	return p.firstname(), nil
}

// FirstName get fake firstname
func FirstName(opts ...options.OptionFunc) string {
	return singleFakeData(FirstNameTag, func() interface{} {
		p := Person{}
		return p.firstname()
	}, opts...).(string)
}

func (p Person) firstnamemale() string {
	return randomElementFromSliceString(firstNamesMale)
}

// FirstNameMale returns first names for males
func (p Person) FirstNameMale(v reflect.Value) (interface{}, error) {
	return p.firstnamemale(), nil
}

// FirstNameMale get fake firstname for male
func FirstNameMale(opts ...options.OptionFunc) string {
	return singleFakeData(FirstNameMaleTag, func() interface{} {
		p := Person{}
		return p.firstnamemale()
	}, opts...).(string)
}

func (p Person) firstnamefemale() string {
	return randomElementFromSliceString(firstNamesFemale)
}

// FirstNameFemale returns first names for females
func (p Person) FirstNameFemale(v reflect.Value) (interface{}, error) {
	return p.firstnamefemale(), nil
}

// FirstNameFemale get fake firstname for female
func FirstNameFemale(opts ...options.OptionFunc) string {
	return singleFakeData(FirstNameFemaleTag, func() interface{} {
		p := Person{}
		return p.firstnamefemale()
	}, opts...).(string)
}

func (p Person) lastname() string {
	return randomElementFromSliceString(lastNames)
}

// LastName returns last name
func (p Person) LastName(v reflect.Value) (interface{}, error) {
	return p.lastname(), nil
}

// LastName get fake lastname
func LastName(opts ...options.OptionFunc) string {
	return singleFakeData(LastNameTag, func() interface{} {
		p := Person{}
		return p.lastname()
	}, opts...).(string)
}

func (p Person) name() string {
	if randNameFlag > 50 {
		return fmt.Sprintf("%s %s %s", randomElementFromSliceString(titlesFemale), randomElementFromSliceString(firstNamesFemale), randomElementFromSliceString(lastNames))
	}
	return fmt.Sprintf("%s %s %s", randomElementFromSliceString(titlesMale), randomElementFromSliceString(firstNamesMale), randomElementFromSliceString(lastNames))
}

// Name returns a random name
func (p Person) Name(v reflect.Value) (interface{}, error) {
	return p.name(), nil
}

// Name get fake name
func Name(opts ...options.OptionFunc) string {
	return singleFakeData(NAME, func() interface{} {
		p := Person{}
		return p.name()
	}, opts...).(string)
}

// Gender returns a random gender
func (p Person) Gender(v reflect.Value) (interface{}, error) {
	return p.gender(), nil
}

func (p Person) gender() string {
	return randomElementFromSliceString(genders)
}

// Gender get fake gender
func Gender(opts ...options.OptionFunc) string {
	return singleFakeData(GENDER, func() interface{} {
		p := Person{}
		return p.gender()
	}, opts...).(string)
}

// ChineseFirstName returns a random chinese first name
func (p Person) ChineseFirstName(v reflect.Value) (interface{}, error) {
	return p.chineseFirstName(), nil
}

func (p Person) chineseFirstName() string {
	return randomElementFromSliceString(chineseFirstNames)
}

// ChineseFirstName get chinese first name
func ChineseFirstName(opts ...options.OptionFunc) string {
	return singleFakeData(ChineseFirstNameTag, func() interface{} {
		p := Person{}
		return p.chineseFirstName()
	}, opts...).(string)
}

// ChineseLastName returns a random chinese last name
func (p Person) ChineseLastName(v reflect.Value) (interface{}, error) {
	return p.chineseLastName(), nil
}

func (p Person) chineseLastName() string {
	return randomElementFromSliceString(chineseLastNames)
}

// ChineseLastName get chinese lsst name
func ChineseLastName(opts ...options.OptionFunc) string {
	return singleFakeData(ChineseLastNameTag, func() interface{} {
		p := Person{}
		return p.chineseLastName()
	}, opts...).(string)
}

// ChineseName returns a random nhinese name
func (p Person) ChineseName(v reflect.Value) (interface{}, error) {
	return p.chineseName(), nil
}

func (p Person) chineseName() string {
	return fmt.Sprintf("%s%s", randomElementFromSliceString(chineseFirstNames), randomElementFromSliceString(chineseLastNames))
}

// ChineseName get chinese lsst name
func ChineseName(opts ...options.OptionFunc) string {
	return singleFakeData(ChineseNameTag, func() interface{} {
		p := Person{}
		return p.chineseName()
	}, opts...).(string)
}
