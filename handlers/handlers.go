package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JavDomGom/sn-test/middleware"
	"github.com/JavDomGom/sn-test/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers set port, handler, listen and serve the HTTP server.*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middleware.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/message", middleware.CheckDB(middleware.ValidateJWT(routers.RecordMsg))).Methods("POST")
	router.HandleFunc("/readMsg", middleware.CheckDB(middleware.ValidateJWT(routers.ReadMsg))).Methods("GET")
	router.HandleFunc("/deleteMsg", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteMsg))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/newFollow", middleware.CheckDB(middleware.ValidateJWT(routers.NewFollow))).Methods("POST")
	router.HandleFunc("/removeFollow", middleware.CheckDB(middleware.ValidateJWT(routers.RemoveFollow))).Methods("DELETE")
	router.HandleFunc("/checkFollow", middleware.CheckDB(middleware.ValidateJWT(routers.CheckFollow))).Methods("GET")

	router.HandleFunc("/getUsers", middleware.CheckDB(middleware.ValidateJWT(routers.GetUsers))).Methods("GET")
	router.HandleFunc("/getFollowersMsg", middleware.CheckDB(middleware.ValidateJWT(routers.GetFollowersMsg))).Methods("GET")
	router.HandleFunc("/getMsg", middleware.CheckDB(middleware.ValidateJWT(routers.GetMsg))).Methods("GET")

	router.HandleFunc("/like", middleware.CheckDB(middleware.ValidateJWT(routers.NewLike))).Methods("POST")
	router.HandleFunc("/unlike", middleware.CheckDB(middleware.ValidateJWT(routers.RemoveLike))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
