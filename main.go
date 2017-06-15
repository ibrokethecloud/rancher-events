package main

import "github/ibrokethecloud/rancher-events/events"
import log "github.com/Sirupsen/logrus"
func init() {
  log.SetFormatter(&log.JSONFormatter{})
}
func main() {
  events.GetContainerEvents()

}
