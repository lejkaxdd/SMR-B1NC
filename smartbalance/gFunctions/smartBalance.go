package smartBalance

import (
	"context"
	"fmt"
	"os"
	"smb/app/handler"
	"smb/app/repository"
	"smb/pkg/api"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type GRPCserver struct {
	api.UnimplementedSmartBalanceServiceServer
}


// gRPC funcs
 
//  = repository.Config{
// 	Host:     "172.26.0.2",
// 	Port:     "5433",
// 	Username: "postgres",
// 	DBName:   "smartbalance",
// 	Password: "secret",
// 	SSLMode:  "disable",
// }



// Insert info from CoolingSystem
func (s *GRPCserver) CoolingSystem(ctx context.Context, req *api.CoolingSystemRequest) (*api.CoolingSystemResponse, error) {

	db, err := repository.NewPangolinDB(repository.Config{
		Host:     "172.26.0.2",
		Port:     "5433",
		Username: "postgres",
		DBName:   "smartbalance",
		Password: "secret",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()


	row, err := db.Query(`CREATE TABLE IF NOT EXISTS coolingsystem( id varchar(70) PRIMARY KEY, CoolingLevel varchar(70) NOT NULL, CoolingFrequency varchar(70) NOT NULL, CoolingType varchar(70) NOT NULL);`)
	if err != nil {
		log.Printf("failed to create table: %s", err.Error())
	}
	// fmt.Println(row)
	defer row.Close()

	id := handler.GenerateUUID()
	stmt, prepErr := db.Prepare("INSERT INTO coolingsystem (id, CoolingLevel, CoolingFrequency, CoolingType) VALUES ($1, $2, $3, $4)")

	if prepErr == nil {
		_, execErr := stmt.Exec(id, req.GetInfo().Coolinglevel, req.GetInfo().Coolingfrequency, req.GetInfo().Coolingtype)
		defer stmt.Close()
		
		if execErr != nil {
			log.Println("error while inserting values", stmt)
		}
	}

	
	myfile, e := os.OpenFile("/application/cooling_audit.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if e != nil {
		log.Printf("Problem with creating file: %s", e)
	}

	defer myfile.Close()
	data := []byte(fmt.Sprintf("{Record="+id+"\tCoolingLevel = %s\tCoolingFrequency = %s\tCoolingType= %s}\n", req.GetInfo().Coolinglevel, req.GetInfo().Coolingfrequency, req.GetInfo().Coolingtype))
	myfile.Write(data)
	fmt.Printf("\nData %s successfully written to file\n", data)

	return &api.CoolingSystemResponse{Record: id}, nil
}


// Check info from CoolingSystem
func (s *GRPCserver) CoolingSystemCheck(ctx context.Context, req *api.CoolingSystemGetRequest) (*api.CoolingSystemGetResponse, error){

	db, err := repository.NewPangolinDB(repository.Config{
		Host:     "172.26.0.2",
		Port:     "5433",
		Username: "postgres",
		DBName:   "smartbalance",
		Password: "secret",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	stmt, prepErr := db.Prepare("SELECT CoolingType FROM coolingsystem WHERE id = $1")

	var value string
	if prepErr == nil {
		
		scanErr := stmt.QueryRow(req.Record).Scan(&value)
		defer stmt.Close()
		fmt.Println(scanErr)
	}

	return &api.CoolingSystemGetResponse{Value: value}, nil
}



// Create User
func (s *GRPCserver) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error){

	db, err := repository.NewPangolinDB(repository.Config{
		Host:     "172.26.0.2",
		Port:     "5433",
		Username: "postgres",
		DBName:   "smartbalance",
		Password: "secret",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	id := handler.GenerateUUID()
	row, err := db.Query(`CREATE TABLE IF NOT EXISTS users( id varchar(70) PRIMARY KEY, Username varchar(70) NOT NULL, Password varchar(70) NOT NULL);`)
	if err != nil {
		log.Printf("failed to create table: %s", err.Error())
	}
	// fmt.Println(row)
	defer row.Close()

	stmt, prepErr := db.Prepare("INSERT INTO users (id, Username, Password) VALUES ($1, $2, $3)")

	if prepErr == nil {
		_, execErr := stmt.Exec(id, req.GetInfo().Username, req.GetInfo().Password)
		defer stmt.Close()
		
		if execErr != nil {
			log.Println("error while inserting values", stmt)
			return &api.CreateUserResponse{Confirm: "User hasn't created"}, err
		}
	}

	return &api.CreateUserResponse{Confirm: "User has successfully created"}, err
}


func (s *GRPCserver) CheckUser(ctx context.Context, req *api.CheckUserRequest) (*api.CheckUserResponse, error){

	db, err := repository.NewPangolinDB(repository.Config{
		Host:     "172.26.0.2",
		Port:     "5433",
		Username: "postgres",
		DBName:   "smartbalance",
		Password: "secret",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	stmt, prepErr := db.Prepare("SELECT Password FROM users WHERE Username = $1")

	var passwd string
	var token string
	// var token boolean
	if prepErr == nil {
		
		scanErr := stmt.QueryRow(req.GetInfo().Username).Scan(&passwd)
		defer stmt.Close()

		if scanErr == nil {
			
			fmt.Println(passwd)
			if passwd == req.GetInfo().Password{
				// token, err = generateJWT(req.Info.GetUsername())

				// if err != nil{
				// 	log.Println(err)
				// }

				// return &api.CheckUserResponse{Token: token}, nil
				token = "1"
				return &api.CheckUserResponse{Token: token}, nil

			} else {

				token = "0"
				return &api.CheckUserResponse{Token: token}, nil
			}

		}

	}
	return &api.CheckUserResponse{Token: "Failed to create token"}, nil	
}
	




/*func checkUser(username string, password string) ([]string, error) {

	//DB connect
	db, err := repository.NewPangolinDB(repository.Config{
		Host:     "172.26.0.2",
		Port:     "5433",
		Username: "postgres",
		DBName:   "smartbalance",
		Password: "secret",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// 	// Checking table exist
	var dbs []string
	// 	// id := handler.GenerateUUID()
	// 	row, err := db.Query(`CREATE TABLE IF NOT EXISTS users(
	// 		"id" varchar(70) PRIMARY KEY,
	// 		"username" varchar(50) NOT NULL,
	// 		"password" varchar(50) NOT NULL);`)

	// 	fmt.Println(row.Next(), "    reere    ", err)
	// 	if err != nil{
	// 		log.Println(err)
	// 	}
	// //
	// 	defer row.Close()

	// 	// = db.QueryRow(
	// 	// 	fmt.Sprintf(`INSERT INTO users ("id", "username", "password") VALUES ('%s', '%s', '%s');`), id, username, password)
	// 	// if err != nil{
	// 	// 	log.Println(err)
	// 	// }
	// 	// defer row.Close()

	// 	rows, err := db.Query(`SELECT * FROM users;`)
	// 	if err != nil{
	// 		log.Println(err)
	// 	}

	// 	defer rows.Close()
	// 	for rows.Next() {
	// 		var id string
	// 		var password string
	// 		var username string
	// 		// var city string
	// 		// var salar string
	// 		err = rows.Scan(&id, &username, &password)
	// 		if err != nil{
	// 			log.Printf("Failed to retrieve row because %s", err)
	// 		}
	// 		dbs = append(dbs, username)
	// 	}

	return dbs, nil
}
*/



// check 	SELECT CoolingType FROM coolingsystem WHERE id='uuid'


// func decodeJWT (tokenString string) {

// 	publicKey := `ne-secret`

// 	decodeToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

//         if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//             return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//         }

//         return publicKey,nil
//     })
// 	if err != nil {
// 		fmt.Println("Errrors")
// 	}
// 	fmt.Println(decodeToken)

// }

// func generateJWT(username string) (string, error) {
	
// 	privKey, err := os.ReadFile("/application/grpcFunc/cert/private/7a8e6b16-dec5-4500-873a-937f0d8e0c0a")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
// 	if err != nil {
// 		log.Println("create: parse key:", err)
// 	}

// 	token := jwt.New(jwt.SigningMethodRS256)
// 	token.Header["kid"] = "7a8e6b16-dec5-4500-873a-937f0d8e0c0a"
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["username"] = username
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	claims["aud"] = "intern"

// 	tokenString, err := token.SignedString(key)


// 	if err != nil {
// 		fmt.Println("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }