package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"jackliu-tts-service/config"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// os install: pip3 install edge-tts
var (
	serverURI      string
	localPort      int
	audioPath      string
	mediaFilesRoot string
	validTokenStr  string
)

func init() {
	configOptions, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	serverURI = configOptions.ServerURI
	localPort = configOptions.LocalPort
	audioPath = configOptions.AudioPath
	mediaFilesRoot = configOptions.MediaFilesRoot
	validTokenStr = configOptions.ValidTokenStr
}

func isValidRole(role string) bool {
	roleList := []string{
		"zh-CN-XiaoxiaoNeural",
		"zh-CN-XiaoyiNeural",
		"zh-CN-YunjianNeural",
		"zh-CN-YunxiNeural",
		"zh-CN-YunxiaNeural",
		"zh-CN-YunyangNeural",
		"zh-CN-liaoning-XiaobeiNeural",
		"zh-CN-shaanxi-XiaoniNeural",
		"zh-HK-HiuGaaiNeural",
		"zh-HK-HiuMaanNeural",
		"zh-HK-WanLungNeural",
		"zh-TW-HsiaoChenNeural",
		"zh-TW-HsiaoYuNeural",
		"zh-TW-YunJheNeural",
	}
	exist := false

	for _, roleItem := range roleList {
		if roleItem == role {
			exist = true
			break
		}
	}
	return exist
}

func generateSpeechText(text string, role string, rate string) (string, error) {
	// check role
	if role == "" {
		role = "zh-CN-XiaoyiNeural"
	} else {
		if !isValidRole(role) {
			return "Role invalidation", nil
		}
	}
	// get md5(filename) value
	fileName := fmt.Sprintf("%x", md5.Sum([]byte(text+"_"+role+"_"+rate))) + ".mp3"

	audioUrl := audioPath + fileName
	// 输出json
	result := "{ \"audio_url\": \"" + audioUrl + "\" }"

	// eheck and respose media file url
	if _, err := os.Stat(mediaFilesRoot + audioPath + fileName); err == nil {
		fmt.Println("File already exists, return directly.")
		return result, nil
	}

	// check cmd
	_, err := exec.LookPath("edge-tts")
	if err != nil {
		fmt.Println("edge-tts not found. Please make sure to follow the installation command of python3 or higher: pip3 install edge_tts.")
		return "Command not found !", nil
	}

	cmd := exec.Command("edge-tts", "--rate", rate, "--voice", role, "--text", text, "--write-media", mediaFilesRoot+audioPath+fileName)
	// cmd output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Command execution error !:", err, output)
		return "Command execution error !", err
	}

	return result, nil
}

func speechHandler(w http.ResponseWriter, r *http.Request) {
	// get form values
	text := r.FormValue("text")
	access_token := r.FormValue("access_token")
	role := r.FormValue("role")
	rate := r.FormValue("rate")
	isExpired := r.FormValue("isExpired")
	if isExpired == "" {
		isExpired = "True"
	}

	if len(text) == 0 || len(access_token) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing text or access_token parameter")
		return
	}

	if access_token != calculateMD5([]byte(validTokenStr)) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing text or access_token parameter")
		return
	}

	// generate audio file
	speechURL, err := generateSpeechText(text, role, rate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to generate speech: %v", err)
		return
	}

	fmt.Fprint(w, speechURL)
}

// generate  MD5 value
func calculateMD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashValue)
	return hashHex
}

// isValidOrigin
func isValidOrigin(origin string) bool {
	allowedOrigins := []string{
		"*",
		fmt.Sprintf("localhost:%d", localPort),
		fmt.Sprintf("127.0.0.1:%d", localPort),
		fmt.Sprintf("%s:%d", serverURI, localPort),
	}

	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}

	return false
}

// cors by middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check origin
		origin := r.Header.Get("Origin")

		//if !isValidOrigin(origin) {
		//	fmt.Println("403 Forbidden")
		//	return
		//}

		// add CORS
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// check method OPTIONS
		if r.Method == "OPTIONS" {
			// OPTIONS
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	// create router
	router := http.NewServeMux()

	// register router
	router.HandleFunc("/tts", speechHandler)

	httpServerUri := serverURI + ":" + strconv.Itoa(localPort)

	fmt.Println(httpServerUri)

	// middleware
	http.ListenAndServe(httpServerUri, corsMiddleware(router))

}
