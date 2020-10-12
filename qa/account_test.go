// +build qa

package qa

import (
	"fmt"
	"os"
	"testing"
	"time"
	"context"

	"carlware/accounts/cli/config"
	"carlware/accounts/cli/dispatchers"
	"carlware/accounts/cli/dispatchers/graphql"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "progress",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	// init graphql server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	config.InitConfig("")
	ctrl := dispatchers.NewController(config.Conf)
	// TODO: wait until is running
	go graphql.NewGraphQL(ctx, config.Conf, ctrl)

	// test
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func iSendRequestTo(arg1, arg2 string) error {
	return fmt.Errorf("err")
}

func theResponseCodeShouldBe(arg1 int) error {
	fmt.Println(arg1)
	return nil
}

func theReponseShouldMatchJson(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, theResponseCodeShouldBe)
	s.Step(`^The reponse should match json:$`, theReponseShouldMatchJson)
}
