package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "options",
			Description: "Options commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "Family Name",
					Description: "Your family name in Black Desert Online",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))
			msgFormat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"

			if option, ok := optionMap["string-option"]; ok {
				margs = append(margs, option.StringValue())
				msgFormat += "> string-option: %s\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgFormat,
						margs...,
					),
				},
			})
		},
	}
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	authToken := os.Getenv("AUTH_TOKEN")
	discordgo.New(authToken)
}
