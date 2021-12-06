/*
Copyright © 2021 Renato Torres <renato.torres@pm.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package internet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

type mediumQuery struct {
	Query string `json:"query"`
}

type mediumPostResponse struct {
	Data mediumData `json:"data"`
}

type mediumData struct {
	Post mediumPost `json:"post"`
}

type mediumPost struct {
	Title   string            `json:"title"`
	Creator mediumPostCreator `json:"creator"`
	Content mediumPostContent `json:"content"`
}

type mediumPostCreator struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type mediumPostContent struct {
	BodyModel mediumPostContentBodyModel `json:"bodyModel"`
}

type mediumPostContentBodyModel struct {
	Paragraphs []mediumPostParagraph `json:"paragraphs"`
}

type mediumPostParagraph struct {
	Text     string                      `json:"text"`
	Type     string                      `json:"type"`
	HRef     string                      `json:"href"`
	IFrame   string                      `json:"iframe"`
	Layout   string                      `json:"layout"`
	Markups  []mediumPostParagraphMarkup `json:"markups"`
	Metadata mediumPostParagraphMetadata `json:"metadata"`
}

type mediumPostParagraphMarkup struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	HRef       string `json:"href"`
	Start      int    `json:"start"`
	End        int    `json:"end"`
	Rel        string `json:"rel"`
	AnchorType string `json:"anchorType"`
}

type mediumPostParagraphMetadata struct {
	TypeName       string `json:"__typename"`
	Id             string `json:"id"`
	OriginalWidth  int    `json:"originalWidth"`
	OriginalHeight int    `json:"originalHeight"`
}

type mediumToMdOutput struct {
	Markdown string `json:"markdown"`
	PostId   string `json:"postId"`
}

const flagId = "post-id"
const flagMdToFile = "md-to-file"
const flagJsonToFile = "json-to-file"

func NewMediumToMdCmd(iostreams iostreams.IOStreams) *cobra.Command {
	var mediumToMdCmd = &cobra.Command{
		Use:   "medium2md",
		Short: "Converts a medium post to markdown",
		Long: heredoc.Doc(`
			Converts a Medium post to markdown!

			Why would I want to use this? There are a number of potential reasons:
			- You believe in an open web (http://scripting.com/liveblog/users/davewiner/2016/01/20/0900.html)
			- You believe more in the author than the platform (https://www.manton.org/2016/01/15/silos-as-shortcuts.html)
			- You don't like the reading experience that Medium provides (https://twitter.com/BretFisher/status/1206766086961745920)
			- You object to Medium's extortionist business tactics (https://www.cdevn.com/why-medium-actually-sucks/)
			- You're concerned about how Medium uses your data (https://tosdr.org/en/service/1003)
			- Other reasons (https://nomedium.dev)

			Inspiration: https://scribe.rip/faq
		`),
		Example: heredoc.Doc(`
			canivete internet medium2md -i 4b63ff0e2bd3
			canivete internet medium2md -i https://medium.com/@bradleyalanlaplante/heres-everything-i-do-when-setting-up-a-new-computer-4b63ff0e2bd3
			canivete internet medium2md -i 4b63ff0e2bd3 -f -d`),
		RunE: func(cmd *cobra.Command, args []string) error {

			postId, _ := cmd.Flags().GetString(flagId)
			outputMdToFile, _ := cmd.Flags().GetBool(flagMdToFile)
			outputJsonToFile, _ := cmd.Flags().GetBool(flagJsonToFile)

			output, err := run(postId, outputMdToFile, outputJsonToFile)
			if err != nil {
				return err
			}

			err = iostreams.PrintOutput(output)
			return err

		},
	}

	mediumToMdCmd.Flags().StringP(
		flagId,
		"i",
		"",
		heredoc.Doc(`the identifier of the post (e.g 4b63ff0e2bd3)
			you can find this in the last part of a Medium URL or
			use directly the Medium URL - check the examples`))
	mediumToMdCmd.MarkFlagRequired(flagId)

	mediumToMdCmd.Flags().BoolP(
		flagMdToFile,
		"f",
		false,
		"writes the markdown to a file named <post-id>.md")

	mediumToMdCmd.Flags().BoolP(
		flagJsonToFile,
		"d",
		false,
		"writes the raw JSON fetched from Medium to a file named <post-id>.json")

	return mediumToMdCmd
}

func run(postId string, outputMdToFile bool, outputJsonToFile bool) (mediumToMdOutput, error) {

	output := mediumToMdOutput{}

	sanitizedPostId := sanitizePostId(postId)
	result := getPostData(sanitizedPostId, outputJsonToFile)
	output.Markdown = postToMarkdown(result)
	output.PostId = sanitizedPostId

	if outputMdToFile {
		err := ioutil.WriteFile(fmt.Sprintf("%s.md", sanitizedPostId), []byte(output.Markdown), fs.ModePerm)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

// Sanitizes the post id
// From the following URL extracts "4b63ff0e2bd3"
// https://medium.com/@bradleyalanlaplante/heres-everything-i-do-when-setting-up-a-new-computer-4b63ff0e2bd3?source=email-bea602c1f3c7-1638752765617-digest.reader--4b63ff0e2bd3----1-72------------------462baf81_c223_437a_8412_a2811df1b2fe-28-
func sanitizePostId(postId string) string {
	if !strings.Contains(postId, "https") {
		return postId
	}

	sanitizePostId := postId
	if strings.Contains(sanitizePostId, "?") {
		index := strings.Index(sanitizePostId, "?")
		sanitizePostId = sanitizePostId[0:index]
	}

	splits := strings.Split(sanitizePostId, "-")
	sanitizePostId = splits[len(splits)-1]

	return sanitizePostId
}

func getPostData(postId string, outputJsonToFile bool) mediumPostResponse {

	post := mediumPostResponse{}

	url := "https://medium.com/_/graphql"

	query := fmt.Sprintf(
		`
		query {
			post(id: "%s") {
			  title
			  createdAt
			  creator {
				id
				name
			  }
			  content {
				bodyModel {
				  paragraphs {
					text
					type
					href
					layout
					markups {
					  title
					  type
					  href
					  userId
					  start
					  end
					  anchorType
					}
					iframe {
					  mediaResource {
						href
						iframeSrc
						iframeWidth
						iframeHeight
					  }
					}
					metadata {
					  id
					  originalWidth
					  originalHeight
					}
				  }
				}
			  }
			}
		  }
		`,
		postId)

	queryStruct := mediumQuery{Query: query}

	data, err := json.Marshal(queryStruct)
	if err != nil {
		log.Fatal("Error marshalling request.", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error reading request.", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	// Convert response to the struct
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal("Error un-marshalling medium response.", err)
	}

	// Writes json to a file
	if outputJsonToFile {
		err := ioutil.WriteFile(fmt.Sprintf("%s.json", postId), body, fs.ModePerm)
		if err != nil {
			log.Fatal("Error writting json file.", err)
		}
	}
	return post
}

func postToMarkdown(post mediumPostResponse) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("# %s\n", post.Data.Post.Title))
	buffer.WriteString(fmt.Sprintf("By %s\n", post.Data.Post.Creator.Name))

	for _, paragraph := range post.Data.Post.Content.BodyModel.Paragraphs {
		if paragraph.Type == "H3" {
			buffer.WriteString(fmt.Sprintf("\n## %s\n", paragraph.Text))
		} else if paragraph.Type == "H4" {
			buffer.WriteString(fmt.Sprintf("\n### _%s_\n", paragraph.Text))
		} else if paragraph.Type == "P" {
			buffer.WriteString(fmt.Sprintf("\n%s\n", paragraph.Text))
		} else if paragraph.Type == "IMG" {
			buffer.WriteString(fmt.Sprintf("\n![%s](https://miro.medium.com/max/1400/%s)\n", paragraph.Text, paragraph.Metadata.Id))

		}
		if len(paragraph.Markups) > 0 {
			textParts := []string{}
			lastStartIndex := 0
			for _, markup := range paragraph.Markups {
				if markup.Type != "A" {
					continue
				}
				textParts = append(textParts, paragraph.Text[lastStartIndex:markup.Start])
				textParts = append(textParts, fmt.Sprintf("[%s](%s)",
					paragraph.Text[markup.Start:markup.End],
					markup.HRef))
				lastStartIndex = markup.End
			}
			buffer.WriteString(fmt.Sprintf("\n%s\n", strings.Join(textParts, "")))
		}
	}

	return buffer.String()
}
