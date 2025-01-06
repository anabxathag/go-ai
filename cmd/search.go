/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	//"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		// Pass the apiKey from arguments to getResponse
		apiKey, _ := cmd.Flags().GetString("api_key")
		getResponse(apiKey, args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Adding flag for GEMINI_API_KEY
	searchCmd.Flags().StringP("api_key", "k", "", "Gemini API Key")
}

// Function can now accept slice parameter
func getResponse(apiKey string, query []string) {
	// Creating a sentence out of a slice
	userArgs := strings.Join(query[0:], " ")

	// Load .env file
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	// Get the API key
	//apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	}

	ctx := context.Background()
	//client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	// change the hardcoded text to userArgs variable
	resp, err := model.GenerateContent(ctx, genai.Text(userArgs))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Candidates[0].Content.Parts[0])
}
