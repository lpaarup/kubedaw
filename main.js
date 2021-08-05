const banjo = require('banjo-sound')

b = new banjo.AudioStreamBuilder()

class MyCallback extends banjo.AudioStreamDataCallback {
    freq = 440
    t = 0
    async onAudioReady(stream, audioData, numFrames) {
        const w = (Math.PI * 2 * this.freq) / stream.sampleRate
        for (let frame = 0; frame < numFrames; frame++, this.t++) {
            const val = Math.round(stream.maxAplitude * Math.sin(w * this.t))
            for (let channel = 0; channel < stream.channelCount; channel++) {
                audioData.writeInt16LE(val, frame * 2 * stream.channelCount + (channel * 2))
            }
        }
    }
}

b.setDataCallback(new MyCallback())

stream = b.newStream()
if (stream.status != banjo.StreamStatus.OK) {
    console.log("Failed to create stream. Error: " + stream.error)
}

/////////////////////// Option 1///////////////////////////////////////
// requestStart and let the stream call onAudioReady when it needs data
stream.requestStart()