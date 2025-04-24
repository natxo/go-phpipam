package main

import (
	"fmt"
	"kk/phpipam"

	"github.com/gdamore/tcell/v2"
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

	ipamhost := tview.NewInputField().SetLabel("phpipam instance url: ")
	app := tview.NewInputField().SetLabel("api app name: ")
	user := tview.NewInputField().SetLabel("user name: ")
	pwd := tview.NewInputField().SetLabel("password: ").SetMaskCharacter(' ')
	nocertval := tview.NewCheckbox().SetLabel("disable certificate validation")

	tvdata := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	tvdata.SetBorder(true).SetTitle("data")

	form := tview.NewForm().
		AddFormItem(ipamhost).
		AddFormItem(app).
		AddFormItem(user).
		AddFormItem(pwd).
		AddFormItem(nocertval).
		AddButton("enter", func() {
			var err error
			c, err = connectipam(user.GetText(), pwd.GetText(), ipamhost.GetText(), app.GetText(), nocertval.IsChecked())
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

// walksectiontree creates a tree from every section, and then every subnet
// inside every section
func walksectiontree(target *tview.TreeNode, path string) {
	sections := getsections()
	for _, sect := range sections.Data {
		node := tview.NewTreeNode(sect.Name).
			SetReference(path)
		fmt.Println(node.GetReference())
		target.AddChild(node)
		subnets := subnets_in_section(sect.ID)
		for _, sbnt := range subnets.Data {
			leaf := tview.NewTreeNode(sbnt.Subnet).
				//SetReference(node.GetReference().(string) + "/" + sect.Name + "/" + sbnt.Subnet)
				SetReference("/" + sect.Name)
			fmt.Println(leaf.GetReference())
			node.AddChild(leaf)
		}
	}
}

func subnets_in_section(id string) phpipam.SectionsSubnets {
	subnets, err := c.GetSectionsSubnets(id)
	if err != nil {
		fmt.Println(err)
		return subnets
	}
	//fmt.Println(subnets)
	return subnets
}
func searchform() {
	kk := tview.NewApplication()
	search := tview.NewInputField().SetLabel("search ...")
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	text.SetBorder(true).SetTitle("results")

	root := "."
	rootnode := tview.NewTreeNode(root).SetColor(tcell.ColorRed)
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

// getsections retrieves all sections to which the user has access in phpipam
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

func connectipam(user, passwd, url, application string, nocertval bool) (client *phpipam.Client, err error) {
	config := phpipam.Config{
		Hostname:      url,
		Application:   application,
		Username:      user,
		Password:      passwd,
		SSLSkipVerify: nocertval,
	}
	client, err = config.NewClient()
	if err != nil {
		//return nil, errors.New("Invalid username or password")
		return nil, err
	}
	return client, err
}
