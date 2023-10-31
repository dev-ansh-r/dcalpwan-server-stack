package blocklist

var builtin = New(
	"about",
	"abuse",
	"access",
	"account",
	"accounts",
	"activate",
	"activities",
	"activity",
	"add",
	"address",
	"adm",
	"admin",
	"administration",
	"administrator",
	"administrators",
	"admins",
	"ads",
	"adult",
	"advertising",
	"affiliate",
	"affiliates",
	"ajax",
	"alerts",
	"all",
	"alpha",
	"analysis",
	"analytics",
	"android",
	"anon",
	"anonymous",
	"api",
	"api-key",
	"api-keys",
	"apikey",
	"apikeys",
	"app",
	"application",
	"applications",
	"applicationserver",
	"apps",
	"archive",
	"archives",
	"article",
	"asct",
	"asset",
	"assets",
	"atom",
	"auth",
	"authentication",
	"avatar",
	"backup",
	"balancer-manager",
	"banner",
	"banners",
	"basic-station",
	"basicstation",
	"batch",
	"beta",
	"billing",
	"bin",
	"blog",
	"blogs",
	"board",
	"book",
	"bookmark",
	"bot",
	"bots",
	"broker",
	"bug",
	"bulk",
	"business",
	"cache",
	"cadastro",
	"calendar",
	"call",
	"campaign",
	"cancel",
	"captcha",
	"career",
	"careers",
	"carreers",
	"cart",
	"categories",
	"category",
	"cdn",
	"cgi",
	"cgi-bin",
	"changelog",
	"chat",
	"check",
	"checking",
	"checkout",
	"cli",
	"client",
	"cliente",
	"clients",
	"cloud",
	"cluster",
	"code",
	"codereview",
	"collaborator",
	"collaborators",
	"comercial",
	"comment",
	"comments",
	"communities",
	"community",
	"company",
	"compare",
	"compras",
	"conference",
	"config",
	"configuration",
	"connect",
	"connection",
	"console",
	"contact",
	"contact-us",
	"contactus",
	"contest",
	"contribute",
	"corp",
	"create",
	"css",
	"cups",
	"customers",
	"dashboard",
	"data",
	"dedicated",
	"dedicated-cloud",
	"default",
	"delete",
	"demo",
	"design",
	"designer",
	"destroy",
	"dev",
	"devel",
	"developer",
	"developers",
	"device",
	"devices",
	"diagram",
	"diary",
	"dict",
	"dictionary",
	"die",
	"dir",
	"direct-messages",
	"directory",
	"dist",
	"dns",
	"doc",
	"docs",
	"documentation",
	"domain",
	"download",
	"downloads",
	"ecommerce",
	"edit",
	"editor",
	"editors",
	"edu",
	"education",
	"email",
	"employment",
	"empty",
	"end",
	"end-device",
	"end-devices",
	"enterprise",
	"entries",
	"entry",
	"error",
	"errors",
	"eval",
	"event",
	"events",
	"example",
	"example-app",
	"example-application",
	"example-device",
	"example-organisation",
	"example-organization",
	"exampleapp",
	"exampleapplication",
	"exampledevice",
	"exampleorganisation",
	"exampleorganization",
	"examples",
	"exit",
	"explore",
	"extension",
	"extensions",
	"facebook",
	"faq",
	"favorite",
	"favorites",
	"feature",
	"features",
	"feed",
	"feedback",
	"feeds",
	"file",
	"files",
	"first",
	"flash",
	"fleet",
	"fleets",
	"flog",
	"follow",
	"followers",
	"following",
	"forgot",
	"form",
	"forum",
	"forums",
	"founder",
	"free",
	"frequency-plan",
	"frequency-plans",
	"friend",
	"friends",
	"ftp",
	"gadget",
	"gadgets",
	"game",
	"games",
	"gateway",
	"gateway-configuration-server",
	"gatewayagent",
	"gatewayconfigurationserver",
	"gatewayconfserver",
	"gateways",
	"gatewayserver",
	"gcs",
	"generic-node",
	"genericnode",
	"get",
	"ghost",
	"gift",
	"gifts",
	"gist",
	"git",
	"github",
	"graph",
	"group",
	"groups",
	"gtw",
	"gtws",
	"guest",
	"guests",
	"handler",
	"help",
	"home",
	"homepage",
	"homes",
	"host",
	"hosting",
	"hostmaster",
	"hostname",
	"howto",
	"hpg",
	"html",
	"http",
	"httpd",
	"https",
	"iamges",
	"icon",
	"icons",
	"idea",
	"ideas",
	"identity",
	"identityserver",
	"image",
	"images",
	"imap",
	"img",
	"index",
	"indice",
	"info",
	"information",
	"inquiry",
	"instagram",
	"intranet",
	"invitations",
	"invite",
	"ipad",
	"iphone",
	"irc",
	"issue",
	"issues",
	"item",
	"items",
	"java",
	"javascript",
	"job",
	"jobs",
	"join",
	"joinserver",
	"json",
	"jump",
	"knowledgebase",
	"language",
	"languages",
	"last",
	"ldap-status",
	"learn",
	"legal",
	"license",
	"link",
	"links",
	"linux",
	"list",
	"lists",
	"local",
	"localdomain",
	"localhost",
	"log",
	"log-in",
	"log-out",
	"login",
	"logout",
	"logs",
	"lora",
	"lorawan",
	"lorawan-stack",
	"mac",
	"mail",
	"mail1",
	"mail2",
	"mail3",
	"mail4",
	"mail5",
	"mailer",
	"mailing",
	"maintenance",
	"manager",
	"manual",
	"map",
	"maps",
	"marketing",
	"marketplace",
	"master",
	"masters",
	"media",
	"member",
	"members",
	"message",
	"messages",
	"messenger",
	"metrics",
	"microblog",
	"microblogs",
	"mine",
	"mis",
	"mob",
	"mobile",
	"monitor",
	"monitoring",
	"moderator",
	"moderators",
	"movie",
	"movies",
	"mp3",
	"msg",
	"msn",
	"music",
	"musicas",
	"myself",
	"mysql",
	"name",
	"named",
	"namespace",
	"nan",
	"navi",
	"navigation",
	"net",
	"network",
	"networkserver",
	"new",
	"news",
	"newsletter",
	"nick",
	"nickname",
	"noc",
	"notes",
	"noticias",
	"notification",
	"notifications",
	"notify",
	"ns1",
	"ns10",
	"ns2",
	"ns3",
	"ns4",
	"ns5",
	"ns6",
	"ns7",
	"ns8",
	"ns9",
	"null",
	"oauth",
	"oauth-clients",
	"offer",
	"offers",
	"official",
	"old",
	"online",
	"open-source",
	"openid",
	"opensource",
	"operator",
	"ops",
	"order",
	"orders",
	"org",
	"organization",
	"organizations",
	"orgs",
	"overview",
	"owner",
	"owners",
	"packetbroker",
	"page",
	"pager",
	"pages",
	"panel",
	"partners",
	"password",
	"payment",
	"peer",
	"peers",
	"perl",
	"phone",
	"photo",
	"photoalbum",
	"photos",
	"php",
	"phpmyadmin",
	"phppgadmin",
	"phpredisadmin",
	"pic",
	"pics",
	"ping",
	"plan",
	"plans",
	"plugin",
	"plugins",
	"policy",
	"pop",
	"pop3",
	"popular",
	"portal",
	"post",
	"postfix",
	"postmaster",
	"posts",
	"premium",
	"press",
	"price",
	"pricing",
	"privacy",
	"privacy-policy",
	"privacypolicy",
	"private",
	"prod",
	"product",
	"production",
	"products",
	"profile",
	"project",
	"projects",
	"promo",
	"pub",
	"public",
	"purpose",
	"put",
	"python",
	"query",
	"random",
	"ranking",
	"read",
	"readme",
	"recent",
	"recruit",
	"recruitment",
	"redirect",
	"redirect-uri",
	"register",
	"registration",
	"release",
	"remove",
	"replies",
	"report",
	"reports",
	"repositories",
	"repository",
	"req",
	"request",
	"requests",
	"reset",
	"roc",
	"root",
	"router",
	"rss",
	"ruby",
	"rule",
	"sag",
	"sale",
	"sales",
	"sample",
	"samples",
	"save",
	"school",
	"script",
	"scripts",
	"search",
	"secure",
	"security",
	"self",
	"send",
	"server",
	"server-info",
	"server-status",
	"service",
	"services",
	"session",
	"sessions",
	"setting",
	"settings",
	"setup",
	"share",
	"shop",
	"show",
	"sign-in",
	"sign-up",
	"signin",
	"signout",
	"signup",
	"site",
	"sitemap",
	"sites",
	"smartphone",
	"smtp",
	"soporte",
	"source",
	"spec",
	"special",
	"sql",
	"src",
	"ssh",
	"ssl",
	"ssladmin",
	"ssladministrator",
	"sslwebmaster",
	"staff",
	"stage",
	"staging",
	"start",
	"stat",
	"state",
	"static",
	"statistics",
	"stats",
	"status",
	"store",
	"stores",
	"stories",
	"style",
	"styleguide",
	"stylesheet",
	"stylesheets",
	"subdomain",
	"subscribe",
	"subscriptions",
	"super",
	"superuser",
	"suporte",
	"support",
	"svn",
	"swf",
	"sys",
	"sysadmin",
	"sysadministrator",
	"system",
	"tablet",
	"tablets",
	"tag",
	"talk",
	"task",
	"tasks",
	"team",
	"teams",
	"tech",
	"telnet",
	"term",
	"terms",
	"terms-of-service",
	"termsofservice",
	"test",
	"test1",
	"test2",
	"test3",
	"teste",
	"testing",
	"tests",
	"the-things-gateway",
	"the-things-industries",
	"the-things-network",
	"the-things-products",
	"the-things-stack",
	"the-thingsindustries",
	"the-thingsnetwork",
	"the-thingsproducts",
	"the-thingsstack",
	"theme",
	"themes",
	"thethings",
	"thethings-gateway",
	"thethings-industries",
	"thethings-network",
	"thethings-products",
	"thethings-stack",
	"thethingsgateway",
	"thethingsindustries",
	"thethingsnetwork",
	"thethingsproducts",
	"thethingsstack",
	"things",
	"this",
	"thread",
	"threads",
	"tls",
	"tmp",
	"todo",
	"tool",
	"tools",
	"top",
	"topic",
	"topics",
	"tos",
	"tour",
	"translations",
	"trends",
	"tti",
	"ttn",
	"ttn-lw-cli",
	"ttn-lw-stack",
	"ttnctl",
	"tts",
	"tutorial",
	"tux",
	"twitter",
	"undef",
	"unfollow",
	"unsubscribe",
	"update",
	"upload",
	"uploads",
	"url",
	"usage",
	"user",
	"username",
	"users",
	"usuario",
	"vendas",
	"ver",
	"version",
	"video",
	"videos",
	"visitor",
	"watch",
	"weather",
	"web",
	"webhook",
	"webhooks",
	"webmail",
	"webmaster",
	"website",
	"websites",
	"webui",
	"welcome",
	"widget",
	"widgets",
	"wiki",
	"win",
	"windows",
	"word",
	"work",
	"works",
	"workshop",
	"wws",
	"www",
	"www1",
	"www2",
	"www3",
	"www4",
	"www5",
	"www6",
	"www7",
	"wwws",
	"wwww",
	"xfn",
	"xml",
	"xmpp",
	"xpg",
	"xxx",
	"yaml",
	"year",
	"yml",
	"you",
	"your-app",
	"your-application",
	"your-device",
	"your-domain",
	"your-name",
	"your-organisation",
	"your-organization",
	"your-site",
	"your-user-name",
	"your-username",
	"yourapp",
	"yourapplication",
	"yourdevice",
	"yourdomain",
	"yourname",
	"yourorganisation",
	"yourorganization",
	"yoursite",
	"yourusername",
)
