package main

import (
	"fmt"
	"kk/phpipam"

	"github.com/rivo/tview"
)

var c *phpipam.Client

func main() {
	loginform()
	searchform()
}

// first window to login to the rackgo ui
// if I get a token, then we continue to the next application
func loginform() bool {
	login := tview.NewApplication()

	user := tview.NewInputField().SetLabel("user name: ")
	pwd := tview.NewInputField().SetLabel("password: ").SetMaskCharacter(' ')

	tvdata := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	tvdata.SetBorder(true).SetTitle("data")

	form := tview.NewForm().
		AddFormItem(user).
		AddFormItem(pwd).
		AddButton("enter", func() {
			var err error
			c, err = connectipam(user.GetText(), pwd.GetText())
			if err != nil {
				refreshdata(tvdata, err.Error())
				user.SetText("")
				pwd.SetText("")
			} else {
				refreshdata(tvdata, "Successfully logged in")
				user.SetText("")
				pwd.SetText("")
				login.Stop()
			}
		})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(tvdata, 0, 1, false)

	if err := login.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
	return true
}

func walksectiontree(target *tview.TreeNode, path string) {
	sections := getsections()
	for _, sect := range sections.Data {
		node := tview.NewTreeNode(sect.Name).
			SetReference(path)
		target.AddChild(node)
	}
}

func searchform() {
	kk := tview.NewApplication()
	search := tview.NewInputField().SetLabel("search ...")
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	text.SetBorder(true).SetTitle("results")

	root := "."
	rootnode := tview.NewTreeNode(root)
	tree := tview.NewTreeView().
		SetRoot(rootnode).
		SetCurrentNode(rootnode)
	walksectiontree(rootnode, root)

	form := tview.NewForm().
		AddFormItem(search).
		AddButton("enter", func() {
			sections := getsections()
			text.Clear()
			fmt.Println(text, sections)

		})
	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(text, 0, 1, false)

	if err := kk.SetRoot(flex, true).SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}

}

func getsections() phpipam.Sections {
	sections, err := c.GetSections()
	if err != nil {
		panic(err)
	}
	return sections

}

func refreshdata(data *tview.TextView, info string) {
	data.Clear()
	fmt.Fprintln(data, info)

}

func connectipam(user, passwd string) (client *phpipam.Client, err error) {
	config := phpipam.Config{
		Hostname:      "http://ipam.example.org/phpipam",
		Application:   "rackgo",
		Username:      user,
		Password:      passwd,
		SSLSkipVerify: false,
	}
	client, err = config.NewClient()
	if err != nil {
		//return nil, errors.New("Invalid username or password")
		return nil, err
	}
	return client, err
}
