package main

type Keyword struct{
	Keyword : string
	Answers : []string
}


var welcome=[]string{"What's bothering you?",
                      "Talk to me.",
					  "How are you today?"}


var keywords=map[string]Keyword{
	"sorry":Keyword{
		Keyword: "sorry",
		Answerd: []string{
			"Don't have to be sorry."
			"Why are you sorry?"
			"Get yourself together, don't say sorry"
		}	
	},

	"xnone":Keyword{
		Keyword: "xnone",
		Answerd: []string{
			"I dont understand you",
			"Go on",
			"Tell me more about this"
		}
	}
}
