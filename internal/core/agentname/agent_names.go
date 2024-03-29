package agentname

import "math/rand"

var names []string

func Generate() string {
	return names[rand.Intn(len(names))]
}

func init() {
	names = []string{
		"101",
		"180",
		"360",
		"44",
		"66",
		"11",
		"22",
		"33",
		"55",
		"77",
		"88",
		"99",
		"00",
		"Achilles Mountain",
		"Acid Gosling",
		"Admiral Tot",
		"Alias",
		"Alley Cat",
		"Alley Fiend",
		"Alley Frog",
		"Alpha",
		"Angler",
		"Apex",
		"Apple Nola",
		"Apple",
		"Are Ess Tee",
		"Athens Fire",
		"Aurora",
		"Bacon",
		"Badger",
		"Badminton",
		"Bambi",
		"Bang Shift",
		"Barbwire",
		"Bat Magenta",
		"Beekeeper",
		"Beetle King",
		"Bentley",
		"Beo",
		"Betty Cricket",
		"Biscuit",
		"Bitmap",
		"Firefly",
		"Mustard",
		"Walnut",
		"Blackfire",
		"Blackout",
		"Blade",
		"Bleach",
		"Bleachers",
		"Bleeker",
		"Blink",
		"Blinker",
		"Bliss",
		"Bobcat",
		"Boost",
		"Bowie",
		"Bowler",
		"Brandy",
		"Breadmaker",
		"Breakfast",
		"Brick Mooch",
		"Bridge Whip",
		"Broomdog",
		"Brown",
		"Brunden",
		"Buckshot",
		"Bug Blitz",
		"Bug Fire",
		"Burn",
		"Butcher",
		"Butters",
		"Canary",
		"Carrot",
		"Cereal",
		"Chameleon",
		"Chaos",
		"Chapstick",
		"Charms",
		"Chew Chew",
		"Chuckles",
		"Cinder Coffee",
		"Clang Glyph",
		"Claws",
		"Clover",
		"Cocoa",
		"Coffee",
		"Coldy",
		"Congo Wire",
		"Cool Whip",
		"Cosmo",
		"Crash Test",
		"Cricket",
		"Criss Cross",
		"Cross Thread",
		"Crumb Cake",
		"Cuffs",
		"Cumulo",
		"Cupid Dust",
		"Cupid",
		"Curio",
		"Curio",
		"Cutlass",
		"D-Hog-Day",
		"DARK HQ",
		"DZE",
		"Daffy Neo",
		"Dahlia Bumble",
		"Dahlia",
		"Dahlia",
		"Daisy Stick",
		"Dakota Bliss",
		"Dallas Burn",
		"Dallas Foxface",
		"Dance Bloom",
		"Dance Cannon",
		"Dandelion",
		"Danger Menace",
		"Daze",
		"Apollo",
		"Hera",
		"Artemis",
		"Sigma",
		"Delta",
		"Prometheus",
		"Thyme",
		"Oregano",
		"Ivy",
		"Monstera",
		"Cotton",
		"Cambria",
		"North",
		"South",
		"West",
		"East",
		"Papers",
		"Keys",
		"Ridgeback",
		"Dark Burn",
		"Dark Horse",
		"Dark Matter",
		"Darkside",
		"Hecate",
		"Staples",
		"Eisenhower",
		"Darth 44",
		"Day Hawk",
		"Deano",
		"Demo Zero",
		"Demo",
		"Desert Haze",
		"Despair",
		"Dez Bonbon",
		"Dez North",
		"Dez",
		"Diamond",
		"Digger",
		"Digger",
		"Digital Equinox",
		"Minimus",
		"Moonshine",
		"Disco",
		"Domino",
		"Don Stab",
		"Double Eerie",
		"Dove Dolce",
		"Doz",
		"Dragon Blood",
		"Drift",
		"Duck Duck",
		"Easy Mac",
		"Easy Street",
		"Eerie",
		"Eight Patrol",
		"Electric Saturn",
		"Ember Rope",
		"Emerald",
		"Emerald",
		"Energy",
		"Engine Eye",
		"Essex",
		"Fabulous",
		"Fadey",
		"Fast Draw",
		"Fennel Dove",
		"Feral Cookie",
		"Finish",
		"Fire Bite",
		"Fire Fish",
		"Firedog",
		"Firefly",
		"Flakes",
		"Flash",
		"Flashpoint",
		"Flint",
		"Flyswat Briggs",
		"Force",
		"Forger",
		"Freesia",
		"Freeway",
		"Frenzy",
		"Fresh Peper",
		"Friday Fox",
		"Friday",
		"Frosty Snazz",
		"Frosty Squid",
		"Frosty Sunshine",
		"Frosty",
		"Bean",
		"Ghost",
		"Ghost",
		"Glyph",
		"Glyph",
		"Goldman",
		"Grabber",
		"Granola Dove",
		"Granola",
		"Grin",
		"Grinch",
		"Gullyway",
		"Guncap Slingbad",
		"Gunhawk",
		"Hairpin",
		"Hemlock",
		"Hiccup",
		"Hidden Tree",
		"High Beam",
		"High Fructose",
		"Highlander",
		"Hightower",
		"Highway",
		"Hitch Frenzy",
		"Homerun",
		"Hoover Spark",
		"Houston Rocket",
		"Houston",
		"Howitzer Rise",
		"Howling Swede",
		"Hyper Kong",
		"Hyper",
		"Ice",
		"Impulse",
		"Indestructible Potato",
		"Indiana",
		"Indigo Red",
		"Iron Butterfly",
		"Iron Jesus",
		"Jack Cassidy",
		"Jade Fox",
		"Jelly Camber",
		"Jelly",
		"Jersey",
		"Jester",
		"Jig Kraken",
		"Jigsaw",
		"Jo Jo Spooky",
		"Journeyman",
		"Judge",
		"Kawaii Red",
		"Keystone",
		"Kickstart",
		"Kill Switch",
		"Kingfisher",
		"Knight Light",
		"Knock Out Star",
		"Koi",
		"Kowhai",
		"Ladybug",
		"Lava",
		"Leaf Assassin",
		"Legacy",
		"LifeRobber",
		"Lilac Lizard",
		"Lime",
		"London Fox",
		"Loot",
		"Lope Lope",
		"Low Menace",
		"Low Voltage",
		"Lucifurious",
		"Lucky",
		"Lucky",
		"Mad Robin",
		"Magenta",
		"Manhattan",
		"Marbles",
		"Marigold Loot",
		"Marshmallow Treat",
		"Marshmallow",
		"Mauve Cactus",
		"Mayhem",
		"Mercury",
		"Metal Aphrodite",
		"Metal Star",
		"Micro Star",
		"Microwave Chardonnay",
		"Microwave",
		"Midnight Bat",
		"Midnight",
		"Mirage",
		"Mirage",
		"Monk",
		"Moon Cricket",
		"Moon Laser",
		"Moon Orchid",
		"Moon Peaches",
		"Moon Radar",
		"Moon",
		"Bloom",
		"Mud Pie",
		"Mule Lock",
		"Libre",
		"Aries",
		"Capsicum",
		"Kumura",
		"Murmur",
		"Mustard Centaur",
		"Mustard",
		"Muzish",
		"Nacho",
		"Natural Gold",
		"Necessary Momentum",
		"Nem X",
		"Neo Germal",
		"Nessie",
		"New Magoo",
		"Nibbler",
		"Night Dream",
		"Night Magnet",
		"Night Train",
		"Nine Lives",
		"Noh Noh",
		"Nola",
		"Nova",
		"Nueva Nova",
		"Nutmeg Riot",
		"Nutmeg",
		"Oblivion",
		"Old Felix",
		"Osprey",
		"Overrun",
		"Fractal",
		"Rapids",
		"Parallax",
		"Paris Boost",
		"Pecan",
		"Pegasus",
		"Pepper Mouse",
		"Pepper",
		"Pepper Burst",
		"Peacock",
		"Flurry",
		"Zuboff",
		"Phoenix Sparrow",
		"Phoenix Tetra",
		"Pinball",
		"Pinball Wizard",
		"Tetra",
		"Wind Hopper",
		"Rock Hopper",
		"Penguin",
		"Pink Hopper",
		"Pink Nightmare",
		"Pink Stream",
		"Riddle",
		"Pistachio",
		"Pistol Hydro",
		"Pitfall Whiskers",
		"Pitfall",
		"Piwakawaka",
		"Pixels",
		"Plenty Orange",
		"Plum Moon",
		"Pluto",
		"Poison",
		"Pockets",
		"Polar Bee",
		"Pompeii",
		"Pop Bee",
		"Rami",
		"Ordo",
		"Uni",
		"Knoll",
		"Popeye",
		"Wipeout",
		"Poppy Coffee",
		"Parker",
		"Prometheus",
		"Prone",
		"Proper",
		"Queen Bee",
		"Tank",
		"Red Delicious",
		"Red Heroine",
		"Earl Grey",
		"Red Pepper",
		"Referee",
		"Slugger",
		"Monarch",
		"Riff Raff",
		"Rimu",
		"Rink Ruler",
		"Risen",
		"Roadblock",
		"Roadspike",
		"Rocky Highway",
		"Roller Turtle",
		"Roma Kabuki",
		"Romane",
		"Rook",
		"Rooster",
		"Rope",
		"Rosie Bird",
		"Round Kick Boomer",
		"Rum Punch",
		"Rummy Stickers",
		"Runway Darling",
		"Rusty Vortex",
		"Saber Red",
		"Sand Whiskers",
		"Sandbox",
		"Santa's Little Helper",
		"Sapiens",
		"Saturn Extreme",
		"Saturnalia",
		"Scare Stone",
		"Scarlet Mary",
		"Scooby Did",
		"Scrapper",
		"Scrapple",
		"Scratch",
		"Scuffs",
		"Scuttlebutt",
		"Seal Snake",
		"Seattle Jay",
		"Serendipity",
		"Shade Nightman",
		"Shadow Bishop",
		"Shadow Gal",
		"ShadowDancer",
		"ShadowFax",
		"Shady Prairie",
		"Shamrock",
		"Shamrock",
		"Shay",
		"Sherm",
		"Sherwood",
		"Shield",
		"Shimmy Shammy",
		"Ship Whip",
		"Shivers",
		"Short Firecracker",
		"Show Boat",
		"Show Off",
		"Shwatson",
		"Shy Warrior",
		"Sick Saurus",
		"Silent",
		"Silver Agent",
		"Silver Cup",
		"Silver Rose",
		"Sir Shark",
		"Sir Shove",
		"Sir Squire",
		"Skirble",
		"Skittle Mine",
		"Skittle",
		"Skool",
		"Sky Dahlia",
		"Sky Herald",
		"Sky Trinity",
		"Sky",
		"Cancannon",
		"Skylark",
		"Slacker Cat",
		"Sleek Assassin",
		"Sleek Zelda",
		"Sleepwalker",
		"Slinger",
		"Sloth",
		"Sly Silvermoon",
		"SmartieQuest",
		"Smurf",
		"Snake Eyes",
		"Snapdragon",
		"Snapple Whistler",
		"Snow Cream",
		"Snow Pharaoh",
		"Snow",
		"Solitaire",
		"Solitaire",
		"Solitaire",
		"Speedwell",
		"Spell",
		"Sphinx",
		"Spitfire",
		"Spoiler",
		"Squatch",
		"Squeeble",
		"Stallion Patton",
		"Jammer",
		"Scratch",
		"Frodo",
		"Zen",
		"Starshine",
		"Gopher",
		"Stealth",
		"Cut Toe",
		"Ginger Steel",
		"Steel Heart",
		"Steel Solstice",
		"Steel",
		"Stick Shift",
		"Stickers",
		"Stone Boomstick",
		"Storm Cake",
		"Storm Master",
		"Street Jolly",
		"Street Squirrel",
		"Stud Buster",
		"Subzero Taffy",
		"Succubus In Training",
		"Sugar",
		"Lemon",
		"Sun Runner",
		"Super Flick",
		"Supernova",
		"Swampmasher",
		"Sweep",
		"Swing Setter",
		"Switch",
		"Tabasco Dracula",
		"Tacklebox",
		"Taffy",
		"Tangerine",
		"Taz Ringer",
		"Tea Kettle",
		"Third Moon",
		"Tin Fox",
		"Tin Mutt",
		"Tokyo",
		"Tomcat",
		"Toronto",
		"Toy Dogwatch",
		"Toy Peep",
		"Toy Town",
		"Trip",
		"Troubadour",
		"Troublemasher",
		"Tui",
		"Turnip",
		"Tweety",
		"Twiddle Twix",
		"Twin Blaze",
		"Twinkle",
		"Twister Hero",
		"Twisty Dew",
		"Twisty Freesia",
		"Twitch",
		"Twix Bond",
		"Twix Esses",
		"Underdog",
		"Undergrad",
		"Universe Bullet",
		"Vice Swerve",
		"Virgo Moon",
		"Vortex",
		"Washer",
		"Wheels",
		"Whipsaw",
		"Whistler",
		"Wildcat Appaloosa",
		"Wildcat Talent",
		"Williams",
		"Wing",
		"Wings",
		"Winter Bite",
		"Wolf Tribune",
		"Wonka",
		"Woodland Beauty",
		"Wrangler Jim",
		"Yellow",
		"Yellowjacket",
		"Zero Corn",
		"Zod",
		"Zorkle Sporkle",
	}
}
