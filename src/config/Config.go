package config

var Config = map[string]string{
   "ServerHost":"206.189.132.17:8080",
   "AuthUrl":"https://secure.religareonline.com/FundAPI/Default.aspx",
   "AuthenticationEnabled":"0",
   "NsePriceFeed":"wss://eftlpf.religareonline.com:9291/nseChannel", // eftl channels for pricefeed
   "BsePriceFeed":"wss://eftlpf.religareonline.com:9291/bseChannel", // eftl channels for pricefeed
   "McxPriceFeed":"wss://eftlpf.religareonline.com:9291/mcxChannel", // eftl channels for pricefeed
   "NcdexPriceFeed":"wss://eftlpf.religareonline.com:9291/ncdexChannel", // eftl channels for pricefeed
   "DPRUpdate":"wss://eftlpf.religareonline.com:9291/DPRUpdate", // dpr update
   "EFTLUser":"bcast",
   "EFTLPassword":"bcast123",
   "Compression":"",
   "TemplatesPath":"/app_data/price_feed_websocket/src/templates",
   "KafkaBrokers": "192.168.90.119:9092, 192.168.90.120:9092, 192.168.90.121:9092",
   "RethinkClusters":"192.168.90.32:28015, 192.168.90.33:28015, 192.168.90.34:28015",
   "RethinkDatabase":"Techyon",
}