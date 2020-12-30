package cns

var Config = map[string]string{
    "MAXPROCS":"MAX",
    "Mongodb":"mongodb://user:cerberus789@RSLINMNMDB2VP:27017,RSLINMNMDB3VP:27017/religare?replicaSet=pfReplicaSet",
    "Mongodb_Odin":"mongodb://ganymedecg:ganymed#2468O@RSLINMNMMOPPVP:27021,RSLINMNMMOPSVP:27021/Religare?replicaSet=ReplPortal",
    "Mongodb_Inmemory": "mongodb://ruser:Rpfp%40ss876@rslinmnmsomsvpha:29015,rslinmnmsrmsvpha:29015,rslinmnmspovpha:29015/admin?replicaSet=Techyon2",
    "Redis":"192.168.90.122:16379, 192.168.90.123:16380, 192.168.90.124:16381",
	"ServerHost":":8080",
	"AuthUrl":"https://secure.religareonline.com/FundAPI/Default.aspx",
	"AuthenticationEnabled":"0",
	"SnapQuoteFilePath": "/home/dynami/price_feed_logs/snap_quote.log",
	"Wss":"wssn.religareonline.com:443", //"wss://stgllfc.religareonline.com:443",
	"Source":"Techyon",
}

var KafkaTopics = map[string]string{
	"SnapQuote": "priceFeed",
}