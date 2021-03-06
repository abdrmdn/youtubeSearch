package main

import (
	"fmt"
	"log"
	"flag"
	"time"
	"net/http"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "io/ioutil"
import "os"
import "github.com/youtube-videos/go-youtube-dl"
import "app/ytvideo"

var (
	query = flag.String("query", "iprice Mannequinchallenge", "")
	listingVideos = flag.String("chart", "mostPopular", "")
	maxResults = flag.Int64("max-results", 50, "Max YouTube results")
	db sql.DB
	debugOutput = true
	number = 0
)
// var db sql.DB
const developerKey = "AIzaSyCq6GaikitWw3X3xMduprZB_soUZqvg9_c"

//////////////////
//START OF Miscilanious 
func handleError(err error) {
	if err != nil {
		log.Println(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}
//END OF Miscilanious 
//////////////////

//////////////////
//START OF YOUTUBE
func listTrending(c chan ytvideo.YTVideo, tPoolNum chan int, videoCategoryId string) {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Videos.List("id,snippet,contentDetails").
		Chart("mostPopular").
		VideoCategoryId(videoCategoryId).
		MaxResults(*maxResults)

	response, err := call.Do()
	if err != nil {
		log.Printf("Error making search API call: %v for category %v", err, videoCategoryId)

		return ;
	}

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		vid := ytvideo.YTVideo{}
		videoObj := vid.ConvertVideoResult(item)
		c <- videoObj
	}
}
//Bring suggestions for the videos Id passed
func getVideoSuggestions(videoId string, videoChan chan ytvideo.YTVideo, pageToken string, tPoolNum chan int) {
	//([]*youtube.Video){
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		MaxResults(*maxResults).
		Type("video").
		RelatedToVideoId(videoId).
		PageToken(pageToken)

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Group video, channel, and playlist results in separate lists.
	//can use type youtube.Video later
	videos := make(map[string]ytvideo.YTVideo)


	// // Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		videoObj := ytvideo.YTVideo{}.ConvertSearchResult(item)
		videos[item.Id.VideoId] = videoObj
		go videosHandler(videoChan, tPoolNum)
		videoChan <- videoObj
	}


	//fetch the rest of th videos if we still have
	if len(response.NextPageToken) > 0 {
		getVideoSuggestions(videoId, videoChan, response.NextPageToken, tPoolNum)
	}
}
//END OF YOUTUBE
//////////////////

//Printers
func videosHandler(videoChan chan ytvideo.YTVideo, tPoolNum chan int) {
	video := <-videoChan
	<-tPoolNum // get a turn in the pool
	defer consumeThread(tPoolNum) // to give turn to other threads
	if debugOutput {
		fmt.Println(video.Id)
	}
	ytdl := youtube_dl.YoutubeDl{}
	ytdl.Path = "$GOPATH/src/app/srts"
	err := ytdl.DownloadVideo(video.Id)
	if err != nil {
		log.Printf("%v", err)
	}
	//if debugOutput {
	//	log.Printf("command : %v", command)
	//}
	fmt.Print(".");
	StoreValue(video)
	getVideoSuggestions(video.Id, videoChan, "12", tPoolNum)// 12 is a random token that works as initial value
}

// START OF MYSQL
//////////////////

// Initialize Mysql Connection
func initializeMysqlConn() {
	dbConn, err := sql.Open("mysql", "admin:admin@tcp(y2search_mysql:3306)/y2search_db?collation=utf8mb4_unicode_ci")
	db = *dbConn
	if err != nil {
		log.Panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Panic(err.Error()) // proper error handling instead of panic in your app
	}
}

// destruct Mysql Connection
func tearDownMysqlConn() {
	db.Close()
}

// store values in mysql connection
func StoreValue(ytVideoObj ytvideo.YTVideo) {
	// Prepare statement for inserting data
	// INSERTING VIDEO
	// Prepairing
	videoInsertQuery := "INSERT INTO videos (`id`, `video_hash_id`,`video_url`,`video_title`) VALUES (NULL, ?, ?, ?)"
	stmtVidIns, err := db.Prepare(videoInsertQuery) // ? = placeholder
	handleError(err)
	defer stmtVidIns.Close() // Close the statement when we leave main() / the program terminates
	result, err := stmtVidIns.Exec(ytVideoObj.Id, ytVideoObj.Url, ytVideoObj.Title)// Inserting
	handleError(err)
	lastInsertedId, _ := result.LastInsertId()

	// Read subtitles file
	file, err := ioutil.ReadFile("srts/" + ytVideoObj.Id + ".en.vtt")// it will be save with this extension regardless
	os.Remove("srts/" + ytVideoObj.Id + ".en.vtt");
	if debugOutput {
		fmt.Println(err)
	}
	if err == nil {
		// INSERTING VIDEO's Subtitles
		// Prepairing
		stmtVidSubIns, err := db.Prepare("INSERT INTO videos_subtitles (`id`, `video_id`,`subtitles`,`language`) VALUES (NULL, ?, ?, ?)") // ? = placeholder
		handleError(err)
		defer stmtVidSubIns.Close() // Close the statement when we leave main() / the program terminates
		// Inserting
		_, err = stmtVidSubIns.Exec(lastInsertedId, file, `en`) // Insert tuples
		handleError(err)

		// INSERTING VIDEO's Meta
		// Prepairing
		stmtVidMetaIns, err := db.Prepare("INSERT INTO videos_meta (`id`, `video_id`,`image_default`,`image_medium`,`image_high`) VALUES (NULL, ?, ?, ?, ?)") // ? = placeholder
		handleError(err)
		defer stmtVidMetaIns.Close() // Close the statement when we leave main() / the program terminates
		// Inserting
		_, err = stmtVidMetaIns.Exec(lastInsertedId, ytVideoObj.ThumbnailDefault, ytVideoObj.ThumbnailMedium, ytVideoObj.ThumbnailHigh) // Insert tuples
		handleError(err)
	}
}

//////////////////
// END OF MYSQL

func threadsPoolManager(tPoolNum chan int) {
	// So that we run only 10 at a time
	for counter := 0; counter < 100; counter++ {
		tPoolNum <- counter
	}
}

// After each thread consume it self, it will call this so it can give one call to another thread.
func consumeThread(tPoolNum chan int) {
	tPoolNum <- 1
}

////////////////////////////////////
////////////////////////////////////
// MAIN APPLICATION START POINT

func main() {
	LogfileName := "/tmp/logs/"+"log_" + time.Now().Format("2006-01-02_15:04:05") + ".log"
	f, err := os.OpenFile(LogfileName, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	// don't forget to close it
	defer f.Close()
	// assign it to the standard logger
	log.SetOutput(f)

	//mysql connection
	initializeMysqlConn()
	var c chan ytvideo.YTVideo = make(chan ytvideo.YTVideo)
	var tPoolNum chan int = make(chan int)
	videoCategoryIds := os.Args[1:]
	for _, catId := range videoCategoryIds {
		go listTrending(c, tPoolNum, catId)
	}
	go threadsPoolManager(tPoolNum)

	for i := 0; i < 5000; i++ {
		go videosHandler(c, tPoolNum)
	}

	var input string
	fmt.Scanln(&input)
	defer tearDownMysqlConn()
}
// MAIN APPLICATION END POINT
////////////////////////////////////
////////////////////////////////////
