package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antchfx/xmlquery"
)

var removeTitleTestInput = `
<?xml version="1.0" encoding="UTF-8"?>
<Staff id="1">
  <VBox>
    <height>10</height>
    <eid>4294967419</eid>
    <linkedMain/>
    <Text>
      <eid>8589934598</eid>
      <linkedMain/>
      <style>title</style>
      <text>Blessed Assurance</text>
      </Text>
    <Text>
      <eid>12884901894</eid>
      <linkedMain/>
      <style>subtitle</style>
      <text>Subtitle</text>
      </Text>
    <Text>
      <eid>17179869190</eid>
      <linkedMain/>
      <style>composer</style>
      <text>Mrs. J. F. Knapp</text>
      </Text>
    <Text>
      <eid>21474836486</eid>
      <linkedMain/>
      <style>poet</style>
      <text>Franny J. Crosby</text>
      </Text>
    </VBox>
</Staff>
`

var removeTitleTestExpectedOutput = `<?xml version="1.0" encoding="UTF-8"?><Staff id="1"><VBox><height>10</height><eid>4294967419</eid><linkedMain></linkedMain><Text><eid>12884901894</eid><linkedMain></linkedMain><style>subtitle</style><text>Subtitle</text></Text><Text><eid>17179869190</eid><linkedMain></linkedMain><style>composer</style><text>Mrs. J. F. Knapp</text></Text><Text><eid>21474836486</eid><linkedMain></linkedMain><style>poet</style><text>Franny J. Crosby</text></Text></VBox></Staff>`

func TestRemoveTitle(t *testing.T) {
	newData, err := removeTextWithTitleStyle([]byte(removeTitleTestInput))
	if err != nil {
		t.Errorf("musescoreRemoveTitle failed with error: %v", err)
	}

	if string(newData) != removeTitleTestExpectedOutput {
		t.Errorf("musescoreRemoveTitle returned wrong output")
		fmt.Println("EXP:", removeTitleTestExpectedOutput)
		fmt.Println("GOT:", string(newData))
	}
}

const xmlDoc = `
    <?xml version="1.0"?>
<catalog>
   <!-- book list-->
   <book id="bk101">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>
   </book>
   <book id="bk102">
      <author>Ralls, Kim</author>
      <title>Midnight Rain</title>
      <genre>Fantasy</genre>
      <price>5.95</price>
      <publish_date>2000-12-16</publish_date>
      <description>A former architect battles corporate zombies,
      an evil sorceress, and her own childhood to become queen
      of the world.</description>
   </book>
   <book id="bk103">
      <author>Corets, Eva</author>
      <title>Maeve Ascendant</title>
      <genre>Fantasy</genre>
      <price>5.95</price>
      <publish_date>2000-11-17</publish_date>
      <description>After the collapse of a nanotechnology
      society in England, the young survivors lay the
      foundation for a new society.</description>
   </book>
</catalog>`

func TestXPath(t *testing.T) {
	doc, err := xmlquery.Parse(strings.NewReader(xmlDoc))

	if err != nil {
		t.Fatalf("could not parse doc %v", err)
	}

	if list := xmlquery.Find(doc, "//book"); len(list) != 3 {
		t.Fatal("count(//book) != 3")
	}
	if node := xmlquery.FindOne(doc, "//book[@id='bk101']"); node == nil {
		t.Fatal("//book[@id='bk101] is not found")
	}
	if list := xmlquery.Find(doc, "//book[genre='Fantasy']"); len(list) != 2 {
		t.Fatal("//book[genre='Fantasy'] items count is not equal 2")
	}
}
