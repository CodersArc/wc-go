/*
Copyright © 2024 Mithun Singh codersarc
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var countBytes, countLines, countChars, countWords, longestLine bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc-go",
	Short: "Learning project to implement unix style wc program but as a golang cli",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]
		if fileName == "" {
			return fmt.Errorf("no file provided, user -c to specify a file")
		}

		if !cmd.Flags().Changed("bytes") &&
			!cmd.Flags().Changed("lines") &&
			!cmd.Flags().Changed("chars") &&
			!cmd.Flags().Changed("words") &&
			!cmd.Flags().Changed("longest") {
			countBytes = true
			countLines = true
			countChars = true
			countWords = true
			longestLine = true
		}

		content, err := os.ReadFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to read file %s due to %s", fileName, err)
		}

		text := string(content)
		var results []string

		if countLines {
			lineCount := strings.Count(text, "\n")
			results = append(results, fmt.Sprintf("%d", lineCount))
		}

		if countWords {
			wordCount := len(strings.Fields(text))
			results = append(results, fmt.Sprintf("%d", wordCount))
		}

		if countBytes {
			results = append(results, fmt.Sprintf("%d", len(content)))
		}

		if countChars {
			results = append(results, fmt.Sprintf("%d", len([]rune(text))))
		}

		if longestLine {
			longest := findLongestLine(fileName)
			results = append(results, fmt.Sprintf("Longest Line (bytes): %d", longest))
		}

		fmt.Printf("%s %s\n", strings.Join(results, " "), fileName)

		return nil
	},
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {

	// Add the `-c` flag for counting bytes
	rootCmd.Flags().BoolVarP(&countBytes, "bytes", "c", false, "Count bytes in the specified file.")

	// Add the `-l` flag for counting lines
	rootCmd.Flags().BoolVarP(&countLines, "lines", "l", false, "Count lines in the specified file.")

	// Add the `-m` flag for counting characters
	rootCmd.Flags().BoolVarP(&countChars, "chars", "m", false, "Count characters in the specified file.")

	// Add the `-w` flag for counting words
	rootCmd.Flags().BoolVarP(&countWords, "words", "w", false, "Count words in the specified file.")

	// Add the `-L` flag for finding the longest line
	rootCmd.Flags().BoolVarP(&longestLine, "longest", "L", false, "Find the longest line in the specified file.")
}

// Helper function to find the length of the longest line
func findLongestLine(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxLength := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	return maxLength
}
