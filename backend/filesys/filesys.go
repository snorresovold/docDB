package filesys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Document struct {
	Id             string `json:"id"`
	Content map[string]interface{} `json:"content"`
}

type Collection struct {
	Name      string     `json:"name"`
	Documents []Document `json:"documents"`
}

func GetCollection(collection string) Collection {
	collection = strings.ReplaceAll(collection, "+", "/")
	entries, err := os.ReadDir("./collections/" + collection)
	coll := Collection{Name: collection}
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		doc := Document{Id: e.Name()}
		content, err := ioutil.ReadFile("./collections/" + collection + "/" + e.Name())
		if err != nil {
			log.Printf("Error reading file '%s': %s", e.Name(), err)
		} else {
			err = json.Unmarshal(content, &doc.Content)
			if err != nil {
				log.Printf("Error unmarshalling JSON for file '%s': %s", e.Name(), err)
			}
		}
		coll.Documents = append(coll.Documents, doc)
	}
	return coll
}

func GetDocument(collection string, document string) Document {
	data, _ := os.ReadFile("./collections/" + collection + "/" + document + ".json")
	doc := Document{Id: document}
	_ = json.Unmarshal([]byte(data), &doc.Content)
	return doc
}

func CreateCollection(name string) {
	err := os.Mkdir(name, 0755) //create a directory and give it required permissions
	if err != nil {
		fmt.Println(err) //print the error on the console
		return
	}
}

func CreateDocument(collection string) {
	collection = strings.ReplaceAll(collection, "+", "/")
	e, err := exists(collection)
	if err != nil {
		fmt.Println(err) //print the error on the console
		return
	}
	if !e {
		CreateCollection(collection)
	}
	id := uuid.New()
	fileName := "./collections/" + collection + "/" + id.String() + ".json"
	_ = os.WriteFile(fileName, []byte(""), 0755)
}

func WriteToDocument(collection string, document string, data map[string]interface{}) {
	e, err := exists("collections/" + collection + "/" + document + ".json")
	if err != nil {
		fmt.Println(err, "error") //print the error on the console
		return
	}
	if !e {
		fmt.Println("collections/" + collection + "/" + document + ".json")
		fmt.Println("document does not exist")
	}
	//fmt.Println(data)
	for k, v := range data {
		fmt.Printf("%v %v\n", k, v)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err, "error") //print the error on the console
		return
	}
	_ = os.WriteFile("collections/" +collection+"/"+document+".json", jsonData, 0755)
}

func DeleteCollection(collection string) {
	// Using Remove() function
	e := os.Remove(collection)
	if e != nil {
		log.Fatal(e)
	}
}

func DeleteDocument(document string) {
	fmt.Println(document + ".json")
	e := os.Remove(document + ".json")
	if e != nil {
		log.Fatal(e)
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	// exists
	if err == nil {
		return true, nil
	}
	// does not exist
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}