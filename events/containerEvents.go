package events

import "github.com/rancher/go-rancher/client"
import "os"
import "time"
import log "github.com/Sirupsen/logrus"
import "encoding/json"
import "github.com/fsouza/go-dockerclient"


func GetContainerEvents(){
  // Function to query container events //
  /* Module to try and use the rancher-client for api data */
  clientConnection := client.ClientOpts {Url:  os.Getenv("CATTLE_URL"), AccessKey: os.Getenv("CATTLE_ACCESS_KEY"), SecretKey: os.Getenv("CATTLE_SECRET_KEY")}


  /* Setting my custom ListOpts Filter */
  myfilter := make(map[string]interface{})
  myfilter["sort"] = "id"
  myfilter["limit"] = 100
  myfilter["order"] = "desc"
  listOptions := client.ListOpts{Filters: myfilter}

  rancherClient,err := client.NewRancherClient(&clientConnection)
  checkError(err)

  containerEventSummary,err := rancherClient.ContainerEvent.List(&listOptions)
  checkError(err)

  // Need to iterate over summary and get individual events //

  var containerEvents = []client.ContainerEvent {}
  var event = client.ContainerEvent {}
  containerEvents = containerEventSummary.Data

  for _,event = range containerEvents {
    parseEvent(event.DockerInspect)
  }
}

func checkError(err error) {
    if err != nil {
        log.Errorf("Fatal Error: %v",err)
        os.Exit(1)
    }
}

func parseEvent(v interface{}) {
  var dockerInspectInfo = docker.Container {}
  jsonInspectInfo, err := json.Marshal(v)
  checkError(err)
  err = json.Unmarshal(jsonInspectInfo, &dockerInspectInfo)
  checkError(err)
  // Call to compareTime //
  if (compareTime(dockerInspectInfo.State.FinishedAt)){
    log.Info(string(jsonInspectInfo))
    notifyWebHook()
  }
}

func compareTime(endTime time.Time) (result bool){
  currentTime := time.Now()
  timeDuration := currentTime.Sub(endTime)
  if (timeDuration > 10*time.Minute ) {
    result = false
  } else {
    result = true
  }
  return result
}

func notifyWebHook() {
  // Define a web hook function here //

}
