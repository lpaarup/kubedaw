const { ipcRenderer } = require('electron')

function sendNote(note) {
    ipcRenderer.send("play", note)
}