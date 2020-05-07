package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	// TODO: replace println with a text template and print to readme
	// TODO: add a cli flag instead of global variable
	// The location of the activity is used to read the source code of the metadata and the implementation
	activityFolder := os.Getenv("DIR")

	// Read the metadata.go source code
	activityMetadata := MustReadFile(fmt.Sprintf("%s/metadata.go", activityFolder))
	cmap, packagename := GetCommentMap(activityMetadata)
	filteredMap := FilterCommentMapPrefix("FromMap", cmap.Comments())
	filteredMap = FilterCommentMapPrefix("ToMap", filteredMap)

	// Read the activity.go source code
	activitySource := MustReadFile(fmt.Sprintf("%s/activity.go", activityFolder))
	start, _ := GetCommentMap(activitySource)
	fmt.Printf("# %s\n\n%s\n", packagename, start.Comments()[0].Text()[8:])

	GetSettings(filteredMap)
	GetInputs(filteredMap)
	GetOutputs(filteredMap)
}

// GetCommentMap creates an ast.CommentMap from the ast.File's comments and returns both the commentmap and package name.
func GetCommentMap(filecontents string) (ast.CommentMap, *ast.Ident) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "file.go", filecontents, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.
	return ast.NewCommentMap(fset, f, f.Comments), f.Name
}

// MustReadFile reads the file named by filename and returns the contents or panics if an error is thrown.
func MustReadFile(filename string) string {
	byteArr, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(byteArr)
}

// FilterCommentMapPrefix removes all comments that start with a specific string and returns the filtered map.
func FilterCommentMapPrefix(prefix string, comments []*ast.CommentGroup) []*ast.CommentGroup {
	t := make([]*ast.CommentGroup, len(comments))
	for _, comment := range comments {
		if !strings.HasPrefix(comment.Text(), prefix) {
			t = append(t, comment)
		}
	}
	return t
}

func GetInputs(comments []*ast.CommentGroup) {
	t := make([]string, 0)

	for _, comment := range comments {
		if strings.HasPrefix(comment.Text(), "Output is") || strings.HasPrefix(comment.Text(), "Settings is") {
			break
		}
		if len(comment.Text()) > 0 && !strings.HasPrefix(comment.Text(), "Input is") {
			t = append(t, comment.Text())
		}
	}

	fmt.Println("## Inputs")

	if len(t) == 0 {
		fmt.Printf("\nNo inputs required\n\n")
		return
	}

	fmt.Printf("\n| Input | Description |\n|:---|:---|\n")

	for idx := range t {
		p := strings.Split(t[idx], "is the")
		re := regexp.MustCompile(`\r?\n`)
		p[1] = re.ReplaceAllString(p[1], "")
		fmt.Printf("| %s | %s |\n", p[0], p[1])
	}

	fmt.Println("")
}

func GetOutputs(comments []*ast.CommentGroup) {
	t := make([]string, 0)

	for _, comment := range comments {
		if strings.HasPrefix(comment.Text(), "Input is") || strings.HasPrefix(comment.Text(), "Settings is") {
			break
		}
		if len(comment.Text()) > 0 && !strings.HasPrefix(comment.Text(), "Output is") {
			t = append(t, comment.Text())
		}
	}

	fmt.Println("## Outputs")

	if len(t) == 0 {
		fmt.Printf("\nNo outputs available\n\n")
		return
	}

	fmt.Printf("\n| Output | Description |\n|:---|:---|\n")

	for idx := range t {
		p := strings.Split(t[idx], "is the")
		re := regexp.MustCompile(`\r?\n`)
		p[1] = re.ReplaceAllString(p[1], "")
		fmt.Printf("| %s | %s |\n", p[0], p[1])
	}

	fmt.Printf("\n\n")
}

func GetSettings(comments []*ast.CommentGroup) {
	t := make([]string, 0)

	for _, comment := range comments {
		if strings.HasPrefix(comment.Text(), "Input is") || strings.HasPrefix(comment.Text(), "Output is") {
			break
		}
		if len(comment.Text()) > 0 && !strings.HasPrefix(comment.Text(), "Settings is") {
			t = append(t, comment.Text())
		}
	}

	fmt.Println("## Settings")

	if len(t) == 0 {
		fmt.Printf("\nNo settings available\n\n")
		return
	}

	fmt.Printf("\n| Setting | Description |\n|:---|:---|\n")

	for idx := range t {
		p := strings.Split(t[idx], "is the")
		re := regexp.MustCompile(`\r?\n`)
		p[1] = re.ReplaceAllString(p[1], "")
		fmt.Printf("| %s | %s |\n", p[0], p[1])
	}

	fmt.Println("")
}
