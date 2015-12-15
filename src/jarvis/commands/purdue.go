// Commands specific to purdue university.
// Other users of jarvis would probably want to disable these commands by removing
// the initialization of the Purdue{} object from handlers/commands.go
package commands

import (
	"fmt"
	"jarvis/service"
	"jarvis/util"
	"jarvis/ws"
	"strings"
	"time"
)

type Purdue struct{}

func NewPurdue() Purdue {
	return Purdue{}
}

func (p Purdue) Name() string {
	return "purdue"
}

func (p Purdue) Description() string {
	return "commands specific to purdue university."
}

func (p Purdue) Examples() []string {
	return []string{"jarvis lunch at earhart today"}
}

func (p Purdue) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{
		util.HelpTopic{
			Name: "dining court locations",
			Body: "earhart, ford, hillenbrand, wiley, windsor",
		},
		util.HelpTopic{
			Name: "dining court meals",
			Body: "breakfast, lunch, dinner",
		},
	}
}

func (p Purdue) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis (breakfast|lunch|dinner) at (?P<location>[^ ]+) (today|tomorrow)$", p.DiningMenu),
		util.NewSubCommand("^jarvis (breakfast|lunch|dinner) at (?P<location>[^ ]+)$", p.DiningMenuNoDay),
	}
}

func (p Purdue) DiningMenuNoDay(m util.IncomingSlackMessage, r util.Regex) {
	meal := strings.ToLower(r.SubExpression(m.Text, 0))
	location := strings.ToLower(r.SubExpression(m.Text, 1))
	day := time.Now()
	p.SendMenus(meal, location, day, m)
}

func (p Purdue) DiningMenu(m util.IncomingSlackMessage, r util.Regex) {
	meal := strings.ToLower(r.SubExpression(m.Text, 0))
	location := strings.ToLower(r.SubExpression(m.Text, 1))
	var day time.Time
	if r.SubExpression(m.Text, 2) == "today" {
		day = time.Now()
	} else {
		day = time.Now().Add(24 * time.Hour)
	}
	p.SendMenus(meal, location, day, m)
}

func (p Purdue) SendMenus(meal string, location string, day time.Time, m util.IncomingSlackMessage) {
	meal = strings.Title(meal)
	if location != "earhart" && location != "ford" && location != "hillenbrand" && location != "wiley" && location != "windsor" {
		ws.SendMessage("The location you provided doesn't appear to be an actual dining court.", m.Channel)
		return
	}
	menus, err := service.Purdue{}.GetDiningMenu(location, day)
	if err != nil {
		ws.SendMessage(err.Error(), m.Channel)
		return
	}
	var mealObj service.DiningMeal
	var found bool
	for _, actMeal := range menus.Meals {
		if actMeal.Name == meal {
			mealObj = actMeal
			found = true
		}
	}
	if !found {
		ws.SendMessage("The meal you requested doesn't appear to be served on that day.", m.Channel)
		return
	}
	if mealObj.Status == "Closed" {
		ws.SendMessage("That dining court is not serving any meals at that time.", m.Channel)
		return
	}
	response := fmt.Sprintf("%v is serving %v from %v to %v\n", strings.Title(location), mealObj.Name, mealObj.Hours.StartTime, mealObj.Hours.EndTime)
	for _, station := range mealObj.Stations {
		response += fmt.Sprintf("*%v*\n", station.Name)
		for _, item := range station.Items {
			response += fmt.Sprintf("> _%v_", item.Name)
			if item.IsVegetarian {
				response += " :herb:"
			}
			for _, allergen := range item.Allergens {
				if allergen.Value {
					switch allergen.Name {
					case "Eggs":
						response += " :egg:"
					case "Gluten":
						response += " :bread:"
					case "Fish":
					case "Shellfish":
						response += " :fish:"
					case "Milk":
						response += " :cow:"
					case "Peanuts":
						response += " :chestnut:"
					case "Soy":
						response += " :rice:"
					case "Tree Nuts":
						response += " :evergreen_tree:"
					case "Wheat":
						response += " :ear_of_rice:"
					}
				}
			}
			response += "\n"
		}
	}
	ws.SendMessage(response, m.Channel)
}
