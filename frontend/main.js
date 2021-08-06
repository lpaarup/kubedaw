const banjo = require('banjo-sound')
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDef = protoLoader.loadSync("../proto/audio_data.proto", {})
const grpcObject = grpc.loadPackageDefinition(packageDef)
const kubedawPackage = grpcObject.kubedaw

const client = new kubedawPackage.AudioData("localhost:8080", grpc.credentials.createInsecure())

b = new banjo.AudioStreamBuilder()

class MyCallback extends banjo.AudioStreamDataCallback {
    freq = 440
    t = 0
    async onAudioReady(stream, audioData, numFrames) {
        client.Request({
            numFrames: numFrames, 
            numChannels: stream.channelCount, 
            sampleRate: stream.sampleRate,
            maxAmplitude: stream.maxAplitude,
        }, (err, response) => {
            if(err != undefined) {
                return
            }
            response.Audio.copy(audioData)
        })
    }
}

b.setDataCallback(new MyCallback())

stream = b.newStream()
if (stream.status != banjo.StreamStatus.OK) {
    console.log("Failed to create stream. Error: " + stream.error)
}

stream.requestStart()