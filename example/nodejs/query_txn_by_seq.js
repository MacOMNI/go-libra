var libra = require("../../gopherjs-libra/gopherjs-libra.js").libra;

const fromHexString = hexString =>
    new Uint8Array(hexString.match(/.{1,2}/g).map(byte => parseInt(byte, 16)));

const toHexString = bytes =>
    bytes.reduce((str, byte) => str + byte.toString(16).padStart(2, '0'), '');

const defaultServer = "http://hk2.wutj.info:38080";

var addrStr = "18b553473df736e5e363e7214bd624735ca66ac22a7048e3295c9b9b9adfc26a"
var addr = fromHexString(addrStr)

var client = libra.client(defaultServer, libra.trustedPeersFile)
client.queryTransactionByAccountSeq(addr, 0, true)
    .then(r => {
        console.log("Txn #", r.getVersion())
        console.log("    Gas used (microLibra): ", r.getGasUsed())
        console.log("    Major status: ", r.getMajorStatus())
        console.log("    Events: ", r.getEvents())
        console.log("    Signed Txn: ", r.getSignedTxn())
    })
    .catch(e => {
        console.log("Error: ", e)
    })