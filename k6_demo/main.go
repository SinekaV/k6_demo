// package main

// import (
//     "fmt"
//     "os"
//     "os/exec"
// )

// func main() {
//     // Define the k6 script as a string
//     k6Script := `
//         import http from 'k6/http';

//         export default function () {
//             const payload = JSON.stringify({
//                 name: 'lorem',
//                 surname: 'ipsum',
//             });
//             const headers = { 'Content-Type': 'application/json' };
//             http.post('https://httpbin.test.k6.io/post', payload, { headers });
//         }
//     `

//     // Create a temporary file to store the k6 script
//     tmpfile, err := os.CreateTemp("", "k6script-*.js")
//     if err != nil {
//         fmt.Println("Error creating temporary file:", err)
//         return
//     }
//     defer tmpfile.Close()
//     defer os.Remove(tmpfile.Name())

//     // Write the k6 script to the temporary file
//     _, err = tmpfile.WriteString(k6Script)
//     if err != nil {
//         fmt.Println("Error writing k6 script to temporary file:", err)
//         return
//     }

//     // Run k6 as a subprocess
//     cmd := exec.Command("k6", "run", tmpfile.Name())
//     cmd.Stdout = os.Stdout
//     cmd.Stderr = os.Stderr

//	    err = cmd.Run()
//	    if err != nil {
//	        fmt.Println("Error running k6:", err)
//	        return
//	    }
//	}
package main

import (
	"fmt"
	"log"

	"context"
	"k6_demo/config"
	"k6_demo/constants"
	"k6_demo/controllers"
	"k6_demo/routes"
	"k6_demo/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	profileCollection := mongoClient.Database(constants.Dbname).Collection("customer")
	profileService := service.InitCustomer(profileCollection, ctx)
	profileController := controllers.InitTransController(profileService)
	routes.CustRoute(server, profileController)
}

// var secretKey = []byte("https://www.postman.com/descent-module-technologist-70995397/workspace/jwt") // Replace with your secret key

// func createToken() (string, error) {
//     claims := jwt.MapClaims{
//         "username": "exampleuser",
//         "exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     return token.SignedString(secretKey)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//     tokenString, err := createToken()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     fmt.Fprintf(w, "JWT Token: %s", tokenString)
// }

//	func main() {
//	    http.HandleFunc("/", handler)
//	    http.ListenAndServe(":4000", nil)
//	}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
