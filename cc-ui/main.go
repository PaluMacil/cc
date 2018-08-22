package main

/*
	binPath = os.Getenv("CCUTIL_BIN_PATH")
	if binPath == "" {
		binPath = filepath.Join(path.Dir(os.Args[0]),"bin")
	}
*/

import (
	"log"
	"github.com/PaluMacil/cc/config"
	"github.com/andlabs/ui"
	"fmt"
)

func main() {
	err := ui.Main(func() {
		mainBox := ui.NewVerticalBox()
		frmSelected := ui.NewGroup("Selected:")
		frmCompilers := ui.NewGroup("Compilers")
		mainBox.Append(frmSelected, true)
		mainBox.Append(frmCompilers, true)

		selectedBox := ui.NewVerticalBox()
		compilersBox := ui.NewVerticalBox()

		frmSelected.SetChild(selectedBox)
		frmCompilers.SetChild(compilersBox)

		//Compiler Selection
		cboCompilers := ui.NewCombobox()
		conf, err := config.Load(config.Path())
		if err != nil {
			log.Fatalln(err)
		}
		for _, c := range conf.Compilers {
			cboCompilers.Append(c.Name)
		}
		if conf.Default != "" {
			for i, c := range conf.Compilers {
				if c.Name == conf.Default {
					cboCompilers.SetSelected(i)
				}
			}
		}
		lblSelectedCompilerLabel := ui.NewLabel("Selected Compiler")
		lblSelectedCompilerValue := ui.NewLabel("(none)")
		btnSelect := ui.NewButton("Select")

		selectedBox.Append(cboCompilers, false)
		selectedBox.Append(btnSelect, false)
		selectedBox.Append(lblSelectedCompilerLabel, false)
		selectedBox.Append(lblSelectedCompilerValue, false)

		//Compiler List
		lblCompilerName := ui.NewLabel("Name")
		txtCompilerName := ui.NewEntry()
		lblCompilerPath := ui.NewLabel("Path")
		txtCompilerPath := ui.NewEntry()
		boxAddRemove := ui.NewHorizontalBox()
		btnRemove := ui.NewButton("Remove")
		btnAdd := ui.NewButton("Add")
		boxAddRemove.Append(btnRemove, false)
		boxAddRemove.Append(btnAdd, false)
		compilersBox.Append(lblCompilerName, false)
		compilersBox.Append(txtCompilerName, false)
		compilersBox.Append(lblCompilerPath, false)
		compilersBox.Append(txtCompilerPath, false)
		compilersBox.Append(boxAddRemove, false)

		window := ui.NewWindow("CC", 400, 500, false)
		window.SetMargined(true)
		window.SetChild(mainBox)
		btnSelect.OnClicked(func(*ui.Button) {
			fmt.Println(conf.Compilers[cboCompilers.Selected()].Name, "selected.")
			fmt.Println(lblSelectedCompilerValue.Text(), "will change to new value.")
			lblSelectedCompilerValue.SetText(conf.Compilers[cboCompilers.Selected()].Name)
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
