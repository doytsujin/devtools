package githubstats

import (
	"math"
	"sort"
	"time"

	"github.com/grafana/devtools/pkg/ghevents"
)

var skipRepos = []string{}
var mapRepos = map[string]string{}

func filterAndPatchRepos(msg interface{}) bool {
	evt := msg.(*ghevents.Event)

	for _, repo := range skipRepos {
		if repo == evt.Repo.Name {
			return false
		}
	}

	if mapToRepo, ok := mapRepos[evt.Repo.Name]; ok {
		evt.Repo.Name = mapToRepo
	}

	return true
}

func fromCreatedDate(msg interface{}) time.Time {
	evt := msg.(*ghevents.Event)
	return evt.CreatedAt
}

func filterByOpenedAndClosedActions(msg interface{}) bool {
	evt := msg.(*ghevents.Event)
	return *evt.Payload.Action == "opened" || *evt.Payload.Action == "closed"
}

func partitionByRepo(msg interface{}) (string, interface{}) {
	evt := msg.(*ghevents.Event)
	return "repo", evt.Repo.Name
}

var bots = []string{
	"CLAassistant",
	"codecov-io",
}

func isBot(login string) bool {
	for _, bot := range bots {
		if bot == login {
			return true
		}
	}

	return false
}

var userLoginGroupMap = map[string]string{}

func mapUserLoginToGroup(login string) string {
	contributorGroup := "Contributor"

	if group, ok := userLoginGroupMap[login]; ok {
		contributorGroup = group
	}

	return contributorGroup
}

func percentile(k float64, values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sort.Float64s(values)
	index := k * float64(len(values))

	if index != float64(int64(index)) {
		index = math.Round(index)
		if int(index) == 0 {
			return values[int(index)]
		}
		return values[int(index)-1]
	}

	slice := values[int(index)-1 : int(index)+1]
	return (slice[0] + slice[1]) / 2
}
