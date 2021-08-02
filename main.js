// var net = require('net');
const Speaker = require('speaker');
const Readable = require('stream').Readable
const Buffer = require('buffer').Buffer

const sine = new Readable()
sine.bitDepth = 16
sine.channels = 1
sine.sampleRate = 44100
sine.samplesGenerated = 0
sine._read = read
sine.pipe(new Speaker())

let s = 0

const freq = 800
function read(n) {
    const sampleSize = this.bitDepth / 8
    const numSamples = n / (sampleSize * this.channels)
    buf = Buffer.alloc(n)
    const t = (Math.PI * 2 * freq) / this.sampleRate
    for (let i = 0; i < numSamples; i++, s++) {
        const val = Math.round(32760 * Math.sin(t * s))
        for (let channel = 0; channel < this.channels; channel++) {
            buf.writeInt16LE(val, i * 2 * this.channels + (channel * 2))
        }
    }

    this.push(buf)
}