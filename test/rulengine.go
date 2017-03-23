package main

import (
	"log"

	"github.com/going/rulengine"
	"github.com/going/rulengine/facts"
	"github.com/going/rulengine/logic"
)

func main() {
	engine := rulengine.NewRuleEngine()
	engine.AddExpression("$results.impression.count > 0", "age_larger_than_35")
	engine.AddExpression("$results._id.ad.id == 5785135de2fc1cc73e8b456b", "male")
	engine.AddRule(&logic.Rule{Expression: "age_larger_than_35 & male", Action: "Pass"})

	data := facts.NewFactCollection()
	user := `
		{
		  "_id" : {
		    "imp_hour" : {
		      "y" : 2017,
		      "m" : 3,
		      "d" : 2,
		      "h" : 15
		    },
		    "device" : {
		      "os" : {
			"name" : "android",
			"version" : "6.0.1"
		      },
		      "model" : "SM-N915F",
		      "geo" : {
			"country" : "ID"
		      },
		      "network" : "wifi"
		    },
		    "slot" : {
		      "site" : {
			"id" : "kdMN29Re",
			"aff_id" : "26781",
			"publisher" : {
			  "id" : "583516eee2fc1cb5018b4567"
			}
		      },
		      "tag" : ""
		    },
		    "ad" : {
		      "id" : "5785135de2fc1cc73e8b456b",
		      "adgroup" : {
			"id" : "577f82ffe2fc1c38758b456c"
		      },
		      "campaign" : {
			"id" : "577f82aee2fc1cba698b4569"
		      },
		      "advertiser" : {
			"id" : "569f6818e2fc1cf15d8b4567"
		      }
		    }
		  },
		  "impression" : {
		    "count" : 4
		  }
		}
	`
	fact, err := facts.NewFact(user)
	if err != nil {
		log.Fatal(err)
	}
	data.Add("results", fact)

	exprNames := engine.GetFiredExpressions(data)
	log.Println(exprNames)

	actions := engine.GetAction(data)
	log.Printf("%#v", actions)
}
