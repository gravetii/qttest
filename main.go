package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

// QmlBridge to communicate with the frontend.
type QmlBridge struct {
	core.QObject
	_ func()            `constructor:"init"`
	_ func(term string) `slot:"search"`
}

func (qmlBridge *QmlBridge) init() {
	fmt.Println("Initialising QmlBridge")
}

func (qmlBridge *QmlBridge) configure() {
	qmlBridge.ConnectSearch(func(term string) {
		fmt.Println("Searching for string: ", term)
	})
}

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML/Quick
	widgets.NewQApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	// create the quick view
	// with a minimum size of 250*200
	// set the window title to "Hello QML/Quick Example"
	// and let the root item of the view resize itself to the size of the view automatically
	view := quick.NewQQuickView(nil)
	view.SetMinimumSize(core.NewQSize2(600, 300))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Diztl")
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.configure()
	view.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	// load the embedded qml file
	// created by either qtrcc or qtdeploy
	view.SetSource(core.NewQUrl3("./qml/main.qml", 0))
	// you can also load a local file like this instead:
	//view.SetSource(core.QUrl_FromLocalFile("./qml/main.qml"))

	// make the view visible
	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	widgets.QApplication_Exec()
}
