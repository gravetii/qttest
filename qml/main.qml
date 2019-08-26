import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.12

Rectangle {
    id:root
    LogArea {
		  id: log
	  }
    FileSearch {
      id: fileSearch
    }
    DownloadsList {
      id: downloadsList
    }
}