import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.3

Rectangle {
    width: 200
    height: 200
    Column {
        spacing: 25
        Row {
            x: 10
            y: 20
            spacing: 20
            TextField {
                width: 300
                height: 50
                id: searchBox
                placeholderText: "Search..."
                color: "#000000"
                selectionColor: "steelblue"
                selectByMouse: true
            }
            Button {
                width: 80
                height: 45
                text: "Go!"
            }
        }

        Row {
            x: 10
            spacing: 10
            ComboBox {
                currentIndex: 0
                width: 200
                height: 50
                model: ["At least", "Less than"]
            }
            TextField {
                width: 80
                height: 50
                text: "0"
                horizontalAlignment: TextInput.AlignHCenter
            }
            ComboBox {
                currentIndex: 1
                width: 100
                height: 50
                model: ["kB", "MB", "GB"]
            }
        }

        Row {
            x: 10
            spacing: 10
            ComboBox {
                currentIndex: 0
                width: 400
                height: 50
                model: [
                    "Any",
                    "video (.mp4, .mkv, .mpeg, .mov, .webm, .flv)",
                    "image (.png, .jpg, .ico, .gif)",
                    "audio (.mp3, .wav)",
                    "document (.txt, .pdf, .ppt, .doc, .xls, .csv)",
                    "compressed (.zip, .gz, .7z, .rar)",
                    "executable (.exe, .dmg, .sh)",
                ]
            }
        }
    }

    
}

