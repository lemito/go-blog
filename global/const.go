package global

const (
	dburi       = "mongodb://Ivan:maz14105@testcluster-shard-00-00.vcevl.mongodb.net:27017,testcluster-shard-00-01.vcevl.mongodb.net:27017,testcluster-shard-00-02.vcevl.mongodb.net:27017/myFirstDatabase?ssl=true&replicaSet=atlas-qbnwmi-shard-0&authSource=admin&retryWrites=true&w=majority"
	dbname      = "blog-application"
	performance = 100
)

var (
	jwtSecret = []byte("blogSecret")
)