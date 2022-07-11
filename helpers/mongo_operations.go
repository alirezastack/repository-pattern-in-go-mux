package helpers

//type MongoDBRepository struct {
//}

//func (mo MongoDBRepository) read() {
//	fmt.Println("read MongoDB")
//}
//
//func (mo MongoDBRepository) update() {
//	fmt.Println("update MongoDB")
//}
//
//func (mo MongoDBRepository) delete() {
//	fmt.Println("delete MongoDB")
//}
//
//func (mo MongoDBRepository) insert() {
//	fmt.Println("insert MongoDB")
//}
//
//func ConnectMongoDB() (*mongo.Client, context.CancelFunc) {
//	client, err := mongo.NewClient(options.Client().ApplyURI(configs.MongoURI()))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//
//	err = client.Connect(ctx)
//
//	err = client.Ping(ctx, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Print("Successfully connected to MongoDB!")
//
//	return client, cancel
//}
